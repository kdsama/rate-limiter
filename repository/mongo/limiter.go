package mongo

import (
	"context"
	"fmt"
	"time"

	"github.com/kdsama/rate-limiter/entity"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	limiter = "limiter"
)

type Limiter struct {
	col string
	db  *MongoRepo
}

func NewLimiterRepo(col string, db *MongoRepo) *Limiter {

	return &Limiter{col: limiter, db: db}
}

func (m *Limiter) Save(l *entity.Limiter) error {

	// col := g.repo.Client.Database(g.repo.Db).Collection(g.current)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	fmt.Println("Expiry being saved is ", l.Expiry)
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
func (m *Limiter) Get(shorturl string) (entity.Limiter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	obj := bson.M{"shorturl": shorturl}
	var respObj entity.Limiter
	res := m.db.GetOne(ctx, &obj, m.col)

	err := res.Decode(&respObj)
	if err != nil && err != mongo.ErrNoDocuments {

		return entity.Limiter{}, err
	}
	return respObj, nil
}
func (m *Limiter) UpdateByShortURL(shortulr string, limiter *entity.Limiter) error {

	return nil
}
