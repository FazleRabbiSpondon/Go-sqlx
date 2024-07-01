package web

import (
	"Go0sqlx/web/handlers"
	"net/http"
)

func Initialize_routes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("user/create", handlers.CreateUser)
	// http.HandleFunc("/update", handlers.Update)
	// http.HandleFunc("/view", handlers.View)
	// http.HandleFunc("/delete", handlers.Delete)
	return router
}
