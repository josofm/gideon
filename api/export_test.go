package api

import "net/http"

func (api *Api) Login(w http.ResponseWriter, r *http.Request) {
	api.login(w, r)
}
