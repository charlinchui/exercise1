package main

import (
	"charlinchui/SomeTimeInc/internal/router"
	"net/http"
)

func main() {
	http.ListenAndServe(":8080", router.Router())
}
