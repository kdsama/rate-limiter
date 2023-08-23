package services

type RateLimiter interface {
	Save()
	Get()
}

type Limiter struct{}

func NewLimiter() *Limiter {
	return &Limiter{}
}
