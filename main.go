package main

import (
	"github.com/go-unit-test/repository"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "user:password@tcp(127.0.0.1:3306)/database_name?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("falha ao conectar ao banco de dados")
	}

	db.AutoMigrate(&repository.User{})

	userService := &repository.UserServiceImpl{DB: db}

	user := repository.User{Name: "John", Email: "john@example.com"}
	err = userService.InsertUser(user)
	if err != nil {
		panic("falha ao inserir usuário")
	}
}
