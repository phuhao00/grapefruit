package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/phuhao00/spoor"
	"github.com/phuhao00/spoor/logger"
	"grapefruit/internal/adapter/psql"
	"grapefruit/internal/adapter/routing"
	"grapefruit/internal/app/service"
	"grapefruit/internal/app/service/chatgpt"
	"grapefruit/kit/log"
	"grapefruit/kit/shutdown"
	"time"
)

type Server struct {
	Router      *gin.Engine
	chatService service.IChatService
}

func NewServer() *Server {
	r := gin.Default()
	routing.Register(r)
	eventSource := chatgpt.NewEventSource()
	return &Server{
		Router:      r,
		chatService: eventSource,
	}
}

func (s *Server) Init() {
	psql.InitGormDB()
	logger.SetLogging(&logger.LoggingSetting{
		Dir:          "./log",
		Level:        int(spoor.DEBUG),
		Prefix:       "",
		WriterOption: nil,
	})
}

func (s *Server) Run() {
	s.Init()

	go func() {
		err := s.Router.Run(":8080")
		if err != nil {
			panic(err.Error())
		}
	}()
	//go func() {
	//	s.chatService.Start()
	//	defer s.chatService.Stop()
	//}()

	s.WaitSignal()
}

func (s *Server) WaitSignal() {
	shutdown.NewHook().Close(
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			if err := s.Shutdown(ctx); err != nil {
				log.Error("server shutdown errwrap:%s", err.Error())
			}
		},
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			if err := s.Shutdown(ctx); err != nil {
				log.Error("server shutdown errwrap:%s", err.Error())
			}
		},
	)
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Info("tutu-backend server closed !")
	return nil
}
