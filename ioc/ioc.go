package ioc

import (
	"github.com/ariefsam/go-chat/configuration"
	"github.com/ariefsam/go-chat/implementation"
	"github.com/ariefsam/go-chat/repository"
	"github.com/ariefsam/go-chat/sms_sender"
	"github.com/ariefsam/go-chat/token_user"
	"github.com/ariefsam/go-chat/usecase"
	"github.com/jinzhu/copier"
)

func Usecase() usecase.Usecase {
	var u usecase.Usecase
	chat := repository.Chat{}
	copier.Copy(&chat, &configuration.Config.MySQL)
	u.ChatRepository = &chat
	chat.AutoMigrate()

	channelRepository := repository.Channel{}
	copier.Copy(&channelRepository, &configuration.Config.MySQL)
	u.ChannelRepository = &channelRepository
	channelRepository.AutoMigrate()

	u.IDGenerator = &implementation.IDGenerator{}

	loginVerificationRepository := repository.LoginVerification{}
	copier.Copy(&loginVerificationRepository, &configuration.Config.MySQL)
	u.LoginVerificationRepository = &loginVerificationRepository
	loginVerificationRepository.AutoMigrate()

	u.SMSSender = &sms_sender.SMSSender{}
	u.Timer = &implementation.Timer{}
	tokenService := token_user.Token{
		Secret: configuration.Config.JWTSecret,
	}
	u.TokenUserService = &tokenService

	return u
}
