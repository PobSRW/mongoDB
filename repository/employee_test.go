package repository

import (
	"context"
	"log"
	"obp-mongoDB-2/models"
	"testing"
	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const mongoDB_URL = "mongodb://test_mongo_db:123456@0.0.0.0:2000"

func newMongoClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	mongoTestClient, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoDB_URL))
	if err != nil {
		log.Fatalf("connect to mongodb url: %s failed: %v ", mongoDB_URL, err)
	}

	err = mongoTestClient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("ping failed: %v ", err)
	}

	log.Printf("connected to mongodb url: %s", mongoDB_URL)

	return mongoTestClient
}

func TestOperationMongoDB(t *testing.T) {
	mongoTestClient := newMongoClient()
	defer mongoTestClient.Disconnect(context.Background())

	// mock data
	empID_1 := uuid.New().String()
	// empID_2 := uuid.New().String()

	// connect to collection
	coll := mongoTestClient.Database("company_db").Collection("employee_test")
	empRepo := EmployeeRepo{MongoCollection: coll}

	// insert employee 1 data
	t.Run("Insert Employee 1", func(t *testing.T) {
		emp := models.Employee{
			EmployeeID: empID_1,
			Name:       "Tony",
			Department: "Physics",
		}

		res, err := empRepo.InsertEmployee(&emp)
		if err != nil {
			t.Fatalf("insert 1 operation failed: %v", err)
		}

		t.Logf("insert 1 operation succesful with result : %v", res)
	})

}
