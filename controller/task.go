package controller

import (
	"encoding/json"
	"fmt"
	"github.com/davidgodeness/tododemo/model"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"strconv"
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
	q := make(gin.H)
	var err error
	status := c.Query("status")
	if status != "" {
		var s int
		s, err = strconv.Atoi(status)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status value"})
			return
		}
		if uint8(s) != model.Active && uint8(s) != model.Done {
			c.JSON(http.StatusBadRequest, gin.H{"error": "invalid status value"})
			return
		}
		q["status"] = int64(s)
	}
	tasks, err := model.GetTasks(c.Request.Context(), q)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {

}

func UpdateTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}

	type Task struct {
		Status uint8 `json:"status" binding:"oneof=0 1"`
	}

	var t Task

	if err := c.ShouldBindJSON(&t); err != nil {
		if err == io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"error": "empty body string"})
			return
		}
		if errs, ok := err.(validator.ValidationErrors); ok {
			for _, fieldError := range errs {
				if fieldError.Field() == "Status" && fieldError.Tag() == "oneof" {
					c.JSON(http.StatusBadRequest, gin.H{"error": "status need to be one of 0,1"})
					return
				}
			}
		}
		if err, ok := err.(*json.UnmarshalTypeError); ok {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("%s type is not %s", err.Field, err.Type)})
			return
		}
	}

	r, err := model.UpdateTask(c.Request.Context(), int64(id), gin.H{"status": t.Status})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if r == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not exist"})
		return
	}

	c.Status(http.StatusOK)
}

func DeleteTask(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id param"})
		return
	}

	r, err := model.DeleteTask(c.Request.Context(), int64(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if r == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "id not exist"})
		return
	}

	c.Status(http.StatusOK)
}
