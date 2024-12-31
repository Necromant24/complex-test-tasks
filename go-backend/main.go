package main

import (
	"fmt"
	"go-backend/config"
	"go-backend/db"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func main() {

	fmt.Println("s")

	port := config.GetConfig().HostPort

	err := db.MigrateDb()
	if err != nil {
		panic(err.Error())
	}

	usersController := InjectUsersController()

	r := chi.NewRouter()
	r.Get("/users/{userId}", usersController.Get)
	r.Get("/users", usersController.GetList)
	r.Post("/users", usersController.Create)
	r.Put("/users", usersController.Update)
	r.Delete("/users/{userId}", usersController.Delete)

	http.ListenAndServe("0.0.0.0:"+strconv.Itoa(port), r)
}
