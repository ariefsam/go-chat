package usecase

import "github.com/ariefsam/go-chat/usecase/dependency"

type Usecase struct {
	ChatRepository              dependency.ChatRepository
	LoginVerificationRepository dependency.LoginVerificationRepository
	UserRepository              dependency.UserRepository
	IDGenerator                 dependency.IDGenerator
	SMSSender                   dependency.SMSSender
	Timer                       dependency.Timer
}
