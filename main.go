package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	repo "jawaban/repository"
	service2 "jawaban/service"
)

func dbConnection() *gorm.DB{
	configDB := map[string]string{
		"DB_Username": "root",
		"DB_Password": "",
		"DB_Port":     "3306",
		"DB_Host":     "localhost",
		"DB_Name":     "merchant",
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		configDB["DB_Username"],
		configDB["DB_Password"],
		configDB["DB_Host"],
		configDB["DB_Port"],
		configDB["DB_Name"])

	db, e := gorm.Open(mysql.Open(connectionString), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if e != nil {
		panic(e)
	}
	return db
}

func main(){
	connection := dbConnection()

	repo := repo.NewGormDBRepository(connection)

	service:= service2.NewService(repo)
	err:= service.InsertArea(10,10,"persegi")
	println(err)

}