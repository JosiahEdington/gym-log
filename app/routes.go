package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JosiahEdington/gym-log/data"
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
	mux.Handle("/user", http.HandlerFunc(handleUserSearch))
}

func handleRootFunc(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	greeting := "Hello there!"
	encode(w, r, 200, greeting)
}

func handleUserSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n%v new user search", time.Now().Format(time.DateTime))
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	db := data.GetGymDB()
	users, _ := db.GetAllUsers()
	// user, _ := db.GetUserByFirstName(r.FormValue("firstname"))
	encode(w, r, 200, users)
}
