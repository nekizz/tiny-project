package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/nekizz/tiny-project/go-gorm/config"
	"github.com/nekizz/tiny-project/go-gorm/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := config.New(mysql.Open("gorm:secret@tcp(localhost:3306)/gorm?charset=utf8mb4&parseTime=true&loc=UTC&autocommit=true"))
	if err != nil {
		panic(err)
	}

	query := db.Model(&model.User{}).Create(&model.User{
		Model:  gorm.Model{},
		Name:   "Hieu",
		Email:  "hieu@gmail.com",
		Gender: "Male",
		Age:    "12",
		Note:   "",
	})

	if err := query.Error; err != nil {
		fmt.Println(err)
	}

}
