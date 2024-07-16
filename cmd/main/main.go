package main

import (
	"charlinchui/SomeTimeInc/internal/router"
	"log"
	"net/http"
)

func main() {
	log.Println("SERVER STARTING ON PORT 8080")
	if err := http.ListenAndServe(":8080", router.Router()); err != nil {
		log.Fatal("ERROR STARTING SERVER")
	}
}
