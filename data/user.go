package data

import (
	"database/sql"
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

func (repo *GymDB) GetUserByFirstName(firstname string) (UserDto, error) {
	var usr UserDto

	row := db.QueryRow("Select * FROM User WHERE FirstName = ?", firstname)
	if err := row.Scan(&usr.UserId, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Username); err != nil {
		if err == sql.ErrNoRows {
			return usr, fmt.Errorf("GetUserByFirstName %s: no such name", firstname)
		}
		return usr, fmt.Errorf("GetUserByFirstName %s: %v", firstname, err)
	}
	return usr, nil
}

func (repo *GymDB) GetUserByUsername(username string) (UserDto, error) {
	var usr UserDto

	row := db.QueryRow("Select UserId, FirstName, LastName, Email, Username FROM User WHERE Username = ?", username)
	if err := row.Scan(&usr.UserId, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Username); err != nil {
		fmt.Println(row.Scan(&usr.UserId, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Username))
		if err == sql.ErrNoRows {
			return usr, fmt.Errorf("GetUserByUsername %s: no such username", username)
		}
		return usr, fmt.Errorf("GetUserByUsername %s: %v", username, err)
	}
	return usr, nil
}
