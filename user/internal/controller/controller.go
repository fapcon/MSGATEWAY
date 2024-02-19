package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"user/internal/service"
)

type HandleUser struct {
	userservice *service.UserService
}

func NewHandleUser(us *service.UserService) *HandleUser {
	return &HandleUser{userservice: us}
}

func (h *HandleUser) Create(w http.ResponseWriter, r *http.Request) {
	id := "1234"
	name := "qwer"
	email := "asdf"
	mess, err := h.userservice.Create(id, name, email)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("created")
	w.Write([]byte(mess))

}

func (h *HandleUser) Get(w http.ResponseWriter, r *http.Request) {
	id := "1234"
	user, err := h.userservice.Get(id)
	if err != nil {
		log.Println("err", err)
	}
	jdata, err := json.Marshal(user)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("got")
	w.Write(jdata)
}
