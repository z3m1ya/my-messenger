package entities

import "time"

type Message struct {
	ID       string    `json:"id"`
	ChatID   string    `json:"chatId"`
	SenderID string    `json:"senderId"`
	Content  string    `json:"content"`
	SentTime time.Time `json:"sentTime"`
}
