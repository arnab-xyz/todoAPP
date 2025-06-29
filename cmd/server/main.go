package main

import (
	handler "github.com/arnab-xyz/todoApp/internal/todo"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	h := handler.NewHandler()

	server.GET("/todos", h.GetAll)

	server.POST("/todos", h.Create)

	server.PUT("/todos/:id", h.Update)

	server.GET("/todos/:id", h.Get)

	server.DELETE("/todos/:id", h.Delete)

	server.Run("localhost:8080")
}
