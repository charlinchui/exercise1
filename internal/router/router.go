package router

import (
	"charlinchui/SomeTimeInc/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() http.Handler {
	router := mux.NewRouter()

	router.HandleFunc("/api/time", handlers.ReturnTime)
	return router
}
