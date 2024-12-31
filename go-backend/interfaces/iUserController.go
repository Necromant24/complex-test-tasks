package interfaces

import "net/http"

type IUserController interface {
	Get(res http.ResponseWriter, req *http.Request)
	GetList(res http.ResponseWriter, req *http.Request)
	Update(res http.ResponseWriter, req *http.Request)
	Create(res http.ResponseWriter, req *http.Request)
	Delete(res http.ResponseWriter, req *http.Request)
}
