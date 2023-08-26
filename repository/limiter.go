package repository

import "github.com/kdsama/rate-limiter/entity"

type LimiterRepository interface {
	Save(le *entity.Limiter) error
	Get(shorturl string) (entity.Limiter, error)
	UpdateByShortURL(shortulr string, limiter *entity.Limiter) error
}
