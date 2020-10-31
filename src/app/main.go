package main

import (
	"github.com/muktevivk/go-gin-rest-api/src/app/config"
	"github.com/jinzhu/gorm"
	"github.com/muktevivk/go-gin-rest-api/src/app/controller"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	_ "github.com/go-sql-driver/mysql"
)

var err error

func main() {
	log.Info("Setting up db configs")
	
	config.DB, err = gorm.Open("mysql", config.DBUrl(config.BuildDBConfig()))

	if err != nil {
		log.Fatal("Error while loading the DB configurations",err)
	} 
	defer config.DB.Close()
	//config.DB.AutoMigrate(&repository.TodoRepository{})
	
	r := setupRouter()
	r.Run()

}

func setupRouter () *gin.Engine {
	r := gin.Default()
	r.GET("/health", controller.HealthCheck)
	group := r.Group("/todo/v1") 
	{
		group.POST("/", controller.CreateTodoTask)
		group.GET("/tasks", controller.GetTodoList)
		group.PUT("/:task_id", controller.UpdateStatus)
		group.DELETE("/:task_id", controller.DeleteTask)
	}

	return r
}