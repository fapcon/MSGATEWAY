package main

import (
	"fmt"
	"net/http"
	"portfolio/internal/controller"
	"portfolio/internal/router"
	"portfolio/internal/service"
)

func main() {
	ps := service.NewPortfService()
	cnt := controller.NewHandleUser(ps)
	r := router.Route(cnt)
	fmt.Println("server portfolioService started on :8082")
	http.ListenAndServe(":8082", r)
}
