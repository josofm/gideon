package api

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"github.com/josofm/gideon/model"
)

type Api struct {
	server     *http.Server
	controller Controller
}

type Controller interface {
	Login(name, pass string) (map[string]interface{}, error)
	CreateUser(user model.User) (string, error)
	GetToken(header string) (model.Token, error)
}

func NewApi(c Controller) *Api {
	api := Api{
		controller: c,
	}
	return &api
}

func (api *Api) StartServer() error {
	router := api.routes()

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
	log.Print("[login] trying login")
	user := model.User{}

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&user)
	if err != nil || (model.User{}) == user {
		sendErrorMessage(w, http.StatusInternalServerError, "Invalid request - Invalid Credentials")
	}
	token, err := api.controller.Login(user.Email, user.Password)
	if err != nil {
		sendErrorMessage(w, http.StatusNotFound, "Invalid request - Invalid Credentials")
		return

	}
	log.Print("[login] login ok")
	send(w, http.StatusOK, token)
	return

}

func (api *Api) register(w http.ResponseWriter, r *http.Request) {
	log.Print("[register] trying register")
	user := model.User{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&user)
	if err != nil || (model.User{}) == user {
		sendErrorMessage(w, http.StatusInternalServerError, "Invalid request - Invalid Credentials")
	}

	email, err := api.controller.CreateUser(user)
	if err != nil { //validate kind of errors
		sendErrorMessage(w, http.StatusInternalServerError, "Invalid request - Name, sex, age, password and email are required")
		return
	}
	message := fmt.Sprintf("Welcome %v", email)
	log.Print("[register] register ok")
	send(w, http.StatusOK, message)
	return
}

func (api *Api) jwtVerify(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("[jwtVerify] triyng get token")
		var header = r.Header.Get("access-token") //Grab the token from the header
		header = strings.TrimSpace(header)
		if header == "" {
			log.Print("[jwtVerify] Token is missing, returns with error code 403 Unauthorized")
			sendErrorMessage(w, http.StatusForbidden, "Missing auth token")
			return
		}
		tk, err := api.controller.GetToken(header)
		if err != nil {
			log.Print("[jwtVerify] Token did not match, Unauthorized")
			sendErrorMessage(w, http.StatusForbidden, "Unauthorized")
			return
		}
		log.Print("[jwtVerify] token ok")
		ctx := context.WithValue(r.Context(), "user", tk)

		next.ServeHTTP(w, r.WithContext(ctx))

	})

}

func (api *Api) getUser(w http.ResponseWriter, r *http.Request) {
	log.Print("[getUser] trying get user")
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Print("[getUser] no id")
		sendErrorMessage(w, http.StatusBadRequest, "Malformed endpoint")
	}

	token, ok := r.Context().Value("user").(model.Token)
	if !ok {
		log.Print("[getUser] wrong user")
		sendErrorMessage(w, http.StatusBadRequest, "Wrong token")
	}
	tokenIdString := fmt.Sprintf("%v", token.UserID)
	if id != tokenIdString {
		log.Print("[getUser] wrong user")
		sendErrorMessage(w, http.StatusForbidden, "field not allowed")
		return
	}
	user, err := api.controller.GetUser(id)
	if err != nil {
		log.Print("[getUser] user not found")
		sendErrorMessage(w, http.StatusNotFound, "User not found")
	}

	log.Print("[getUser] User ok")
	send(w, http.StatusOK, user)
	return

}

func send(w http.ResponseWriter, code int, val interface{}) {
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, TRACE, GET, HEAD, POST, PUT")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Accept, X-Requested-With")
	w.WriteHeader(code)

	if val != nil {
		err := json.NewEncoder(w).Encode(val)
		if err != nil {
			log.Printf("error on json encoder err: %s", err.Error())
		}
	}
}

func sendErrorMessage(w http.ResponseWriter, code int, msg string) {
	log.Printf("Error - %s", msg)
	send(w, code, msg)
}

// router := mux.NewRouter()

// router.HandleFunc("/decks", controller.GetDecks(db)).Methods("GET")
// router.HandleFunc("/decks/{id}", controller.GetDeck(db)).Methods("GET")
// router.HandleFunc("/decks", controller.AddDeck(db)).Methods("POST")
// router.HandleFunc("/decks", controller.UpdateDeck(db)).Methods("PUT")
// router.HandleFunc("/decks/{id}", controller.RemoveDeck(db)).Methods("DELETE")

// fmt.Println("Server is running at port 8000")
// log.Fatal(http.ListenAndServe(":8000", router))
