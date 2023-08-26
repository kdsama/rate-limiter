package mongo

import (
	"context"

	mongoUtils "github.com/kdsama/rate-limiter/infra/mongo"
	"go.mongodb.org/mongo-driver/bson"
)

type MongoRepo struct {
	*mongoUtils.MongoClient
}

func NewMongoRepo() *MongoRepo {
	return &MongoRepo{}
}

func (m *MongoRepo) Save(ctx context.Context, obj *bson.M, collection string) error {
	col := m.Client.Database(m.Db).Collection(collection)

	_, err := col.InsertOne(ctx, obj)
	return err
}

// Get from db
func (m *MongoRepo) GetOne(ctx context.Context, obj *bson.M, collection string) (*any, error) {
	var result any
	col := m.Client.Database(m.Db).Collection(collection)
	err := col.FindOne(ctx, obj).Decode(&result)
	return &result, err
}
