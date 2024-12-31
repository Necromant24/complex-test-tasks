package services

import (
	"go-backend/interfaces"
	"go-backend/models"
)

type UserService struct {
	interfaces.IUserRepository
}

func (service *UserService) Get(id string) (models.User, error) {
	return service.IUserRepository.Get(id)
}

func (service *UserService) GetList() ([]models.User, error) {
	return service.IUserRepository.GetList()
}
func (service *UserService) Update(user models.User) (models.User, error) {
	return service.IUserRepository.Update(user)
}
func (service *UserService) Create(user models.User) (models.User, error) {
	return service.IUserRepository.Create(user)
}
func (service *UserService) Delete(user models.User) (models.User, error) {
	return service.IUserRepository.Delete(user)
}
