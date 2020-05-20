package models

// Chat desctibe chat
type Chat struct {
	ID      int64  `json:"chat_id"`
	UserUID string `json:"user_uid"`
}
