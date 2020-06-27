package usecase

import "github.com/ariefsam/go-chat/entity"

func (u *Usecase) GetChannel(filter entity.FilterChannel) (channels []entity.Channel) {
	channels = u.ChannelRepository.Get(filter)
	return
}
