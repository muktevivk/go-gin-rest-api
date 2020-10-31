package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/muktevivk/go-gin-rest-api/src/model"
	"github.com/muktevivk/go-gin-rest-api/src/repository"
	log "github.com/sirupsen/logrus"
)

// GetTodoList controller function
func GetTodoList(context *gin.Context) {
	var todo []model.Todo
	log.Info("Calling get TodoList function")
	err := repository.GetTodoList(&todo)
	if err != nil {
		log.Error("Something went wrong for getTodoList",err)
		context.AbortWithStatus(http.StatusInternalServerError)
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

//CreateTodoTask function is POST call to create a task 
func CreateTodoTask (context *gin.Context) {
	var todo model.Todo
	context.BindJSON(&todo)
	todo.ID = uuid.New().String()
	log.Info("Calling create todo task function")
	err := repository.CreateTask(&todo)
	if err != nil {
		log.Error("Something went wrong while posting todo", err)
		context.AbortWithStatus(http.StatusInternalServerError)
	} else {
		context.JSON(http.StatusCreated, todo)
	}

}

//UpdateStatus function is PUT call to update the task
func UpdateStatus (context *gin.Context) {
	var todo model.Todo
	context.BindJSON(&todo)
	todo.ID = context.Params.ByName("task_id")
	err := repository.UpdateStatus(&todo)
	if err != nil {
		log.Error("Something went wrong while updating todo stauts", err)
		context.AbortWithStatus(http.StatusInternalServerError)
	} else {
		context.JSON(http.StatusOK, todo)
	}
}

//DeleteTask function is DELETE call to delete todo task
func DeleteTask(context *gin.Context) {
	var todo model.Todo
	id := context.Params.ByName("task_id")
	err := repository.DeleteTask(&todo, id)
	if err != nil {
		log.Error("Something went wrong while deleting todo task", err)
		context.AbortWithStatus(http.StatusInternalServerError)
	} else {
		context.JSON(http.StatusOK, id)
	}
}