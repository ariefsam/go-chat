package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/ariefsam/go-chat/ioc"
	"github.com/ariefsam/go-chat/usecase"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var post map[string]string
	response := map[string]interface{}{}
	json.NewDecoder(r.Body).Decode(&post)

	var phoneNumber, deviceID string
	var ok bool
	if phoneNumber, ok = post["phoneNumber"]; !ok {
		response["error"] = "Phone number needed"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	if deviceID, ok = post["deviceID"]; !ok {
		response["error"] = "Device id needed"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	var u usecase.Usecase
	u = ioc.Usecase()
	err := u.LoginBySMS(phoneNumber, deviceID)
	if err != nil {
		response["error"] = "Unable to login"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	response["message"] = "SMS Sent, Please check your inbox."
	response["status"] = "ok"
	JSONView(w, response, http.StatusOK)

}
