package models

import "time"

type (
	Chat struct {
		ChatId     int        `json:"chat_id,omitempty"`
		RoomId     string     `json:"room_id"`
		SenderId   int        `json:"sender_id"`
		ReceiverId int        `json:"receiver_id"`
		Message    string     `json:"message"`
		CreatedAt  *time.Time `json:"created_at,omitempty"`
		UpdatedAt  *time.Time `json:"updated_at,omitempty"`
	}

	Chats []Chat
)

func (m *Chat) ToLocal() {
	if m.CreatedAt != nil {
		*m.CreatedAt = m.CreatedAt.Local()
	}

	if m.UpdatedAt != nil {
		*m.UpdatedAt = m.UpdatedAt.Local()
	}
}
