package main

import (
	"charlinchui/SomeTimeInc/internal/router"
	"log"
	"net/http"
)

func main() {
	log.Println("Server starting on port 8080")
	if err := http.ListenAndServe(":8080", router.Router()); err != nil {
		log.Fatal("ERROR STARTING SERVER")
	}
}
