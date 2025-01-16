package handler

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JosiahEdington/gym-log/app"
	"github.com/JosiahEdington/gym-log/data"
	"github.com/JosiahEdington/gym-log/logs"
)

type Route struct {
	mux    *http.ServeMux
	logger *logs.Logger
	config *app.Config
}

func addRoutes(
	mux *http.ServeMux,
	logger *logs.Logger,
	config *app.Config,
) {
	mux.Handle("/", http.HandlerFunc(handleRootFunc))
	mux.Handle("/user", http.HandlerFunc(handleUserSearch))
	mux.Handle("/user/new", http.HandlerFunc(handleNewUser))
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
	fmt.Printf("\n%v new user search\n", time.Now().Format(time.DateTime))
	switch r.Method {
	case http.MethodGet:
		var (
			err    error
			result any
			db     = data.GetGymDB()
		)

		if r.FormValue("SearchBy") == "" {
			fmt.Println("Get All Users")
			result, _ = db.GetAllUsers()
		} else {
			var (
				by  = r.FormValue("SearchBy")
				val = r.FormValue("SearchValue")
			)
			fmt.Printf("Getting Users where %v = '%v'\n", by, val)
			result, err = db.GetUserBySearch(by, val)
			if err != nil {
				fmt.Printf("%v\n", err)
			}
		}
		encode(w, r, 200, result)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func handleNewUser(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n%v handling new user", time.Now().Format(time.DateTime))
}

func handleWorkoutSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n%v new workout search", time.Now().Format(time.DateTime))

}
