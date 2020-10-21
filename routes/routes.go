package routes

import (
	"employees/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Routes(router *gin.Engine) {
	router.GET("/", welcome)
	router.GET("mongo/persons", controller.GetAllPersonsFromDB)

	router.GET("mongo/persons/:id", controller.FindOnePersonInDB)
	router.POST("mongo/persons", controller.AddPersonInDB)
	router.PUT("mongo/persons/:id", controller.UpdatePersonInDB)
	router.DELETE("mongo/persons/:id", controller.DeletePersonFromDB)

	router.GET("redis/persons/:id", controller.FindOnePersonFromCache)
	router.POST("redis/persons", controller.AddPersonInCache)
	router.PUT("redis/persons/:id", controller.UpdatePersonInCache)
	router.DELETE("redis/persons/:id", controller.DeletePersonFromCache)

	router.NoRoute(notFound)
}

func welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Welcome To Employee API's",
	})
	return
}

func notFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status":  404,
		"message": "Route Not Found",
	})
	return
}
