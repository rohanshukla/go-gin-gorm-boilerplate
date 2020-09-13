package main

import (
	"fmt"
	"log"
	"os"
	"webapp/routers"
	"webapp/utils"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func middleware(c *gin.Context) {
	// Go Log Request
	// Authenticate Tokens and much more
	fmt.Println("I am Middleware")
}

func main() {

	envLoadError := godotenv.Load()
	if envLoadError != nil {
		log.Fatal("Error loading .env file")
	}

	// go run . --> To Run all files
	r := gin.Default()
	r.Use(middleware)

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "Rohan",
			"store": utils.DBStore,
		})
	})

	v1Routes := r.Group("/v1")
	{
		// User Routes
		userRoutes := v1Routes.Group("/user")
		{
			userRoutes.GET("/", routers.GetUsers)
			userRoutes.POST("/", routers.CreateUser)

			// Database Call
			userRoutes.POST("/get", routers.GetAllUsersDB) // This API is made with Field Projection thats why POST is used
			userRoutes.GET("/get/:id", routers.GetUserDB)
			userRoutes.POST("/create", routers.CreateUserDB)
			userRoutes.PUT("/update/:id", routers.UpdateUser)
		}

		// Todo Routes
		todoRoutes := v1Routes.Group("/todo")
		{
			todoRoutes.POST("/", routers.CreateTodo)
			todoRoutes.GET("/:uid", routers.GetUserTodo)

			// Database Calls
			todoRoutes.POST("/create", routers.CreateTodoDB)
			todoRoutes.POST("/todoByUser", routers.GetAllTodoByUser)
			todoRoutes.DELETE("/:id", routers.DeleteTodo)
			todoRoutes.PUT("/", routers.UpdateTodo)
		}
	}

	port := os.Getenv("PORT")
	if err := r.Run(":" + port); err != nil {
		log.Fatal(err.Error())
	}
}
