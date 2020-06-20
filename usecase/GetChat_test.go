package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/usecase"
	"github.com/ariefsam/go-chat/usecase/dependency/mockdependency"
)

func TestGetChat(t *testing.T) {
	var u usecase.Usecase
	filter := entity.FilterChat{}
	expectedChats := []entity.Chat{
		entity.Chat{
			ReceiverID: "groupIDxxx1",
			Timestamp:  1500,
			SenderID:   "userIDxx1",
			ChatType:   "group",
			Message:    "Yeyeye",
		},
		entity.Chat{
			ReceiverID: "groupIDxxx1",
			Timestamp:  1520,
			SenderID:   "userIDxx2",
			ChatType:   "group",
			Message:    "Yeyeye",
		},
	}
	var mockChatRepository mockdependency.ChatRepository
	u.ChatRepository = &mockChatRepository
	mockChatRepository.On("Get", filter).Return(expectedChats)
	chats := u.GetChat(filter)
	assert.Equal(t, expectedChats, chats)

}
