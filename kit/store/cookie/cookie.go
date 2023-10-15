package cookie

import (
	gsessions "github.com/gorilla/sessions"
	"grapefruit/internal/app/session"
)

type Store interface {
	session.Store
}

func NewStore(keyPairs ...[]byte) Store {
	return &store{gsessions.NewCookieStore(keyPairs...)}
}

type store struct {
	*gsessions.CookieStore
}

func (c *store) Options(options session.Options) {
	c.CookieStore.Options = options.ToGorillaOptions()
}
