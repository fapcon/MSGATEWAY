package main

import (
	"cart/internal/controller"
	"cart/internal/router"
	"cart/internal/service"
	"fmt"
	"net/http"
)

func main() {
	cs := service.NewCartService()
	cnt := controller.NewHandleUser(cs)
	r := router.Route(cnt)
	fmt.Println("server cartService started on :8083")
	http.ListenAndServe(":8083", r)
}
