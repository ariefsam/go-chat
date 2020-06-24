package usecase

import "github.com/ariefsam/go-chat/entity"

func (u *Usecase) AddChat(chat entity.Chat) (savedChat entity.Chat, err error) {
	if chat.ID == "" {
		chat.ID = u.IDGenerator.Generate()
	}
	err = u.ChatRepository.Save(chat)
	if err != nil {
		return
	} else {
		savedChat = chat
	}
	return
}
