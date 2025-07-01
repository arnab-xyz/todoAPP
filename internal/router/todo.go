package router

import (
	"github.com/arnab-xyz/todoApp/internal/handler"
	"github.com/gin-gonic/gin"
)

func NewRouter(h *handler.Handler) *gin.Engine {
	server := gin.Default()

	server.POST("/todos", h.Create)

	server.GET("/todos/:id", h.Get)

	server.GET("/todos", h.GetAll)

	server.PUT("/todos/:id", h.Update)

	server.DELETE("/todos/:id", h.Delete)

	return server

}
