package domain

import (
	"net/http"

	"githut.com/warley-juneo/bexs-bank-challenger/core/dto"
)

// User is a entity of table user database column
type User struct {
	ID						int32  `json:"id"`
	Trading_name	string `json:"trading_name"`
	Document			string `json:"document"`
	Currency			string `json:"currency"`
}

// UserService is a contract of http adapter layer
type UserService interface {
	Create(response http.ResponseWriter, request *http.Request)
	Fetch(response http.ResponseWriter, request *http.Request)
}

// UserUseCase is a contract of business rule layer
type UserUseCase interface {
	Create(userRequest *dto.CreateUserRequest) (*User, error)
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination[[]User], error)
}

// UserRepository is a contract of database connection adapter layer
type UserRepository interface {
	Create(userRequest *dto.CreateUserRequest) (*User, error)
	Fetch(paginationRequest *dto.PaginationRequestParms) (*Pagination[[]User], error)
}