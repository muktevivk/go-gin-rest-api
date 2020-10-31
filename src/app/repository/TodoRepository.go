package repository

import (
	"github.com/muktevivk/go-gin-rest-api/src/config"
	log "github.com/sirupsen/logrus"
	"github.com/muktevivk/go-gin-rest-api/src/model"
)

//GetTodoList from the DB based on the state
func GetTodoList(todoList *[]model.Todo) (err error) {
	log.Info("Getting todo list from repository")
	if err := config.DB.Find(todoList).Error; err != nil {
		log.Error("Something went wrong while fetching data from DB!")
		return err
	}
	return nil
}


//CreateTask is to add tasks to the todo list
func CreateTask(todo *model.Todo) (err error) {
	log.Info("Creating a task in the todo list")
	if err := config.DB.Create(todo).Error; err != nil {
		log.Error("Something went wrong while creating a task in DB ",err)
		return err
	}
	return nil
}

//UpdateStatus is to update status of a task
func UpdateStatus(todo *model.Todo) (err error) {
	log.Info("Updating task status")
	if err := config.DB.Save(todo).Error; err != nil {
		log.Error("Something went wrong while updating task status in DB ",err)
		return err
	}
	return nil
}


//DeleteTask is to delete task from todo
func DeleteTask (todo *model.Todo, id string) (err error) {
	log.Info("Deleting task from todo list")
	if err := config.DB.Where("id = ?", id).Delete(todo).Error; err != nil {
		log.Error("Something went wrong while deleting task in DB ",err)
		return err
	}
	return nil

}