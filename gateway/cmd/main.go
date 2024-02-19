package main

import (
	"gateway/internal/router"
	"log"
	"net/http"
)

func main() {
	r := router.Route()
	log.Println("gateway server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("failed to start gateway: %v", err)
	}
}
