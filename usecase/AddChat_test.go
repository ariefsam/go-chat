package usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/usecase"
	"github.com/ariefsam/go-chat/usecase/dependency/mockdependency"
)

func TestAddChat(t *testing.T) {
	var u usecase.Usecase
	generatedID := "idtosave"
	mockID := mockdependency.IDGenerator{}
	u.IDGenerator = &mockID
	mockID.On("Generate").Return(generatedID)
	inputChat := entity.Chat{
		SenderID:   "xxx",
		Timestamp:  34400,
		ChatType:   "group",
		ReceiverID: "groupid",
		Message:    "Testing Message",
	}
	expectedChat := inputChat
	expectedChat.ID = generatedID
	mockChatRepository := mockdependency.ChatRepository{}
	u.ChatRepository = &mockChatRepository
	mockChatRepository.On("Save", expectedChat).Return(nil)
	savedChat, err := u.AddChat(inputChat)
	mockChatRepository.AssertCalled(t, "Save", expectedChat)
	assert.NoError(t, err)
	assert.Equal(t, expectedChat, savedChat)
}
