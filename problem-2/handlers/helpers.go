package handlers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	log "github.com/aditya-nagare/take-home-test/problem-2/log"
	"github.com/aditya-nagare/take-home-test/problem-2/models"
)

//RespondWithJSON returns json with success response
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {

	response, err := json.Marshal(payload)
	if nil != err {
		RespondWithError(w, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
	return
}

//RespondWithError returns with errors
func RespondWithError(w http.ResponseWriter, code int) {
	w.WriteHeader(code)
	return
}

func getUserEmail(ctx context.Context, serverURL, userID string) (userEmail *models.ReqData, err error) {

	log.Logger(ctx).Info("get user email request initiated")

	resp, err := http.Get(serverURL + userID)
	if nil != err {
		log.Logger(ctx).Error(err)
		return nil, err
	}

	log.Logger(ctx).Info("get user email response received")

	body, err := ioutil.ReadAll(resp.Body)
	if nil != err {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	defer resp.Body.Close()

	json.Unmarshal([]byte(body), &userEmail)
	if nil != err {
		log.Logger(ctx).Error(err)
		return nil, err
	}
	return
}
