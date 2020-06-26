package usecase

import "github.com/ariefsam/go-chat/entity"

func (u *Usecase) CreateChannel(channel entity.Channel) (savedChannel entity.Channel, err error) {
	if channel.ID == "" {
		channel.ID = u.IDGenerator.Generate()
	}

	err = u.ChannelRepository.Save(channel)
	if err != nil {
		return
	}

	savedChannel = channel
	return
}
