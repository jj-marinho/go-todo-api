package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-basic-api/models"
)

func GetAllTodos(c *gin.Context) {
	todos := models.GetAllTodos()

	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	err := c.BindJSON(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	models.CreateTodo(&todo)
	c.JSON(http.StatusOK, todo)
}

func GetTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	todo, err := models.GetTodo(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var todo models.Todo
	err = c.BindJSON(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = models.UpdateTodo(id, &todo)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func DeleteTodo(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	models.DeleteTodo(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}
