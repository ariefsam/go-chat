package usecase

import "github.com/ariefsam/go-chat/usecase/dependency"

type Usecase struct {
	ChatRepository dependency.ChatRepository
	UserRepository dependency.UserRepository
	IDGenerator    dependency.IDGenerator
	SMSSender      dependency.SMSSender
}
