package mongo

import (
	"context"

	mongoUtils "github.com/kdsama/rate-limiter/infra/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepo struct {
	*mongoUtils.MongoClient
}

func NewMongoRepo(mu *mongoUtils.MongoClient) *MongoRepo {
	return &MongoRepo{mu}
}

func (m *MongoRepo) Save(ctx context.Context, obj *bson.M, collection string) error {
	col := m.Client.Database(m.Db).Collection(collection)

	_, err := col.InsertOne(ctx, obj)
	return err
}

// Get from db
func (m *MongoRepo) GetOne(ctx context.Context, obj *bson.M, collection string) *mongo.SingleResult {
	col := m.Client.Database(m.Db).Collection(collection)
	result := col.FindOne(ctx, obj)

	return result
}
