package data

import (
	"fmt"
	"time"
)

type Workout struct {
	WorkoutId       int       `json:"workoutid"`
	UserId          int       `json:"userid"`
	Name            string    `json:"name"`
	StartedDateTime time.Time `json:"starteddatetime"`
	EndedDateTime   time.Time `json:"endeddatetime"`
	SecondDuration  int       `json:"secondduration"`
	Category        string    `json:"category"`
	Type            string    `json:"type"`
	CreatedDateTime time.Time `json:"createddatetime"`
	CreatedBy       string    `json:"createdby"`
	UpdatedDateTime time.Time `json:"updateddatetime"`
	UpdatedBy       string    `json:"updatedby"`
	IsDeleted       bool      `json:"isdeleted"`
	Rating          int       `json:"rating"`
}

type WorkoutDto struct {
	WorkoutId       int     `json:"workoutid"`
	User            UserDto `json:"user"`
	Name            string  `json:"name"`
	StartedDateTime string  `json:"starteddatetime"`
	EndedDateTime   string  `json:"endeddatetime"`
	SecondDuration  int     `json:"secondduration"`
	Category        string  `json:"category"`
	Type            string  `json:"type"`
	Rating          int     `json:"rating"`
}

type NewWorkoutDto struct {
	Name            string `json:"name"`
	StartedDateTime string `json:"starteddatetime"`
	EndedDateTime   string `json:"endeddatetime"`
	SecondDuration  int    `json:"secondduration"`
	Category        string `json:"category"`
	Type            string `json:"type"`
	Rating          int    `json:"rating"`
}

type GetWorkoutFunc func(search string) (workout Workout)
type GetWorkoutListFunc func(search string) (workouts []Workout)

func (repo *GymDB) GetWorkoutBySearch(by string, val string, user UserDto) ([]WorkoutDto, error) {
	var (
		workouts []WorkoutDto
		q        = fmt.Sprintf("Select User, Name, StartedDateTime, EndedDateTime, SecondDuration, Category, Type, Rating FROM Workout WHERE UserId = %v and %v = %v", user.UserId, by, val)
	)

	rows, err := db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("GetWorkoutBySearch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var wo WorkoutDto
		rows.Scan(
			&wo.User,
			&wo.Name,
			&wo.StartedDateTime,
			&wo.EndedDateTime,
			&wo.SecondDuration,
			&wo.Category,
			&wo.Type,
			&wo.Rating,
		)
		workouts = append(workouts, wo)
	}

	return workouts, nil
}

func (repo *GymDB) GetUserWorkouts(user UserDto) ([]WorkoutDto, error) {
	var (
		workouts []WorkoutDto
		q        = fmt.Sprintf("Select Name, StartedDateTime, EndedDateTime, SecondDuration, Category, Type, Rating FROM Workout WHERE UserId = %v", user.UserId)
	)

	rows, err := db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("GetUserWorkouts error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var wo WorkoutDto
		wo.User = user
		rows.Scan(
			&wo.Name,
			&wo.StartedDateTime,
			&wo.EndedDateTime,
			&wo.SecondDuration,
			&wo.Category,
			&wo.Type,
			&wo.Rating,
		)
		workouts = append(workouts, wo)
	}

	return workouts, nil
}

func (repo *GymDB) SaveWorkout(workout NewWorkoutDto, user UserDto) (int64, error) {
	var newID int64
	q := `INSERT INTO Workout (UserId, Name, StartedDateTime, EndedDateTime, SecondDuration, Category, Type, Rating, CreatedDateTime, CreatedBy)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?);`

	start, err := time.Parse(time.DateTime, workout.StartedDateTime)
	if err != nil {
		fmt.Printf("Invalid StartedDateTime %v: %v\n", workout.StartedDateTime, err)
	}

	end, err := time.Parse(time.DateTime, workout.EndedDateTime)
	if err != nil {
		fmt.Printf("Invalid StartedDateTime %v: %v\n", workout.EndedDateTime, err)
	}

	insert, err := db.Prepare(q)
	if err != nil {
		return 0, err
	}

	resp, err := insert.Exec(
		user.UserId,
		workout.Name,
		start,
		end,
		workout.SecondDuration,
		workout.Category,
		workout.Type,
		workout.Rating,
		time.Now(),
		user.Username,
	)
	insert.Close()

	if err != nil {
		return 0, err
	}

	newID, err = resp.LastInsertId()
	if newID == 0 {
		fmt.Printf("New WorkoutId is: %v\n", newID)
		return newID, nil
	}
	if err != nil {
		fmt.Println(err)
	}

	return newID, nil

}
