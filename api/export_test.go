package api

import "net/http"

func (api *Api) Login(w http.ResponseWriter, r *http.Request) {
	api.login(w, r)
}

func (api *Api) Register(w http.ResponseWriter, r *http.Request) {
	api.register(w, r)
}

func (api *Api) GetUser(w http.ResponseWriter, r *http.Request) {
	api.getUser(w, r)
}

func (api *Api) JwtVerify(next http.Handler) http.Handler {
	return api.jwtVerify(next)
}
