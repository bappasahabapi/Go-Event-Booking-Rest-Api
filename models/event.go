package models

import (
	"log"
	"time"

	"bappa.com/rest/db"
	// "github.com/pelletier/go-toml/query"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int       `json:"user_id"`
}

var events = []Event{} // event obj er array

// func Save(e Event) error {
func (e Event) Save() error {
	query := `
	INSERT INTO events (name, description, location, date_time, user_id)
	VALUES (?, ?, ?, ?, ?)
	`
	db.DB.Prepare(query)
	// statement, err := db.DB.Exec(query, e.Name, e.Description, e.Location, e.DateTime, e.UserID)
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
	// events = append(events, e)
}

func GetAllEvents() ([]Event, error) {

	query := `SELECT * FROM events`
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	log.Printf("📦 Retrieved %d events: %+v\n", len(events), events)

	return events, err
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id =?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

	if err != nil {
		return nil, err
	}
	log.Printf("🎯 Retrieved event with ID %d: %+v\n", id, event)
	return &event,nil
	
}
