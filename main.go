package main

import (
	"api/src/router"
	"log"
	"net/http"
)

func main() {
	println("API IS RUNNING...")

	r := router.Generate()

	log.Fatal(http.ListenAndServe(":5000", r))
}
