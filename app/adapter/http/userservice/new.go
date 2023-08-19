package userservice

import (
	"githut.com/warley-juneo/bexs-bank-challenger/core/domain"
)

type service struct {
	usecase domain.UserUseCase
}

// New return contract implementation of UseService
func New(usecase domain.UserUseCase) domain.UserService {
	return &service {
		usecase: usecase,
	}
}