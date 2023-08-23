package services

import (
	"errors"

	"github.com/kdsama/rate-limiter/repository"
	"github.com/kdsama/rate-limiter/utils"
)

type RateLimiter interface {
	Save(url string, browserCache bool, limit int, expiry int64, throttle int32) (string, error)
	Get(shortUrl string) (string, error)
}

type LimiterService struct {
	prefix string
	repo   repository.LimiterRepository
	us     UserServicer
}

func NewLimiterService(prefix string) *LimiterService {
	return &LimiterService{
		prefix: prefix, // prefix for the shortened Url. Each server will have a random prefix to avoid collision
	}
}

// save the information and returns the shortened url
func (l *LimiterService) Save(userID string, url string, browserCache bool, expiry int64) (string, error) {

	//
	//

	limit, err := l.us.GetLimitByUserID(userID)
	if err != nil {
		return "", err
	}
	if limit == 0 {
		return "", errors.New("cannot save another url")
	}
	sequence := utils.GenerateShortURL()

	// check if sequence exists
	seqExists, err := l.repo.Exists(sequence)
	if err != nil {
		return "", err
	}
	if expiry == int64(0) {
		expiry = utils.OneWeekFromNow()
	}
	if !seqExists {
		err := l.repo.Save(url, browserCache, limit, expiry)
		if err != nil {
			return "", err
		}
	}

	return sequence, nil
}

// Gets the full url, and mentions the type of Redirection (browser Cache or not)
func (l *LimiterService) Get(shortUrl string) (string, error) {

	return "", nil
}
