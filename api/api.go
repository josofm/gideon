package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/josofm/gideon/model"
)

type Api struct {
	server     *http.Server
	controller Controller
}

type Controller interface {
	Login(name, pass string) (model.User, error)
}

func NewApi(c Controller) *Api {
	api := Api{
		controller: c,
	}
	return &api
}

func (api *Api) StartServer() error {
	router := mux.NewRouter()

	//routes

	router.HandleFunc("/up", api.Up).Methods("GET")
	router.HandleFunc("/login", api.login).Methods("POST")

	api.server = &http.Server{Addr: ":8000", Handler: router}
	log.Print("Server is running at port 8000")

	err := api.server.ListenAndServe()
	return err

}

func (api *Api) Up(w http.ResponseWriter, r *http.Request) {
	log.Print("[UP] Server is Up")
	w.WriteHeader(http.StatusOK)
}

func (api *Api) login(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(user)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	resp, err := api.controller.Login(user.Email, user.Password)
	if err != nil {
		var resp = map[string]interface{}{"status": false, "message": "Invalid request"}
		json.NewEncoder(w).Encode(resp)
		return
	}
	json.NewEncoder(w).Encode(resp)
	return

}

// router := mux.NewRouter()

// router.HandleFunc("/decks", controller.GetDecks(db)).Methods("GET")
// router.HandleFunc("/decks/{id}", controller.GetDeck(db)).Methods("GET")
// router.HandleFunc("/decks", controller.AddDeck(db)).Methods("POST")
// router.HandleFunc("/decks", controller.UpdateDeck(db)).Methods("PUT")
// router.HandleFunc("/decks/{id}", controller.RemoveDeck(db)).Methods("DELETE")

// fmt.Println("Server is running at port 8000")
// log.Fatal(http.ListenAndServe(":8000", router))
