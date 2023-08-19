package userusecase

import (
	"githut.com/warley-juneo/bexs-bank-challenger/core/domain"
	"githut.com/warley-juneo/bexs-bank-challenger/core/dto"
)

func (usecase usecase) Create(userRequest *dto.CreateUserRequest) (*domain.User, error) {
	user, err := usecase.repository.Create(userRequest)

	if err != nil {
		return nil, err
	}

	return user, nil
}