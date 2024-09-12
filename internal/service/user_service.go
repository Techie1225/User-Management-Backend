package service

import (
	"user-management-backend/internal/models"
	"user-management-backend/internal/repository"
)

func GetUsers() ([]models.User, error) {
	return repository.GetUsers()
}

func GetUserByID(id int) (models.User, error) {
	return repository.GetUserByID(id)
}

func AddUser(user models.User) error {
	return repository.AddUser(user)
}

func UpdateUser(user models.User) error {
	return repository.UpdateUser(user)
}

func DeleteUser(id int) error {
	return repository.DeleteUser(id)
}
