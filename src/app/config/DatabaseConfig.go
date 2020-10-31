package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
)



var DB *gorm.DB

type DBConfig struct {
	Host string
	Port int
	User string
	DBName string
	Password string
}

func BuildDBConfig () *DBConfig {
	dbConfig := DBConfig {
		Host: "localhost",
		Port: 3306,
		User: "todo-user",
		DBName: "db",
		Password: "",
	}

	return &dbConfig
}

func DBUrl (config *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.DBName,
	   )
}