package dto

import (
	"encoding/json"
	"io"
)

type CreateUserRequest struct {
	Trading_name 	string `json:"trading_name"`
	Document 			string `json:"document"`
	Currency 			string `json:"currency"`
}

func FromJSONCreateUserRequest(body io.Reader) (*CreateUserRequest, error) {
	createUserRequest := CreateUserRequest{}
	if err := json.NewDecoder(body).Decode(&createUserRequest); err != nil {
		return nil, err
	}

	return &createUserRequest, nil
}