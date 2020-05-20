package models

import "time"

// Message containt
type Message struct {
	UID       string    `json:"uid"`
	To        string    `json:"to_recipient"`
	Text      string    `json:"text"`
	Processed bool      `json:"is_processed"`
	CreatedAt time.Time `json:"created_at"`
}
