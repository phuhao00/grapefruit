package psql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"grapefruit/config"
	"log"
	"os"
	"sync"
	"time"
)

var gormDb *gorm.DB
var onceDo sync.Once

func InitGormDB() {
	onceDo.Do(func() {
		var err error

		psqlConfig := config.GetPSQLConfig()
		connStr := "user=%s password=%s dbname=%s host=%s port=%s sslmode=disable"
		dsn := fmt.Sprintf(connStr,
			psqlConfig.User, psqlConfig.Pwd, psqlConfig.DataBase, psqlConfig.Host, psqlConfig.Port)
		//
		dialector := postgres.Open(dsn)
		gormDb, err = gorm.Open(dialector, &gorm.Config{
			Logger: logger.New(
				log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold:             time.Second, // Slow SQL threshold
					LogLevel:                  logger.Info, // Log level
					IgnoreRecordNotFoundError: false,       // Ignore ErrRecordNotFound error for logger
				},
			),
		})

		if err != nil {
			panic(err)
		}
	})

}

func GetGormDB() *gorm.DB {
	return gormDb
}
