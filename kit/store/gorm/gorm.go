package gormstore

import (
	"github.com/wader/gormstore/v2"
	"gorm.io/gorm"
	"grapefruit/internal/app/session"
	"time"
)

type Store interface {
	session.Store
}

func NewStore(d *gorm.DB, expiredSessionCleanup bool, keyPairs ...[]byte) Store {
	s := gormstore.New(d, keyPairs...)
	if expiredSessionCleanup {
		quit := make(chan struct{})
		go s.PeriodicCleanup(1*time.Hour, quit)
	}
	return &store{s}
}

type store struct {
	*gormstore.Store
}

func (s *store) Options(options session.Options) {
	s.Store.SessionOpts = options.ToGorillaOptions()
}
