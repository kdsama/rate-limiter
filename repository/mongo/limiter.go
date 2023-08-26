package mongo

import (
	"context"
	"time"

	"github.com/kdsama/rate-limiter/entity"
	"go.mongodb.org/mongo-driver/bson"
)

const (
	limiter = "limiter"
)

type Limiter struct {
	col string
	db  *MongoRepo
}

func NewLimiterRepo() *Limiter {
	db := NewMongoRepo()
	return &Limiter{col: limiter, db: db}
}

func (m *Limiter) Save(l *entity.Limiter) error {

	// col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	obj := bson.M{
		"uuid":         l.ID,
		"userid":       l.UserID,
		"url":          l.Url,
		"shorturl":     l.ShortUrl,
		"expiry":       l.Expiry,
		"limit":        l.Limit,
		"browserCache": l.BrowserCache,
		"createdAt":    l.CreatedAt,
		"updatedAt":    l.UpdatedAt,
	}
	m.db.Save(ctx, &obj, m.col)
	// _, err := col.InsertOne(
	// 	ctx,
	// 	bson.M{
	// 		"uuid":        NewBook.ID,
	// 		"name":        NewBook.Name,
	// 		"image_url":   NewBook.Image_Url,
	// 		"authors":     NewBook.Authors,
	// 		"co_authors":  NewBook.Co_Authors,
	// 		"audio_urls":  NewBook.AudiobookUrls,
	// 		"ebook_urls":  NewBook.EbookUrls,
	// 		"hard_copies": NewBook.Hardcopies,
	// 		"categories":  NewBook.Categories,
	// 		"createdAt":   NewBook.Created_Timestamp,
	// 		"updatedAt":   NewBook.Updated_Timestamp,
	// 		"verified":    false},
	// )

	// if err != nil {
	// 	return repository.ErrWriteRecord
	// }
	// return nil
	return nil
}
func (m *Limiter) Get(shorturl string) (*entity.Limiter, error) {

	return nil, nil
}
func (m *Limiter) UpdateByShortURL(shortulr string, limiter *entity.Limiter) error {

	return nil
}
