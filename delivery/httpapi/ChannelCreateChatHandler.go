package httpapi

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/ioc"
)

func ChannelCreateChatHandler(w http.ResponseWriter, r *http.Request) {
	var post map[string]string
	response := map[string]interface{}{}
	json.NewDecoder(r.Body).Decode(&post)
	var channelID, token, text string
	var ok bool
	usecase := ioc.Usecase()

	if token, ok = post["token"]; !ok {
		response["error"] = "Parameter 'token' is needed"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	isValid, user := usecase.ParseToken(token)
	if !isValid {
		response["error"] = "Invalid token"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	if channelID, ok = post["channelID"]; !ok {
		response["error"] = "Parameter 'channelID' is needed"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	if text, ok = post["text"]; !ok {
		response["error"] = "Parameter 'text' is needed"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	filter := entity.FilterChannel{}
	filter.ID = &channelID

	channels := usecase.GetChannel(filter)
	if len(channels) > 0 {
		response["channel"] = channels[0]
	} else {
		response["error"] = "Channel not found"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	channel := channels[0]

	var chat entity.Chat
	chat.ReceiverID = channel.ID
	chat.ChatType = "channel"
	chat.SenderID = user.ID
	chat.SenderName = user.Name
	chat.Timestamp = time.Now().Unix()
	chat.Message = text
	savedChat, err := usecase.AddChat(chat)
	if err != nil {
		response["error"] = "Bad gateway"
		JSONView(w, response, http.StatusBadGateway)
		return
	} else {
		response["chat"] = savedChat
	}
	JSONView(w, response, http.StatusOK)
	return
}
