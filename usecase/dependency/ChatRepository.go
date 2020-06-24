package dependency

import "github.com/ariefsam/go-chat/entity"

type ChatRepository interface {
	Get(filter entity.FilterChat) (chats []entity.Chat)
	Save(chat entity.Chat) error
}
