package controller

import (
	"context"
	"employees/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
	"strconv"
)

var collection *mongo.Collection
var ctx context.Context

func EmployeeCollection(c *mongo.Database) {
	collection = c.Collection("persons")
	ctx = context.TODO()
}

func GetAllPersonsFromDB(c *gin.Context) {
	persons := []models.Person{}
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Error while getting all persons, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}
	for cursor.Next(ctx) {
		var person models.Person
		cursor.Decode(&person)
		persons = append(persons, person)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Persons",
		"data":    persons,
	})
	return
}

func AddPersonInDB(c *gin.Context) {
	var person models.Person
	c.BindJSON(&person)

	_, err := collection.InsertOne(ctx, person)

	if err != nil {
		log.Printf("Error while inserting new person into db, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Person added Successfully",
	})
	return
}

func FindOnePersonInDB(c *gin.Context) {

	personId, convErr := strconv.Atoi(c.Param("id"))
	if convErr != nil {
		log.Printf("Error, Reason: %v\n", convErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Can't convert parameter into integer",
		})
		return
	}

	person := models.Person{}
	err := collection.FindOne(ctx, bson.M{"id": personId}).Decode(&person)
	if err != nil {
		log.Printf("Error while getting a person, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Person not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Found the person",
		"data":    person,
	})
	return
}

func UpdatePersonInDB(c *gin.Context) {

	personId, convErr := strconv.Atoi(c.Param("id"))
	if convErr != nil {
		log.Printf("Error, Reason: %v\n", convErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Can't convert parameter into integer",
		})
		return
	}

	findErr := collection.FindOne(ctx, bson.M{"id": personId}).Err()
	if findErr != nil {
		log.Printf("Error while getting a person, Reason: %v\n", findErr)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Person not found",
		})
		return
	}

	var person models.Person
	c.BindJSON(&person)

	newPerson := bson.M{
		"$set": bson.M{
			"name":    person.Name,
			"active":  person.Active,
			"address": person.Address,
		},
	}

	_, err := collection.UpdateOne(ctx, bson.M{"id": personId}, newPerson)
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Person Updated Successfully",
	})
	return
}

func DeletePersonFromDB(c *gin.Context) {
	personId, convErr := strconv.Atoi(c.Param("id"))
	if convErr != nil {
		log.Printf("Error, Reason: %v\n", convErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Can't convert parameter into integer",
		})
		return
	}

	_, err := collection.DeleteOne(ctx, bson.M{"id": personId})
	if err != nil {
		log.Printf("Error while deleting a person, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Person deleted successfully",
	})
	return
}
