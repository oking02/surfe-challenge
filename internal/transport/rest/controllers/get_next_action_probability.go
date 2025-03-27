package controllers

import (
	"fmt"
	"net/http"

	"github.com/oking02/surfe-challenge/internal/datasources"
	"github.com/oking02/surfe-challenge/internal/domain"
)

type GetNextActionProbabilityController struct {
	probabilityFetcher datasources.ActionProbability
	operationID        string
}

func NewGetNextActionProbabilityController(probabilityFetcher datasources.ActionProbability) *GetNextActionProbabilityController {
	return &GetNextActionProbabilityController{
		probabilityFetcher: probabilityFetcher,
		operationID:        "v1_action_probability_next_get",
	}
}

func (ctrl *GetNextActionProbabilityController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("next action probability request %s", r.PathValue("type"))

	clientID, actionType, errs := ctrl.parseRequest(r)
	if len(errs) > 0 {
		writeJSONResponse(w, http.StatusBadRequest, ErrorResponse{
			OperationID: ctrl.operationID,
			Message:     "failed to parse request",
			Errors:      errs,
		})

		return
	}

	result, err := ctrl.probabilityFetcher.NextActionProbability(r.Context(), clientID, actionType)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{
			OperationID: ctrl.operationID,
			Message:     "failed to calculate next action stats",
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, Response{
		OperationID: ctrl.operationID,
		Data:        result,
	})

}

func (ctrl *GetNextActionProbabilityController) parseRequest(r *http.Request) (domain.ClientID, domain.ActionType, map[string][]string) {

	var (
		errs     = make(map[string][]string)
		userID   domain.ActionType
		clientID domain.ClientID
		err      error
	)

	clientID, err = getClientID(r)
	if err != nil {
		errs["header.client_id"] = []string{err.Error()}
	}

	userID, err = getPathAction(r)
	if err != nil {
		errs["path.action"] = []string{err.Error()}
	}

	return clientID, userID, errs

}
