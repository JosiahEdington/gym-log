package data

import (
	"fmt"
	"time"
)

type User struct {
	UserId          int
	FirstName       string
	LastName        string
	Email           string
	Username        string
	DateOfBirth     time.Time
	Sex             string
	CreatedDateTime time.Time
	CreatedBy       string
	UpdatedDateTime time.Time
	UpdatedBy       string
	IsDeleted       bool
	IsInactive      bool
	IsAdmin         bool
}

type UserDto struct {
	UserId    int    `json:"userid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}

type UserNewDto struct {
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	DateOfBirth string `json:"dateofbirth"`
	Sex         string `json:"sex"`
}

type GetUserFunc func(search string) (user User)
type GetUserListFunc func(search string) (users []User)

func (repo *GymDB) GetUserBySearch(by string, val string) ([]UserDto, error) {
	var (
		users []UserDto
		q     = fmt.Sprintf("Select UserId, FirstName, LastName, Email, Username FROM User WHERE %v = '%v'", by, val)
	)

	rows, err := db.Query(q)
	if err != nil {
		return nil, fmt.Errorf("GetUserBySearch error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var usr UserDto
		rows.Scan(
			&usr.UserId,
			&usr.FirstName,
			&usr.LastName,
			&usr.Email,
			&usr.Username,
		)
		users = append(users, usr)
	}

	return users, nil
}

func (repo *GymDB) GetAllUsers() ([]UserDto, error) {
	var users []UserDto

	rows, err := db.Query("SELECT UserId, FirstName, LastName, Email, Username FROM User")
	if err != nil {
		return nil, fmt.Errorf("GetAllUsers error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var usr UserDto
		rows.Scan(
			&usr.UserId,
			&usr.FirstName,
			&usr.LastName,
			&usr.Email,
			&usr.Username,
		)
		users = append(users, usr)
	}

	return users, nil
}
func (repo *GymDB) SaveUser(user UserNewDto) (int64, error) {
	var newID int64
	query := `INSERT INTO User (FirstName, LastName, Username, Email, DateOfBirth, Sex, CreatedDateTime, CreatedBy)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?);`
	dob, err := time.Parse(time.DateOnly, user.DateOfBirth)
	if err != nil {
		fmt.Printf("Invalid User Birthday %v: %v\n", user.DateOfBirth, err)
	}

	insert, err := db.Prepare(query)
	if err != nil {
		return 0, err
	}

	resp, err := insert.Exec(
		user.FirstName,
		user.LastName,
		user.Username,
		user.Email,
		dob,
		user.Sex,
		time.Now(),
		user.Username,
	)
	insert.Close()

	if err != nil {
		return 0, err
	}

	newID, err = resp.LastInsertId()
	if newID == 0 {
		fmt.Printf("New UserId is: %v\n", newID)
		return newID, nil
	}
	if err != nil {
		fmt.Println()
	}

	return newID, nil
}
