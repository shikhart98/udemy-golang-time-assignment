package app

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Start() {
	port := os.Getenv("PORT")
	router := mux.NewRouter()

	router.HandleFunc("/api/time", getTime).Methods(http.MethodGet)
	fmt.Println("Server running at port" + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
