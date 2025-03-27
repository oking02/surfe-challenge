package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/oking02/surfe-challenge/internal/datasources"
	"github.com/oking02/surfe-challenge/internal/domain"
	"github.com/oking02/surfe-challenge/internal/transport/rest/controllers/models"
)

type GetUserActionsController struct {
	actionFetcher datasources.UserActionLister
	operationID   string
}

func NewGetUserActionsController(actionFetcher datasources.UserActionLister) *GetUserActionsController {
	return &GetUserActionsController{
		actionFetcher: actionFetcher,
		operationID:   "v1_user_actions_get",
	}
}

func (ctrl *GetUserActionsController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("user actions request %s", r.PathValue("user_id"))

	clientID, userID, errs := ctrl.parseRequest(r)
	if len(errs) > 0 {
		writeJSONResponse(w, http.StatusBadRequest, ErrorResponse{
			OperationID: ctrl.operationID,
			Message:     "failed to parse request",
			Errors:      errs,
		})

		return
	}

	actions, err := ctrl.actionFetcher.ListUserActions(r.Context(), clientID, userID)
	if err != nil {
		var (
			status = http.StatusInternalServerError
			msg    = "failed to fetch user actions"
		)
		if errors.Is(err, domain.ErrUserNotFound) {
			status = http.StatusNotFound
			msg = "user not found"
		}

		writeJSONResponse(w, status, ErrorResponse{
			OperationID: ctrl.operationID,
			Message:     msg,
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, Response{
		OperationID: ctrl.operationID,
		Data:        models.FromDomainActions(actions),
		Metadata: map[string]any{
			"count":   len(actions),
			"user_id": userID,
		},
	})
}

func (ctrl *GetUserActionsController) parseRequest(r *http.Request) (domain.ClientID, domain.UserID, map[string][]string) {

	var (
		errs     = make(map[string][]string)
		userID   domain.UserID
		clientID domain.ClientID
		err      error
	)

	clientID, err = getClientID(r)
	if err != nil {
		errs["header.client_id"] = []string{err.Error()}
	}

	userID, err = getUserID(r)
	if err != nil {
		errs["header.user_id"] = []string{err.Error()}
	}

	return clientID, userID, errs

}
