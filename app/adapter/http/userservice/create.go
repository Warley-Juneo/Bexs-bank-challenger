package userservice

import (
	"encoding/json"
	"net/http"


	"githut.com/warley-juneo/bexs-bank-challenger/core/dto"
)

func (service service) Create(response http.ResponseWriter, request *http.Request) {
	userRequest, err := dto.FromJSONCreateUserRequest(request.Body)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	user, err := service.usecase.Create(userRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(user)
}