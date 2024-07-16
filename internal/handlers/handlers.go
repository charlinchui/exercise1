package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

func ReturnTime(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]string, 0)
	tz := r.URL.Query().Get("tz")
	timezones := strings.Split(tz, ",")

	if len(timezones) <= 1 {
		if loc, err := time.LoadLocation(tz); err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(fmt.Sprintf("That timezone is not valid %s", tz)))
		} else {
			response[loc.String()] = time.Now().In(loc).String()
		}
	} else {
		for _, zone := range timezones {
			if loc, err := time.LoadLocation(zone); err != nil {
				w.WriteHeader(http.StatusNotFound)
				w.Write([]byte(fmt.Sprintf("That timezone is not valid %s", zone)))
				return
			} else {
				response[loc.String()] = time.Now().In(loc).String()
			}
		}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
