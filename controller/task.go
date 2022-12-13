package controller

import (
	"github.com/davidgodeness/tododemo/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

func CreateTask(c *gin.Context) {
	var t model.Task

	err := c.Bind(&t)
	if err != nil {
		return
	}

	t.Status = model.Active
	now := time.Now()
	t.CreatedAt, t.UpdatedAt = now, now
	log.Printf("%+v\n", t)
	err = model.CreateTask(t)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
}

func GetTasks(c *gin.Context) {
	tasks, err := model.GetTasks()
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func GetTask(c *gin.Context) {

}

func UpdateTask(c *gin.Context) {

}

func DeleteTask(c *gin.Context) {

}
