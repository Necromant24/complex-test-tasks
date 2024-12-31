package interfaces

import (
	"go-backend/models"
)

type IUserRepository interface {
	Get(id string) (models.User, error)
	GetList() ([]models.User, error)
	Update(user models.User) (models.User, error)
	Delete(user models.User) (models.User, error)
	Create(user models.User) (models.User, error)
}
