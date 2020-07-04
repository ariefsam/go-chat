package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/ariefsam/go-chat/entity"
	"github.com/ariefsam/go-chat/ioc"
)

func SearchChannelHandler(w http.ResponseWriter, r *http.Request) {
	var post map[string]string
	response := map[string]interface{}{}
	json.NewDecoder(r.Body).Decode(&post)
	var channelName, token string
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

	if channelName, ok = post["name"]; !ok {
		response["error"] = "Parameter 'name' is needed"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	filter := entity.FilterChannel{}
	filter.Name = &channelName

	channels := usecase.GetChannel(filter)
	response["channels"] = channels
	JSONView(w, response, http.StatusOK)
	return
}
