package services

import (
	"log"

	"bappa.com/rest/db"
	"bappa.com/rest/models"
)

// var Event =[]models.Event{}
// var events = []models.Event{}// event obj er array



// func Save(e Event) error {
func SaveEvent(e *models.Event) error {
	query := `
	INSERT INTO events (name, description, location, date_time, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	prepareStatement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer prepareStatement.Close()

	result, err := prepareStatement.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]models.Event, error) {

	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []models.Event

	for rows.Next() {
		var event models.Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	log.Printf("📦 Retrieved %d events: %+v\n", len(events), events)

	return events, err
}

func GetEventById(id int64) (*models.Event, error) {
	query := "SELECT * FROM events WHERE id =?"
	row := db.DB.QueryRow(query, id)

	var event models.Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}
	log.Printf("🎯 Retrieved event with ID %d: %+v\n", id, event)
	return &event,nil
	
}

func UpdateEvent(event *models.Event) error {
	query := `
	UPDATE events
	SET 
		name=?,
		description=?,
		location=?,
		date_time=?
	WHERE id = ?
	`
	prepareStatement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer prepareStatement.Close()

	_, err = prepareStatement.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	return err
}

func DeleteEvent(event *models.Event) error {
	query := "DELETE FROM events WHERE id =?"
	prepareStatement, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer prepareStatement.Close()
	_, err = prepareStatement.Exec(event.ID)
	return err
}

