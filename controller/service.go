package controller

import (
	"context"
	"net/http"
	"obp-mongoDB-2/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeService struct {
	MongoCollection *mongo.Collection
}

type Response struct {
	Data  interface{} `json:"data,omitempty"`
	Error error       `json:"error,omitempty"`
}

// func (s *EmployeeService) CreateEmployee(c *gin.Context) {
// }

func (s *EmployeeService) FindEmployeeByID(c *gin.Context) {
	empID := c.Query("employee_id")

	var emp models.Employee
	err := s.MongoCollection.
		FindOne(context.Background(), bson.D{{Key: "employee_id", Value: empID}}).
		Decode(&emp)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error:": err.Error()})
	}

	c.JSON(http.StatusOK, emp)
}
