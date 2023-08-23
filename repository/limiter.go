package repository

import "github.com/kdsama/rate-limiter/entity"

type LimiterRepository interface {
	Save(url string, browserCache bool, limit int, expiry int64) error
	Get(shorturl string) (*entity.Limiter, error)
	UpdateByShortURL(shortulr string, limiter *entity.Limiter) error
}
