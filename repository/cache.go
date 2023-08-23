package repository

import (
	"errors"
	"sync"

	"github.com/kdsama/rate-limiter/entity"
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
		return nil, errors.New("key not found")
	}
	return hm.mp[key], nil
}
func (hm *LimiterCache) Set(key string, value *entity.Limiter) error {
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
