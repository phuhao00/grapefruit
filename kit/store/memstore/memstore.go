package memstore

import (
	"github.com/quasoft/memstore"
	"grapefruit/internal/app/session"
)

type Store interface {
	session.Store
}

func NewStore(keyPairs ...[]byte) Store {
	return &store{memstore.NewMemStore(keyPairs...)}
}

type store struct {
	*memstore.MemStore
}

func (c *store) Options(options session.Options) {
	c.MemStore.Options = options.ToGorillaOptions()
}
