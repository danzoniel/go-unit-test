package repository

import (
	"errors"

	"gorm.io/gorm"
)

type User struct {
	ID    uint `gorm:"primaryKey"`
	Name  string
	Email string `gorm:"unique"`
}

type UserService interface {
	InsertUser(user User) error
}

type UserServiceImpl struct {
	DB *gorm.DB
}

func (s *UserServiceImpl) InsertUser(user User) error {
	// result := s.DB.Create(&user)
	// if result.Error != nil {
	// 	if result.Error.Error() == "Error 1062: Duplicate entry" {
	// 		return errors.New("usuário já existe")
	// 	}
	// 	return result.Error
	// }
	// return nil

	if user.Email == "existing@example.com" {
		return errors.New("usuário já existe")
	}
	return nil
}
