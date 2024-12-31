package repositories

import (
	"fmt"
	"go-backend/db"
	"go-backend/models"
)

type UserRepository struct {
}

func (repo *UserRepository) Get(id string) (models.User, error) {

	conn, err := db.InitConnection()
	if err != nil {
		fmt.Println("userRepo get(id) error - " + err.Error())
	}
	defer db.CloseConnection(conn)

	var user models.User
	conn.First(&user, "id = ?", id)

	return user, err
}

func (service *UserRepository) GetList() ([]models.User, error) {
	conn, err := db.InitConnection()
	if err != nil {
		fmt.Println("userRepo get error - " + err.Error())
	}
	defer db.CloseConnection(conn)

	var users []models.User
	conn.Find(&users)

	return users, err
}
func (service *UserRepository) Update(user models.User) (models.User, error) {
	conn, err := db.InitConnection()
	if err != nil {
		fmt.Println("userRepo get error - " + err.Error())
	}
	defer db.CloseConnection(conn)

	err = conn.Model(&user).Where("id = ?", user.Id).Update("name", user.Name).Error

	return user, err
}
func (service *UserRepository) Create(user models.User) (models.User, error) {
	conn, err := db.InitConnection()
	if err != nil {
		fmt.Println("userRepo get error - " + err.Error())
	}
	defer db.CloseConnection(conn)

	err = conn.Create(user).Error

	return user, err
}
func (service *UserRepository) Delete(user models.User) (models.User, error) {
	conn, err := db.InitConnection()
	if err != nil {
		fmt.Println("userRepo get error - " + err.Error())
	}
	defer db.CloseConnection(conn)

	err = conn.Delete(user).Error

	return user, err
}
