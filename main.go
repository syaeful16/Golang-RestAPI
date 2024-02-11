package main

import (
	"go-restapi/router"
	"log"
	"net/http"
)

func main() {
	router.Routes()

	log.Println("running on http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
