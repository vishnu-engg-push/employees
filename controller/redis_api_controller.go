package controller

import (
	"context"
	"employees/models"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"log"
	"net/http"
	"strconv"
)

var redisCtx = context.TODO()
var client *redis.Client

func EmployeeCache(redisClient *redis.Client) {
	client = redisClient
}

func AddPersonInCache(c *gin.Context) {

	var person models.Person
	c.BindJSON(&person)

	newPerson := models.Person{
		Id:      person.Id,
		Name:    person.Name,
		Active:  person.Active,
		Address: person.Address,
	}

	objToInsertInJson, _ := json.Marshal(newPerson)

	err := client.Set(redisCtx, strconv.Itoa(person.Id), objToInsertInJson, 0).Err()

	if err != nil {
		log.Printf("Error while inserting new person into cache, Reason: %v\n", err)
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

func FindOnePersonFromCache(c *gin.Context) {

	objStr, err := client.Get(redisCtx, c.Param("id")).Result()
	if err != nil {
		log.Printf("Error while getting a person from cache, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Person not found",
		})
		return
	}
	strInBytes := []byte(objStr)
	var person models.Person
	err = json.Unmarshal(strInBytes, &person)
	if err != nil {
		log.Printf("Error while converting fetched person from cache, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
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

func UpdatePersonInCache(c *gin.Context) {

	findErr := client.Get(redisCtx, c.Param("id")).Err()
	if findErr != nil {
		log.Printf("Error while getting a person from cache, Reason: %v\n", findErr)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Person not found",
		})
		return
	}

	personId, convErr := strconv.Atoi(c.Param("id"))
	if convErr != nil {
		log.Printf("Error, Reason: %v\n", convErr)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  500,
			"message": "Can't convert parameter into integer",
		})
		return
	}

	var person models.Person
	c.BindJSON(&person)

	newPerson := models.Person{
		Id:      personId,
		Name:    person.Name,
		Active:  person.Active,
		Address: person.Address,
	}

	objToUpdateInJson, _ := json.Marshal(newPerson)

	err := client.Set(redisCtx, c.Param("id"), objToUpdateInJson, 0).Err()

	if err != nil {
		log.Printf("Error while inserting new person into cache, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
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

func DeletePersonFromCache(c *gin.Context) {

	keys := []string{c.Param("id")}
	err := client.Del(redisCtx, keys...).Err()
	if err != nil {
		log.Printf("Error while deleting a person from cache, Reason: %v\n", err)
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
