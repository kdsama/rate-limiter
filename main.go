package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/kdsama/rate-limiter/handler"
	mc "github.com/kdsama/rate-limiter/infra/mongo"
	"github.com/kdsama/rate-limiter/repository/mongo"
	"github.com/kdsama/rate-limiter/services"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	mclient := ConnectMongo(ctx)
	commonRepo := mongo.NewMongoRepo(mclient)
	lmp := mongo.NewLimiterRepo("limiter", commonRepo)
	lserv := services.NewLimiterService("prefix", lmp)
	lhandle := handler.NewLimiter(lserv)
	router := http.NewServeMux()
	router.HandleFunc("/api/v1/set", lhandle.HandleSave)
	router.HandleFunc("/", lhandle.Handle)

	log.Fatal("listening now ", http.ListenAndServe(":8090", router))

}
func ConnectMongo(ctx context.Context) *mc.MongoClient {
	mongoClient := mc.GetMongoClient("mongodb://localhost:27017", "somethingelse")
	err := mongoClient.Client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	return mongoClient
}
