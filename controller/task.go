package controller

import (
	"encoding/json"
	"fmt"
	"github.com/davidgodeness/tododemo/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

func CreateTask(c *gin.Context) {
	type Task struct {
		Name string `json:"name" binding:"required"`
	}

	var t Task

	if err := c.ShouldBindJSON(&t); err != nil {
		if err == io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty body string"})
			return
		}
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, fieldErr := range errs {
				if fieldErr.Field() == "Name" && fieldErr.Tag() == "required" {
					c.JSON(http.StatusBadRequest, gin.H{"error": "name is required"})
					return
				}
			}
		}
		if _, ok := err.(*json.SyntaxError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": "json syntax error"})
			return
		}
		if err, ok := err.(*json.UnmarshalTypeError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s type is not %s", err.Field, err.Type)})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := model.CreateTask(c.Request.Context(), t.Name)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func GetTasks(c *gin.Context) {
	tasks, err := model.GetTasks(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusNotFound, map[string]string{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {

}

func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)

	var t model.Task

	if err := c.ShouldBindJSON(&t); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
			return
		}
		errs := err.(validator.ValidationErrors)
		for _, fieldErr := range errs {
			if fieldErr.Field() == "Name" && fieldErr.Tag() == "required" {
				c.JSON(http.StatusBadRequest, map[string]string{"error": "name is required"})
				return
			}
		}
	}
}

func DeleteTask(c *gin.Context) {

}
