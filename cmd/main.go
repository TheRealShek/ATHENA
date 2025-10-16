package main

import (
	"Athena/internal/handlers"
	"Athena/internal/hub"
	"fmt"
	"net/http"
)

func main() {
	handlers.MainHub = hub.NewHub()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "ATHENA server is running.")
	})
	http.HandleFunc("/ws", handlers.WsHandler)

	fmt.Println("Server started on :8080")
	http.ListenAndServe(":8080", nil)
}
