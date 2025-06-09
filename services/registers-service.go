package services

import (
	"bappa.com/rest/db"
	"bappa.com/rest/models"
)

func RegisterEvent(userId int64, e models.Event) error {
   query := `
	INSERT INTO registrations (event_id, user_id)
	VALUES (?, ?)
	`

	prepareStatement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer prepareStatement.Close()

	_, err = prepareStatement.Exec(e.ID,userId)
	if err != nil {
		return err
	}

	return err
}
