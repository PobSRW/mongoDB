package router

import (
	"obp-mongoDB-2/controller"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetupRouter(mongoClient *mongo.Client) *gin.Engine {

	dbName := os.Getenv("DB_NAME")
	dbCollectionName := os.Getenv("COLLECTION_NAME")

	coll := mongoClient.Database(dbName).Collection(dbCollectionName)
	empService := controller.EmployeeService{
		MongoCollection: coll,
	}

	r := gin.Default()

	r.GET("/ping", controller.ShowPong)
	r.GET("/find_employee", empService.FindEmployeeByID)

	return r
}
