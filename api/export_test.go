package api

import "net/http"

func (api *Api) Login(w http.ResponseWriter, r *http.Request) {
	api.login(w, r)
}

func (api *Api) Register(w http.ResponseWriter, r *http.Request) {
	api.register(w, r)
}
