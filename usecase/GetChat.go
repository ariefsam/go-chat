package usecase

import "github.com/ariefsam/go-chat/entity"

func (u *Usecase) GetChat(filter entity.FilterChat) (chats []entity.Chat) {
	chats = u.ChatRepository.Get(filter)
	return
}
