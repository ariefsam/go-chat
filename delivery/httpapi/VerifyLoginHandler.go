package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/ariefsam/go-chat/ioc"
	"github.com/ariefsam/go-chat/usecase"
)

func VerifyLoginHandler(w http.ResponseWriter, r *http.Request) {
	var post map[string]string
	response := map[string]interface{}{}
	json.NewDecoder(r.Body).Decode(&post)

	var phoneNumber, deviceID, verificationCode string
	var ok bool
	if phoneNumber, ok = post["phoneNumber"]; !ok {
		response["error"] = "phoneNumber is needed"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	if deviceID, ok = post["deviceID"]; !ok {
		response["error"] = "deviceID is needed"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	if verificationCode, ok = post["verificationCode"]; !ok {
		response["error"] = "Verification code is needed"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	var u usecase.Usecase
	u = ioc.Usecase()
	isValid, user := u.ValidateLogin(phoneNumber, deviceID, verificationCode)
	if !isValid {
		response["error"] = "invalid token"
		JSONView(w, response, http.StatusBadRequest)
		return
	}

	token := u.GenerateToken(user)

	response["token"] = token
	response["status"] = "ok"
	JSONView(w, response, http.StatusOK)

}
