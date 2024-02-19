package main

import (
	"fmt"
	"net/http"
	"user/internal/controller"
	"user/internal/router"
	"user/internal/service"
)

func main() {
	us := service.NewUserService()
	cnt := controller.NewHandleUser(us)
	r := router.Route(cnt)
	fmt.Println("server userService started on :8081")
	http.ListenAndServe(":8081", r)
}
