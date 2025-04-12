package server

import (
	"context"
	"sync"
)

type sessionStoreKey struct{}

type SessionStore struct {
	lock sync.RWMutex
	Meta map[string]string
}

func (s *SessionStore) Get(key string) string {
	s.lock.RLock()
	defer s.lock.RUnlock()
	return s.Meta[key]
}

func (s *SessionStore) Set(key, value string) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.Meta[key] = value
}

func SessionStoreFromCtx(ctx context.Context) *SessionStore {
	session := ctx.Value(sessionStoreKey{})
	if session == nil {
		return nil
	}
	return session.(*SessionStore)
}
