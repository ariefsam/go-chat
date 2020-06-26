package repository

import (
	"testing"

	"github.com/jinzhu/copier"

	"github.com/stretchr/testify/assert"

	"github.com/ariefsam/go-chat/configuration"
	"github.com/ariefsam/go-chat/entity"
)

func TestChatRepository(t *testing.T) {
	var chatRepository Chat

	copier.Copy(&chatRepository, &configuration.Config.MySQL)
	chatRepository.Flush()
	chatRepository.AutoMigrate()

	t.Run("Flush data", func(t *testing.T) {
		err := chatRepository.Flush()
		assert.NoError(t, err)
	})

	chat := entity.Chat{
		ID:         "xi123",
		SenderID:   "xxx1",
		Timestamp:  34400,
		ChatType:   "group",
		ReceiverID: "groupid",
		Message:    "Testing Message",
	}

	t.Run("Save first data", func(t *testing.T) {
		err := chatRepository.Save(chat)
		assert.NoError(t, err)
	})

	t.Run("Get first data", func(t *testing.T) {
		expectedChat := []entity.Chat{
			chat,
		}
		filter := entity.FilterChat{ReceiverID: &chat.ReceiverID}
		getChat := chatRepository.Get(filter)
		assert.Equal(t, expectedChat, getChat)
	})

	t.Run("Edit First Data", func(t *testing.T) {
		chat.Message = "Testing edit message"
		err := chatRepository.Save(chat)
		assert.NoError(t, err)
	})

	t.Run("Get edited First Data", func(t *testing.T) {
		expectedChat := []entity.Chat{
			chat,
		}
		filter := entity.FilterChat{ReceiverID: &chat.ReceiverID}
		getChat := chatRepository.Get(filter)
		assert.Equal(t, expectedChat, getChat)
	})

	t.Run("Multiple chats", func(t *testing.T) {
		chats := []entity.Chat{
			entity.Chat{
				ID:         "xi001",
				SenderID:   "xxx1",
				Timestamp:  34400,
				ChatType:   "group",
				ReceiverID: "groupid",
				Message:    "Testing Message",
			},
			entity.Chat{
				ID:         "xi002",
				SenderID:   "xxx1",
				Timestamp:  34400,
				ChatType:   "group",
				ReceiverID: "groupid",
				Message:    "Testing Message",
			},
		}

		chats2 := []entity.Chat{
			entity.Chat{
				ID:         "xi003",
				SenderID:   "xxx1",
				Timestamp:  34200,
				ChatType:   "group",
				ReceiverID: "groupid2",
				Message:    "Testing Message 1",
			},
			entity.Chat{
				ID:         "xi004",
				SenderID:   "xxx1",
				Timestamp:  34400,
				ChatType:   "group",
				ReceiverID: "groupid2",
				Message:    "Testing Message 2",
			},
			entity.Chat{
				ID:         "xi005",
				SenderID:   "xxx1",
				Timestamp:  34500,
				ChatType:   "group",
				ReceiverID: "groupid2",
				Message:    "Testing Message 3",
			},
		}

		for _, val := range chats {
			chatRepository.Save(val)
		}
		for _, val := range chats2 {
			chatRepository.Save(val)
		}

		t.Run("Testing get chat id in group", func(t *testing.T) {
			filter := entity.FilterChat{ReceiverID: &chats2[0].ReceiverID}
			getChat := chatRepository.Get(filter)
			assert.Equal(t, chats2, getChat)
		})

		t.Run("Testing get chat id in group with before id", func(t *testing.T) {
			filter := entity.FilterChat{ReceiverID: &chats2[0].ReceiverID, BeforeID: &chats2[2].ID}
			getChat := chatRepository.Get(filter)
			assert.Equal(t, chats2[:len(chats2)-1], getChat)
		})

		t.Run("Testing get chat id in group with before id limit 1", func(t *testing.T) {
			var limit int
			limit = 1
			filter := entity.FilterChat{ReceiverID: &chats2[0].ReceiverID, BeforeID: &chats2[2].ID, Limit: &limit}
			getChat := chatRepository.Get(filter)
			assert.Equal(t, chats2[1:len(chats2)-1], getChat)
		})
	})

}
