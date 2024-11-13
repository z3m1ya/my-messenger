package usecases

import (
	"chats/internal/repo"
	"sync"

	"chats/internal/entities"
)

type ChatUseCase struct {
	userRepo *repo.InMemoryUserRepository
	chats    map[string]*Chat
	mu       sync.RWMutex
}

type Chat struct {
	ID           string
	Participants []string
	Messages     []entities.Message
}

func NewChatUseCase(userRepo *repo.InMemoryUserRepository) *ChatUseCase {
	return &ChatUseCase{
		userRepo: userRepo,
		chats:    make(map[string]*Chat),
	}
}

func (uc *ChatUseCase) CreateChat(participants []string) *Chat {
	uc.mu.Lock()
	defer uc.mu.Unlock()
	chatID := "some chat id"
	chat := &Chat{
		ID:           chatID,
		Participants: participants,
		Messages:     []entities.Message{},
	}
	uc.chats[chatID] = chat
	return chat
}

func (uc *ChatUseCase) SendMessage(chatID string, message entities.Message) bool {
	uc.mu.Lock()
	defer uc.mu.Unlock()
	chat, exists := uc.chats[chatID]
	if !exists {
		return false
	}
	chat.Messages = append(chat.Messages, message)
	return true
}

func (uc *ChatUseCase) GetChat(chatID string) (*Chat, bool) {
	uc.mu.RLock()
	defer uc.mu.RUnlock()
	chat, exists := uc.chats[chatID]
	return chat, exists
}
