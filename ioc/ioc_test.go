package ioc_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/ioc"
)

func TestUsecase(t *testing.T) {
	u := ioc.Usecase()
	assert.NotNil(t, u.ChannelRepository)
	assert.NotNil(t, u.ChatRepository)
	assert.NotNil(t, u.IDGenerator)
	assert.NotNil(t, u.LoginVerificationRepository)
	assert.NotNil(t, u.SMSSender)
	assert.NotNil(t, u.Timer)
	assert.NotNil(t, u.TokenUserService)
	assert.NotNil(t, u.UserRepository)

}
