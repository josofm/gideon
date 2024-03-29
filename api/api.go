package api

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/josofm/gideon/model"
	"github.com/josofm/mtg-sdk-go"
)

type Api struct {
	server     *http.Server
	controller Controller
}

type Controller interface {
	Login(name, pass string) (map[string]interface{}, error)
	CreateUser(user model.User) (string, error)
	GetToken(header string) (model.Token, error)
	GetUser(id uint) (model.User, error)
	GetCardByName(name string) ([]*mtg.Card, error)
	CreateDeck(deck model.Deck, userId uint) (string, error)
	DeleteUser(id uint) error
	UpdateUser(user model.User) error
}

func NewApi(c Controller) *Api {
	api := Api{
		controller: c,
	}
	return &api
}

func (api *Api) StartServer() error {
	router := api.routes()
	muxWithMiddlewares := http.TimeoutHandler(router, time.Second*30, "Timeout!")

	api.server = &http.Server{
		Addr:    ":80",
		Handler: muxWithMiddlewares,
	}
	log.Print("Server is running at port 80")

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
	defer r.Body.Close()

	email, err := api.controller.CreateUser(user)
	if err != nil { //validate kind of errors
		if err.Error() == "user already registred" {
			sendErrorMessage(w, http.StatusConflict, "This user is already register in our system")
			return
		}
		sendErrorMessage(w, http.StatusInternalServerError, "Invalid request - Name, sex, age, password and email are required")
		return
	}
	message := fmt.Sprintf("Welcome %v", email)
	log.Print("[register] register ok")
	send(w, http.StatusOK, message)
	return
}

func (api *Api) jwtVerify(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("[jwtVerify] triyng get token")
		var header = r.Header.Get("access-token")
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
	methodName := "getUser"
	log.Printf("[%v] trying get user", methodName)
	statusCode, userID, err := api.validateTokenUser(methodName, r)
	if err != nil {
		sendErrorMessage(w, statusCode, err.Error())
		return
	}

	user, err := api.controller.GetUser(uint(userID))
	if err != nil {
		log.Print("[getUser] user not found")
		sendErrorMessage(w, http.StatusNotFound, "User not found")
		return
	}
	log.Print("[getUser] User ok")
	send(w, http.StatusOK, user)
	return

}

func (api *Api) getCardByName(w http.ResponseWriter, r *http.Request) {
	log.Print("[getCard] trying get card")
	vars := mux.Vars(r)
	cardName, ok := vars["name"]
	if !ok {
		log.Print("[getCard] no name")
		sendErrorMessage(w, http.StatusBadRequest, "Malformed endpoint")
		return
	}
	cards, err := api.controller.GetCardByName(cardName)
	if err != nil {
		sendErrorMessage(w, http.StatusNotFound, "Card not found")
		return
	}
	send(w, http.StatusOK, cards)
	return

}

func (api *Api) addDeck(w http.ResponseWriter, r *http.Request) {
	log.Print("[addDeck] trying add deck")
	deck := model.Deck{}

	token, ok := r.Context().Value("user").(model.Token)
	if !ok || token == (model.Token{}) {
		log.Print("[getUser] wrong user")
		sendErrorMessage(w, http.StatusBadRequest, "Wrong token")
		return
	}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&deck)
	if err != nil || reflect.DeepEqual(deck, model.Deck{}) {
		log.Print("[addDeck] invalid token or body request")
		sendErrorMessage(w, http.StatusInternalServerError, "Invalid request - Invalid Credentials")
		return
	}
	defer r.Body.Close()
	deckName, err := api.controller.CreateDeck(deck, token.UserID)
	if err != nil {
		log.Print("[addDeck] problems saving deck")
		if err.Error() == "A commander must be a legendary card" {
			sendErrorMessage(w, http.StatusInternalServerError, err.Error())
			return
		}
		sendErrorMessage(w, http.StatusInternalServerError, "problems saving deck")
		return
	}
	send(w, http.StatusOK, deckName)
	return

}

func (api *Api) deleteUser(w http.ResponseWriter, r *http.Request) {
	methodName := "deleteUser"
	log.Printf("[%v] trying delete user", methodName)
	statusCode, userID, err := api.validateTokenUser(methodName, r)
	if err != nil {
		sendErrorMessage(w, statusCode, err.Error())
		return
	}
	err = api.controller.DeleteUser(uint(userID))
	if err != nil {
		log.Printf("[%v] user not found", methodName)
		sendErrorMessage(w, http.StatusNotFound, "User not found")
		return
	}
	log.Printf("[%v] Delete user ok", methodName)
	send(w, http.StatusOK, "User deleted")
	return

}

func (api *Api) updateUser(w http.ResponseWriter, r *http.Request) {
	methodName := "updateUser"
	log.Printf("[%v] trying update user", methodName)

	statusCode, userID, err := api.validateTokenUser(methodName, r)
	if err != nil {
		sendErrorMessage(w, statusCode, err.Error())
		return
	}
	var user model.User

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&user)
	if err != nil || (model.User{}) == user {
		sendErrorMessage(w, http.StatusInternalServerError, "Invalid request - Invalid Credentials")
	}
	defer r.Body.Close()
	user.ID = uint(userID)
	err = api.controller.UpdateUser(user)
	if err != nil {
		sendErrorMessage(w, http.StatusInternalServerError, err.Error())
		return
	}
	log.Printf("[%v] Update user ok", methodName)
	send(w, http.StatusOK, "User updated successfully")
	return

}

func (api *Api) validateTokenUser(method string, r *http.Request) (int, int, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		log.Printf("[%v] no id", method)
		return http.StatusBadRequest, 0, errors.New("Malformed endpoint")
	}
	token, ok := r.Context().Value("user").(model.Token)
	if !ok || token == (model.Token{}) {
		log.Printf("[%v] wrong user", method)
		return http.StatusBadRequest, int(token.UserID), errors.New("Wrong token")
	}
	tokenIdString := fmt.Sprintf("%v", token.UserID)
	log.Printf("[%v] TOKEN user", tokenIdString)
	if id != tokenIdString {
		log.Printf("[%v] wrong user", method)
		return http.StatusForbidden, int(token.UserID), errors.New("field not allowed")
	}
	return http.StatusOK, int(token.UserID), nil

}

func send(w http.ResponseWriter, code int, val interface{}) {
	w.Header().Set("Access-Control-Allow-Methods", "OPTIONS, TRACE, GET, HEAD, POST, PATCH")
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
