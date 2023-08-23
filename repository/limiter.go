package repository

type LimiterRepository interface {
	Save(url string, browserCache bool, limit int, expiry int64) error
	Exists(shorturl string) (bool, error)
}
