package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

const INVLAID_TIMEZONE string = "INVALID_TIMEZONE"

func getTime(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	timezones := r.URL.Query().Get("tz")
	timezoneList := strings.Split(timezones, ",")
	response := make(map[string]string)
	defaultTimezone := "UTC"
	for _, timezone := range timezoneList {
		timezone = strings.TrimSpace(timezone)
		if timezone == "" {
			timezone = defaultTimezone
		}
		loc, err := time.LoadLocation(timezone)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "invalid timezone")
			return
		}
		response[timezone] = time.Now().In(loc).String()
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		fmt.Fprintf(w, "Something went wrong")
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}
