package ioc

import (
	"github.com/ariefsam/go-chat/repository"
	"github.com/ariefsam/go-chat/usecase"
)

func Usecase() usecase.Usecase {
	var u usecase.Usecase
	u.ChatRepository = &repository.Chat{}
	return u
}
