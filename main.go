package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	// Http Server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		slog.Info(fmt.Sprintf("Request: %s %s", r.Method, r.URL.Path))

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Hello World",
		})
	})

	// Start Server
	slog.Info("Server started at :8080")
	http.ListenAndServe(":8080", nil)
}
