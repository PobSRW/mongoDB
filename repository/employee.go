package repository

import (
	"context"
	"obp-mongoDB-2/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	MongoCollection *mongo.Collection
}

func (r *EmployeeRepo) InsertEmployee(emp *models.Employee) (*mongo.InsertOneResult, error) {
	result, err := r.MongoCollection.InsertOne(context.Background(), emp, nil)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (r *EmployeeRepo) FindByEmployeeID(empID string) (*models.Employee, error) {
	var emp models.Employee

	err := r.MongoCollection.
		FindOne(context.Background(), bson.D{{Key: "employee_id", Value: empID}}).
		Decode(&emp)

	if err != nil {
		return nil, err
	}

	return &emp, nil
}

func (r *EmployeeRepo) FindAllEmployee() ([]models.Employee, error) {
	results, err := r.MongoCollection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	var emps []models.Employee
	err = results.All(context.Background(), &emps)
	if err != nil {
		return nil, err
	}

	return emps, err
}

func (r *EmployeeRepo) UpdateEmployeeByID(empID string, updateEmp *models.Employee) (int64, error) {
	result, err := r.MongoCollection.UpdateOne(context.Background(),
		bson.D{{Key: "employee_id", Value: empID}},
		bson.D{{Key: "$set", Value: updateEmp}})
	if err != nil {
		return 0, err
	}

	return result.ModifiedCount, nil
}

func (r *EmployeeRepo) DeleteEmployeeByID(empID string) (int64, error) {
	result, err := r.MongoCollection.DeleteOne(context.Background(),
		bson.D{{Key: "employee_id", Value: empID}})
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}

func (r *EmployeeRepo) DeleteAllEmployee() (int64, error) {
	result, err := r.MongoCollection.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return 0, err
	}

	return result.DeletedCount, nil
}
