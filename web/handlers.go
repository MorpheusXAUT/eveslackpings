package web

import (
	"fmt"
	"net/http"

	"github.com/morpheusxaut/eveslackpings/misc"
	"github.com/morpheusxaut/eveslackpings/models"
)

func (controller *Controller) PingPostHandler(w http.ResponseWriter, r *http.Request) {
	response := make(map[string]interface{})
	response["pageType"] = 2
	response["pageTitle"] = "Login"

	err := r.ParseForm()
	if err != nil {
		misc.Logger.Warnf("Failed to parse form: [%v]", err)

		response["success"] = false
		response["error"] = fmt.Errorf("Failed to parse form, please try again!")

		//controller.SendResponse(w, r, "ping", response)
		http.Error(w, "Failed to parse form, please try again!", http.StatusBadRequest)

		return
	}

	target := r.Form.Get("target")
	token := r.Form.Get("token")
	username := r.Form.Get("user_name")
	text := r.Form.Get("text")

	if len(target) == 0 || len(token) == 0 || len(username) == 0 || len(text) == 0 {
		misc.Logger.Warnf("Received empty target, token, username or text")

		response["success"] = false
		response["error"] = fmt.Errorf("Empty target, token, username or text, please try again!")

		//controller.SendResponse(w, r, "ping", response)
		http.Error(w, "Empty target, token, username or text, please try again!", http.StatusBadRequest)

		return
	}

	if !misc.ValidSlackToken(controller.Config.SlackTokens, token) {
		misc.Logger.Warnf("Received invalid token")

		response["success"] = false
		response["error"] = fmt.Errorf("Invalid token, please try again!")

		//controller.SendResponse(w, r, "ping", response)
		http.Error(w, "Invalid token, please try again!", http.StatusUnauthorized)

		return
	}

	ping := models.NewPing(username, text, fmt.Sprintf("#%s", target))

	err = ping.Send(controller.Config.SlackWebhookURL)
	if err != nil {
		misc.Logger.Warnf("Failed to send ping")

		response["success"] = false
		response["error"] = fmt.Errorf("Failed to send ping, please try again!")

		//controller.SendResponse(w, r, "ping", response)
		http.Error(w, "Failed to send ping, please try again!", http.StatusInternalServerError)

		return
	}
}
