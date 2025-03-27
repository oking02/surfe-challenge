package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/oking02/surfe-challenge/internal/domain"
)

type Response struct {
	OperationID string `json:"operation_id"`
	Data        any    `json:"data"`
	Metadata    any    `json:"metadata,omitempty"`
}

type ErrorResponse struct {
	OperationID string `json:"operation_id"`
	Message     any    `json:"message"`
	Errors      any    `json:"errors,omitempty"`
}

func writeJSONResponse(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		fmt.Println(err)
	}
}

/*
	Assuming this would be added by an API gateway type of process
	Possibly using JWTs or other Token based auth systems
*/

const clientHeader = "X-Client-ID"

func getClientID(r *http.Request) (domain.ClientID, error) {
	clientIDInput := r.Header.Get(clientHeader)
	// nolint:staticcheck
	if clientIDInput == "" {
		// would normally error
		// this challenge is will always be ""
	}
	return domain.ClientID(clientIDInput), nil
}

func getUserID(r *http.Request) (domain.UserID, error) {

	userIDInput := r.PathValue("id")
	if userIDInput == "" {
		return 0, errors.New("missing user ID")
	}

	userIDInt, err := strconv.Atoi(userIDInput)
	if err != nil {
		return 0, errors.New("user ID must be an integer")
	}

	return domain.UserID(userIDInt), nil
}

func getPathAction(r *http.Request) (domain.ActionType, error) {

	actionTypeInput := r.PathValue("type")
	if actionTypeInput == "" {
		return "", errors.New("missing action type")
	}

	// could validate if it was a set list or enum

	return domain.ActionType(actionTypeInput), nil
}
