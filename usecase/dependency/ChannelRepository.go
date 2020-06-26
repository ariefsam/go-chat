package dependency

import "github.com/ariefsam/go-chat/entity"

type ChannelRepository interface {
	Get(filter entity.FilterChannel) (channels []entity.Channel)
	Save(channel entity.Channel) error
}
