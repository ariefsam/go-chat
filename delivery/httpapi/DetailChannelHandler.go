package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/ioc"
)

func DetailChannelHandler(w http.ResponseWriter, r *http.Request) {
	var post map[string]string
	response := map[string]interface{}{}
	json.NewDecoder(r.Body).Decode(&post)
	var channelID, token string
	var ok bool
	usecase := ioc.Usecase()

	if token, ok = post["token"]; !ok {
		response["error"] = "Parameter 'token' is needed"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	isValid, _ := usecase.ParseToken(token)
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

	filter := entity.FilterChannel{}
	filter.ID = &channelID

	channels := usecase.GetChannel(filter)
	if len(channels) > 0 {
		response["channel"] = channels[0]
	}
	filterChat := entity.FilterChat{
		ReceiverID: &channelID,
	}
	chats := usecase.GetChat(filterChat)
	response["chats"] = chats
	JSONView(w, response, http.StatusOK)
	return
}
