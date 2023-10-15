package server

import (
	"context"
	"github.com/gin-gonic/gin"
	"grapefruit/internal/app/service"
	"grapefruit/internal/app/service/chat"
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
	eventSource := chat.NewEventSource()
	return &Server{
		Router:      r,
		chatService: eventSource,
	}
}

func (s *Server) Init() {

}

func (s *Server) Run() {
	s.Init()
	go func() {
		err := s.Router.Run(":8080")
		if err != nil {
			log.Error("[Run] err:%v", err.Error())
		}
	}()
	go func() {
		s.chatService.Start()
		defer s.chatService.Stop()
	}()
}

func (s *Server) WaitSignal() {
	shutdown.NewHook().Close(
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
			defer cancel()

			if err := s.Shutdown(ctx); err != nil {
				log.Error("server shutdown err:%s", err.Error())
			}
		})
}

func (s *Server) Shutdown(ctx context.Context) error {
	log.Info("tutu-backend server closed !")
	return nil
}
