package data

import (
	"fmt"
	"time"
)

type SettingsTemplate struct {
	SettingId       int       `json:"settingid"`
	Type            string    `json:"type"`
	Category        string    `json:"category"`
	Name            string    `json:"name"`
	Options         string    `json:"options"`
	OptionType      string    `json:"optiontype"`
	IsDeleted       bool      `json:"isdeleted"`
	CreatedDateTime time.Time `json:"createddatetime"`
	CreatedBy       string    `json:"createdby"`
	UpdatedDateTime time.Time `json:"updateddatetime"`
	UpdatedBy       string    `json:"updatedby"`
}

type UserSettings struct {
	Id              int       `json:"id"`
	UserID          int       `json:"userid"`
	SettingsId      int       `json:"settingsid"`
	Value           string    `json:"value"`
	IsDeleted       bool      `json:"isdeleted"`
	CreatedDateTime time.Time `json:"createddatetime"`
	CreatedBy       string    `json:"createdby"`
	UpdatedDateTime time.Time `json:"updateddatetime"`
	UpdatedBy       string    `json:"updatedby"`
}

type UserSettingsDto struct {
	Id       int    `json:"id"`
	Type     string `json:"type"`
	Category string `json:"category"`
	Name     string `json:"name"`
	Value    string `json:"value"`
}

func (repo *GymDB) GetUserSettings(username string) ([]UserSettingsDto, error) {
	var (
		err      error
		settings []UserSettingsDto
		db       = gymDB.db
	)

	rows, err := db.Query(`SELECT us.SettingsId, s.Type, s.Category, s.Name, us.Value
			FROM UserSettings us
			LEFT JOIN SettingsTemplate s on us.SettingsId = s.SettingId
			LEFT JOIN User u on u.UserId = us.UserId
			WHERE u.Username = ?`, username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var s UserSettingsDto
		err := rows.Scan(
			&s.Id,
			&s.Type,
			&s.Category,
			&s.Name,
			&s.Value,
		)
		if err != nil {
			fmt.Printf("Error reading information: %v\n", err)
		}
		settings = append(settings, s)
	}
	return settings, nil

}
