package services

import (
	"errors"
	"time"

	"github.com/kdsama/rate-limiter/entity"
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
	lc     repository.LimiterCacher
}

func NewLimiterService(prefix string) *LimiterService {
	lc := repository.NewLimiterCache()
	return &LimiterService{
		prefix: prefix, // prefix for the shortened Url. Each server will have a random prefix to avoid collision
		lc:     lc,
	}
}

// save the information and returns the shortened url
func (l *LimiterService) Save(userID string, url string, browserCache bool, expiry int64) (string, error) {

	limit, err := l.us.GetLimitByUserID(userID)
	if err != nil {
		return "", err
	}
	if limit == 0 {
		return "", errors.New("cannot save another url")
	}
	sequence := utils.GenerateShortURL()

	// check if sequence exists
	_, err = l.repo.Get(sequence)
	if err != nil {
		return "", err
	}
	if expiry == int64(0) {
		expiry = utils.OneWeekFromNow()
	}

	err = l.repo.Save(url, browserCache, limit, expiry)
	if err != nil {
		return "", err
	}

	return sequence, nil
}

// Gets the full url, and mentions the type of Redirection (browser Cache or not)
func (l *LimiterService) Get(shortUrl string) (string, error, bool) {
	// Check in memory
	var rlc *entity.Limiter
	rlc, err := l.lc.Get(shortUrl)
	if err != nil {
		rlc, err = l.repo.Get(shortUrl)
		if err != nil {
			return "", err, false
		}
	}
	if time.Now().Unix() > rlc.Expiry {
		return "", errors.New("url expired"), false
	}

	// else find in db

	return rlc.Url, nil, rlc.BrowserCache
}

// Gets the full url, and mentions the type of Redirection (browser Cache or not)
func (l *LimiterService) Reset(shortUrl string) error {
	// Check in memory
	var rlc *entity.Limiter
	rlc, err := l.repo.Get(shortUrl)
	if err != nil {
		return err
	}
	rlc.Expiry = utils.OneWeekFromNow()
	_, err = l.lc.Get(rlc.ShortUrl)
	if err == nil {
		// means it probably doesnot exist
		l.lc.Set(rlc.ShortUrl, rlc)
	}
	return l.repo.UpdateByShortURL(rlc.ShortUrl, rlc)
	// else find in db
}
