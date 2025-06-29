package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) Create(c *gin.Context) {
	id := uuid.New().String()
	s := GetStore()
	todo := &Todo{}
	err := c.BindJSON(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	todo.ID = id
	s.Todos[id] = todo
	c.JSON(http.StatusCreated, todo)
	return
}

func (h *Handler) Get(c *gin.Context) {
	id := c.Param("id")
	todos := GetStore().Todos
	if todo, ok := todos[id]; ok {
		c.JSON(http.StatusOK, todo)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})
	}
}

func (h *Handler) GetAll(c *gin.Context) {
	todos := GetStore().Todos
	c.JSON(http.StatusOK, todos)
}

func (h *Handler) Update(c *gin.Context) {
	id := c.Param("id")
	updatedTodo := &Todo{}
	err := c.BindJSON(updatedTodo)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	todos := GetStore().Todos
	if _, ok := todos[id]; ok {
		todos[id] = updatedTodo
		c.JSON(http.StatusNoContent, nil)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})
	}
}

func (h *Handler) Delete(c *gin.Context) {
	id := c.Param("id")
	todos := GetStore().Todos
	if _, ok := todos[id]; ok {
		delete(todos, id)
		c.JSON(http.StatusNoContent, nil)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Not Found",
		})
	}
}
