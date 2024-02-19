package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"portfolio/internal/service"
)

type HandleUser struct {
	portfservice *service.PortfService
}

func NewHandleUser(ps *service.PortfService) *HandleUser {
	return &HandleUser{portfservice: ps}
}

func (h *HandleUser) Create(w http.ResponseWriter, r *http.Request) {
	userid := "1234"
	cur := make(map[string]float64)
	mess, err := h.portfservice.Create(userid, cur)
	if err != nil {
		log.Println("err", err)
	}
	log.Println("created")
	w.Write([]byte(mess))

}

func (h *HandleUser) Get(w http.ResponseWriter, r *http.Request) {
	userid := "1234"
	user, err := h.portfservice.Get(userid)
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
