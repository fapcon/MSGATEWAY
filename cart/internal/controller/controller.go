package controller

import (
	"cart/internal/service"
	"encoding/json"
	"log"
	"net/http"
)

type HandleUser struct {
	cartservice *service.CartService
}

func NewHandleUser(cs *service.CartService) *HandleUser {
	return &HandleUser{cartservice: cs}
}

func (h *HandleUser) Create(w http.ResponseWriter, r *http.Request) {
	userid := "1234"
	prods := make(map[string]int)
	mess, err := h.cartservice.Create(userid, prods)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("created")
	w.Write([]byte(mess))

}

func (h *HandleUser) Get(w http.ResponseWriter, r *http.Request) {
	userid := "1234"
	user, err := h.cartservice.Get(userid)
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
