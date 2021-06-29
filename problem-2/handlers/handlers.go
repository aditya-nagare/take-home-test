package handlers

import (
	"net/http"

	log "github.com/aditya-nagare/take-home-test/problem-2/log"
	"github.com/aditya-nagare/take-home-test/problem-2/models"
	"github.com/gorilla/mux"
)

//Handler implements methods in handler
type Handler interface {
	GetUserEmail(http.ResponseWriter, *http.Request)
}

//HandlerImplementation returns dependencies for handler
type HandlerImplementation struct {
	serverURL string
}

//NewHandlerImplementation returns new object for handlers
func NewHandlerImplementation(serverURL string) *HandlerImplementation {
	return &HandlerImplementation{
		serverURL: serverURL,
	}
}

// Ping returns server status
func (handler HandlerImplementation) GetUserEmail(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	log.Logger(ctx).Info("get user email request")

	vars := mux.Vars(r)
	userID := vars["ID"]

	log.Logger(ctx).Info("user email request for userID:", userID)

	userEmail, err := getUserEmail(ctx, handler.serverURL, userID)
	if nil != err {
		log.Logger(ctx).Error(err)
		RespondWithError(w, http.StatusInternalServerError)
		return
	}
	var getUserEmailResp models.Response
	getUserEmailResp.Email = userEmail.Data.Email

	defer log.Logger(ctx).Info("get user email response sent to user")

	RespondWithJSON(w, http.StatusOK, getUserEmailResp)
}
