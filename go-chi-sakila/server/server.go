package server

import (
	"log"
	"net/http"
	"os"
)

func Init() {
	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	router := Router()

	log.Printf("Server starting at http://localhost:%s\n", port)

	err := http.ListenAndServe(":"+port, router)

	log.Fatal(err)
}
