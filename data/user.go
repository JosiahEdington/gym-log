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
	UserId    int
	FirstName string
	LastName  string
	Email     string
	Username  string
}

type GetUserFunc func(search string) (user User)
type GetUserListFunc func(search string) (users []User)

func (repo *GymDB) GetUserListByUsername(username string) ([]User, error) {
	var users []User

	fmt.Printf("Searching for username: '%v'\n", username)

	rows, err := repo.db.Query("SELECT Username FROM User WHERE Username = ?", username)
	if err != nil {
		return nil, fmt.Errorf("getUsersByUsername %q: %v", username, err)
	}

	defer rows.Close()

	for rows.Next() {
		var usr User
		fmt.Println(rows.Scan(&usr.Username))
		if err := rows.Scan(&usr.Username); err != nil {
			return nil, fmt.Errorf("getUsersByUsername %v: %v", username, err)
		}
		users = append(users, usr)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("getUsersByUsername %q: %v", username, err)
	}
	return users, nil

}

func (repo *GymDB) GetUserByFirstName(firstname string) (UserDto, error) {
	var usr UserDto
	fmt.Printf("Searching for %v\n", firstname)

	row := db.QueryRow("Select * FROM User WHERE FirstName = ?", firstname)
	if err := row.Scan(&usr.UserId, &usr.FirstName, &usr.LastName, &usr.Email, &usr.Username); err != nil {
		if err == sql.ErrNoRows {
			return usr, fmt.Errorf("GetUserByFirstName %s: no such name", firstname)
		}
		return usr, fmt.Errorf("GetUserByFirstName %s: %v", firstname, err)
	}
	return usr, nil
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
			// &usr.DateOfBirth,
			// &usr.Sex,
			// &usr.CreatedDateTime,
			// &usr.CreatedBy,
			// &usr.UpdatedDateTime,
			// &usr.UpdatedBy,
			// &usr.IsDeleted,
			// &usr.IsInactive,
			// &usr.IsAdmin,
		)
		users = append(users, usr)
	}

	return users, nil
}
