package repository

import (
	"errors"
	"sync"

	"github.com/kdsama/rate-limiter/entity"
)

var (
	ErrKeyNotFound  = errors.New("key not found in the cache")
	ErrLimiterIsNil = errors.New("cannot accept nil as value for map")
)

type LimiterCacher interface {
	Get(key string) (*entity.Limiter, error)
	Set(key string, value *entity.Limiter) error
	Delete(key string) error
}

type mp map[string]*entity.Limiter

type LimiterCache struct {
	mp
	mut sync.RWMutex
}

func NewLimiterCache() *LimiterCache {
	return &LimiterCache{
		mp:  map[string]*entity.Limiter{},
		mut: sync.RWMutex{},
	}
}

func (hm *LimiterCache) Get(key string) (*entity.Limiter, error) {
	hm.mut.RLock()
	defer hm.mut.RUnlock()
	if _, ok := hm.mp[key]; !ok {
		return nil, ErrKeyNotFound
	}
	return hm.mp[key], nil
}
func (hm *LimiterCache) Set(key string, value *entity.Limiter) error {
	if value == nil {
		return ErrLimiterIsNil
	}
	hm.mut.Lock()
	defer hm.mut.Unlock()
	hm.mp[key] = value
	return nil
}

func (hm *LimiterCache) Delete(key string) error {
	hm.mut.Lock()
	defer hm.mut.Unlock()
	delete(hm.mp, key)
	return nil
}
