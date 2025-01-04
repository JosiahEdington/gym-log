package app

import (
	"net/http"

	"github.com/JosiahEdington/gym-log/logs"
)

type Route struct {
	mux    *http.ServeMux
	logger *logs.Logger
	config *Config
}

func addRoutes(
	mux *http.ServeMux,
	logger *logs.Logger,
	config *Config,
) {
	mux.Handle("/", http.HandlerFunc(handleRootFunc))
}

func handleRootFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}
