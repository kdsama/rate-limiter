package entity

import (
	"time"

	"github.com/kdsama/rate-limiter/utils"
)

type Limiter struct {
	ID           string `json:"uuid" bson:"uuid"`
	UserID       string `json:"userid" bson:"userid"`
	Url          string `json:"url" bson:"url"`
	ShortUrl     string `json:"shorturl" bson:"shorturl"`
	Expiry       int64  `json:"expiryTimestamp" bson:"expiryTimestamp"`
	Limit        int64  `json:"limit" bson:"limit"`
	BrowserCache bool   `json:"bcache" bson:"bcache"`

	CreatedAt int64 `bson:"createdAt"`
	UpdatedAt int64 `bson:"updatedAt"`
}

func NewLimiter(url, userid, shorturl string, expiry int64) *Limiter {
	id := utils.GenerateUUID()
	t := time.Now().Unix()
	return &Limiter{
		ID:        id,
		UserID:    userid,
		Url:       url,
		ShortUrl:  shorturl,
		CreatedAt: t,
		UpdatedAt: t,
	}
}
