package main

import (
	"github.com/PattanasakGit/sa-65-example/controller"
	"github.com/PattanasakGit/sa-65-example/entity"
	"github.com/gin-gonic/gin"
)

func main() {

	entity.SetupDatabase()
	r := gin.Default()

	// User Routes
	r.GET("/users", controller.ListUsers)
	r.GET("/user/:id", controller.GetUser)
	r.POST("/users", controller.CreateUser)
	r.PATCH("/users", controller.UpdateUser)
	r.DELETE("/users/:id", controller.DeleteUser)

	// Run the server
	r.Run()
}
