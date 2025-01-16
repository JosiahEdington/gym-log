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

func (repo *GymDB) SaveUser(user UserNewDto) error {
	var newID int64
	query := `INSERT INTO User (FirstName, LastName, Username, Email, DateOfBirth, Sex, CreatedDateTime, CreatedBy)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?);`
	dob, err := time.Parse(time.DateOnly, user.DateOfBirth)
	if err != nil {
		fmt.Printf("Invalid User Birthday %v: %v\n", user.DateOfBirth, err)
	}

	insert, err := db.Prepare(query)
	if err != nil {
		return err
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
		return err
	}

	newID, err = resp.LastInsertId()
	if newID == 0 {
		fmt.Printf("New UserId is: %v\n", newID)
		return nil
	}
	if err != nil {
		fmt.Println()
	}

	return nil
}
