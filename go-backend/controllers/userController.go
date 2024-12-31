package controllers

import (
	"encoding/json"
	"fmt"
	"go-backend/interfaces"
	"go-backend/models"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/unrolled/render"
)

type UserController struct {
	interfaces.IUserService
}

func (controller *UserController) Get(res http.ResponseWriter, req *http.Request) {

	userId := chi.URLParam(req, "userId")

	r := render.New()

	user, err := controller.IUserService.Get(userId)
	if err != nil {
		fmt.Println(err)
	}

	if user.Id == "" {
		r.JSON(res, http.StatusNotFound, userId)
		return
	}

	r.JSON(res, http.StatusOK, user)

}

func (controller *UserController) GetList(res http.ResponseWriter, req *http.Request) {
	r := render.New()
	users, err := controller.IUserService.GetList()
	if err != nil {
		r.JSON(res, http.StatusInternalServerError, err.Error())
	}

	r.JSON(res, http.StatusOK, users)
}
func (controller *UserController) Update(res http.ResponseWriter, req *http.Request) {
	r := render.New()
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		r.JSON(res, http.StatusBadRequest, err.Error())
	}

	_, err = controller.IUserService.Update(user)
	if err != nil {
		r.JSON(res, http.StatusInternalServerError, err.Error())
	}

	r.JSON(res, http.StatusOK, user)
}
func (controller *UserController) Create(res http.ResponseWriter, req *http.Request) {
	r := render.New()
	var user models.User
	err := json.NewDecoder(req.Body).Decode(&user)
	if err != nil {
		r.JSON(res, http.StatusBadRequest, err.Error())
	}

	_, err = controller.IUserService.Create(user)
	if err != nil {
		r.JSON(res, http.StatusInternalServerError, err.Error())
	}

	r.JSON(res, http.StatusOK, user)
}

func (controller *UserController) Delete(res http.ResponseWriter, req *http.Request) {
	userId := chi.URLParam(req, "userId")
	r := render.New()
	users, err := controller.IUserService.Delete(models.User{Id: userId})
	if err != nil {
		r.JSON(res, http.StatusInternalServerError, err.Error())
	}

	r.JSON(res, http.StatusOK, users)
}
