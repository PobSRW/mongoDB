package main

import (
	"context"
	"log"
	"net/http"
	"obp-mongoDB-2/router"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var mongoClient *mongo.Client

func init() {
	// load env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("load env failed, error : %v", err)
	}

	log.Println("env file loaded")

	// create mongo client
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	var mongoDB_URL = os.Getenv("MONGO_URL")

	mongoClient, err = mongo.Connect(ctx, options.Client().ApplyURI(mongoDB_URL))
	if err != nil {
		log.Fatalf("connection to mongo_db failed, error : %v", err)
	}

	err = mongoClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("ping failed: %v ", err)
	}

	log.Println("connected to mongodb")
}

func main() {
	// close the mongo connection when func main doesn't working
	ctx := context.Background()
	defer mongoClient.Disconnect(ctx)

	r := router.SetupRouter(mongoClient)

	http.ListenAndServe(":8000", r)
	log.Println("server is running on port 8000")

	r.Run()
}
