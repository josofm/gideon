package api

import (
	"log"
	"net/http"

	"github.com/josofm/gideon/controller"

	"github.com/gorilla/mux"
)

type Api struct {
	server     *http.Server
	controller controller.Controller
}

func NewApi(c controller.Controller) *Api {
	api := Api{
		controller: c,
	}
	return &api
}

func (api *Api) StartServer() error {
	router := mux.NewRouter()

	//routes

	router.HandleFunc("/up", api.Up).Methods("GET")
	router.HandleFunc("/login", api.Login).Methods("POST")

	api.server = &http.Server{Addr: ":8000", Handler: router}
	log.Print("Server is running at port 8000")

	err := api.server.ListenAndServe()
	return err

}

func (api *Api) Up(w http.ResponseWriter, r *http.Request) {
	log.Print("[UP] Server is Up")
	w.WriteHeader(http.StatusOK)
}

func (api *Api) Login(w http.ResponseWriter, r *http.Request) {
	log.Print("TODO implement this")
	w.WriteHeader(http.StatusOK)
}

// router := mux.NewRouter()

// router.HandleFunc("/decks", controller.GetDecks(db)).Methods("GET")
// router.HandleFunc("/decks/{id}", controller.GetDeck(db)).Methods("GET")
// router.HandleFunc("/decks", controller.AddDeck(db)).Methods("POST")
// router.HandleFunc("/decks", controller.UpdateDeck(db)).Methods("PUT")
// router.HandleFunc("/decks/{id}", controller.RemoveDeck(db)).Methods("DELETE")

// fmt.Println("Server is running at port 8000")
// log.Fatal(http.ListenAndServe(":8000", router))
