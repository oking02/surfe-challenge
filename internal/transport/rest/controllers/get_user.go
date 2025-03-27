package controllers

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/oking02/surfe-challenge/internal/transport/rest/controllers/models"

	"github.com/oking02/surfe-challenge/internal/datasources"
	"github.com/oking02/surfe-challenge/internal/domain"
)

type GetUserController struct {
	userFetcher datasources.UserFetcher
	operationID string
}

func NewGetUserController(userFetcher datasources.UserFetcher) *GetUserController {
	return &GetUserController{
		userFetcher: userFetcher,
		operationID: "v1_user_get",
	}
}

func (ctrl *GetUserController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("user id request %s", r.PathValue("id"))

	clientID, userID, errs := ctrl.parseRequest(r)
	if len(errs) > 0 {
		writeJSONResponse(w, http.StatusBadRequest, ErrorResponse{
			OperationID: ctrl.operationID,
			Message:     "failed to parse request",
			Errors:      errs,
		})

		return
	}

	user, err := ctrl.userFetcher.GetUser(r.Context(), clientID, userID)
	if err != nil {
		var (
			status = http.StatusInternalServerError
			msg    = "failed to fetch user"
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
		Data:        models.FromDomainUser(user),
	})
}

func (ctrl *GetUserController) parseRequest(r *http.Request) (domain.ClientID, domain.UserID, map[string][]string) {

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
		errs["path.user_id"] = []string{err.Error()}
	}

	return clientID, userID, errs

}
