package main

import (
	"fmt"
	"github.com/arnab-xyz/todoApp/internal/dal"
	"github.com/arnab-xyz/todoApp/internal/handler"
	"github.com/arnab-xyz/todoApp/internal/model"
	"github.com/arnab-xyz/todoApp/internal/router"
	"github.com/arnab-xyz/todoApp/internal/service"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(sqlite.Open("todos.db"), &gorm.Config{})
	if err != nil {
		panic(fmt.Sprintf("Failed to start server: reason- %v", err.Error()))
	}
	db.AutoMigrate(&model.Todo{})
	dal := dal.NewDal(db)
	service := service.NewService(dal)
	handler := handler.NewHandler(service)
	server := router.NewRouter(handler)
	server.Run("localhost:8080")
}
