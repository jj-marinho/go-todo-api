package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-basic-api/models"
	"gopkg.in/mgo.v2/bson"
)

func GetAllTodos(c *gin.Context) {
	todos, err := models.GetAllTodos()
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, todos)
}

func CreateTodo(c *gin.Context) {
	var todo models.Todo

	err := c.BindJSON(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = models.CreateTodo(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

func GetTodo(c *gin.Context) {
	id := c.Param("id")

	todo, err := models.GetTodo(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, todo)
}

func UpdateTodo(c *gin.Context) {
	id := c.Param("id")

	var todo models.Todo
	err := c.BindJSON(&todo)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	todo.Id = bson.ObjectIdHex(id)

	err = models.UpdateTodo(id, &todo)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}

func DeleteTodo(c *gin.Context) {
	id := c.Param("id")

	err := models.DeleteTodo(id)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Writer.WriteHeader(http.StatusOK)
}
