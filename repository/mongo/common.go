package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type MongoRepo struct {
	// repo *mongoUtils.MongoClient
	db string
}

func NewMongoRepo() *MongoRepo {
	return &MongoRepo{}
}

func (m *MongoRepo) Save(ctx context.Context, obj *bson.M, collection string) error {
	col := g.repo.Client.Database(m.db).Collection(collection)

	_, err := col.InsertOne(ctx, obj)

	return nil
}

func (m *MongoRepo) Get(ctx context.Context, obj *bson.M, collection string) (*bson.M, error) {

	return nil, nil
}
