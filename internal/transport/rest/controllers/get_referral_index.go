package controllers

import (
	"fmt"
	"net/http"

	"github.com/oking02/surfe-challenge/internal/datasources"
	"github.com/oking02/surfe-challenge/internal/domain"
)

type GetReferralIndexController struct {
	referralStats datasources.ReferralStatistics
	operationID   string
}

func NewGetReferralIndexController(referralStats datasources.ReferralStatistics) *GetReferralIndexController {
	return &GetReferralIndexController{
		referralStats: referralStats,
		operationID:   "v1_action_referral_index_get",
	}
}

func (ctrl *GetReferralIndexController) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println("referral index request")

	clientID, errs := ctrl.parseRequest(r)
	if len(errs) > 0 {
		writeJSONResponse(w, http.StatusBadRequest, ErrorResponse{
			OperationID: ctrl.operationID,
			Message:     "failed to parse request",
			Errors:      errs,
		})

		return
	}

	result, err := ctrl.referralStats.ReferralIndex(r.Context(), clientID)
	if err != nil {
		writeJSONResponse(w, http.StatusInternalServerError, ErrorResponse{
			OperationID: ctrl.operationID,
			Message:     "failed to calculate referral index",
		})
		return
	}

	writeJSONResponse(w, http.StatusOK, Response{
		OperationID: ctrl.operationID,
		Data:        result,
	})
}

func (ctrl *GetReferralIndexController) parseRequest(r *http.Request) (domain.ClientID, map[string][]string) {

	var (
		errs     = make(map[string][]string)
		clientID domain.ClientID
		err      error
	)

	clientID, err = getClientID(r)
	if err != nil {
		errs["header.client_id"] = []string{err.Error()}
	}

	return clientID, errs

}
