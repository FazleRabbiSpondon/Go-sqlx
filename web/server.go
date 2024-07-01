package web

import (
	"log"
	"net/http"
)

func Start_Server() *http.ServeMux {
	log.Println("Server started on :8080")
	return Initialize_routes()
}
