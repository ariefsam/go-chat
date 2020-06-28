package ioc

import (
	"github.com/ariefsam/go-chat/configuration"
	"github.com/ariefsam/go-chat/repository"
	"github.com/ariefsam/go-chat/usecase"
	"github.com/jinzhu/copier"
)

func Usecase() usecase.Usecase {
	var u usecase.Usecase
	chat := repository.Chat{}
	copier.Copy(&chat, configuration.Config.MySQL)
	u.ChatRepository = &chat
	return u
}
