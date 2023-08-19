package userservice

import (
	"encoding/json"
	"net/http"


	"githut.com/warley-juneo/bexs-bank-challenger/core/dto"
)

func (service service) Fetch(response http.ResponseWriter, request *http.Request) {
	paginationRequest, err := dto.FromValuePaginationRequestParms(request)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	users, err := service.usecase.Fetch(paginationRequest)

	if err != nil {
		response.WriteHeader(500)
		response.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(response).Encode(users)
}