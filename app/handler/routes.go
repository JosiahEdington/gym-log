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
	mux.Handle("/workout", http.HandlerFunc(handleWorkoutSearch))
	mux.Handle("/workout/new", http.HandlerFunc(handleNewWorkout))
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
	var valid bool

	usr, err := decoder[data.NewUser](r)
	if err != nil {
		fmt.Printf("Invalid user information: %v\n", err)
	}
	db := data.GetGymDB()

	result, err := db.GetUserBySearch("Username", usr.Username)
	if err != nil {
		fmt.Printf("Error reading new user: %v\n", err)
		valid = false
	} else if result != nil {
		fmt.Printf("Username already exists: %v\n", result)
		valid = false
	} else {
		valid = true
	}

	result, err = db.GetUserBySearch("Email", usr.Email)
	if err != nil {
		fmt.Printf("Error reading new user: %v\n", err)
		valid = false
	} else if result != nil {
		fmt.Printf("Email already exists: %v\n", result)
		valid = false
	} else {
		valid = true
	}

	if valid {
		newID, err := db.SaveNewUser(usr)
		if err != nil {
			fmt.Printf("Error saving new user %v: %v", usr, err)
			w.WriteHeader(http.StatusNotModified)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Printf("User saved successfully with UserID: %v\n", newID)
		}
	}

}

func handleWorkoutSearch(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("\n%v Searching workout for user: %v\n", time.Now().Format(time.DateTime), r.Header.Get("User"))

	switch r.Method {
	case http.MethodGet:
		var (
			err    error
			result any
			db     = data.GetGymDB()
		)

		usr, err := db.GetUserByUsername(r.Header.Get("User"))
		if err != nil {
			encode(w, r, 404, "not logged in")
		}

		fmt.Printf("Searching workouts for username: %v\n", usr.Username)
		if r.FormValue("SearchBy") == "" {
			fmt.Println("Get All Workouts")
			result, _ = db.GetAllWorkouts(usr)
		} else {
			var (
				by  = r.FormValue("SearchBy")
				val = r.FormValue("SearchValue")
			)
			fmt.Printf("Searching workouts for user %v where %v = '%v'\n", usr.Username, by, val)
			result, err = db.GetWorkoutBySearch(by, val, usr)
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

func handleNewWorkout(w http.ResponseWriter, r *http.Request) {
	var (
		valid    bool
		username = r.Header.Get("User")
	)
	fmt.Printf("\n%v Saving new workout for user: %v\n", time.Now().Format(time.DateTime), username)

	wo, err := decoder[data.NewWorkoutDto](r)
	if err != nil {
		fmt.Printf("Invalid workout information: %v\n", err)
	}
	db := data.GetGymDB()

	fmt.Printf("Workout info: %v\n", wo)
	user, err := db.GetUserByUsername(username)
	fmt.Println("Result of user: ", user)
	if err != nil {
		fmt.Printf("Error reading associated user: %v\n", err)
		valid = false
	} else if user.Username == "" {
		fmt.Printf("Username is blank")
		valid = false
	} else {
		valid = true
	}

	if valid {
		newID, err := db.SaveWorkout(wo, user)
		if err != nil {
			fmt.Printf("Error saving new workout %v: %v", wo, err)
			w.WriteHeader(http.StatusNotModified)
			return
		} else {
			w.WriteHeader(http.StatusOK)
			fmt.Printf("Workout saved successfully with WorkoutId: %v\n", newID)
		}
	}
}
