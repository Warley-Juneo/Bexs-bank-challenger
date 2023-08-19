package userusecase

import (
	"githut.com/warley-juneo/bexs-bank-challenger/core/domain"
	"githut.com/warley-juneo/bexs-bank-challenger/core/dto"
)

func (usecase usecase) Fetch(paginationRequest *dto.PaginationRequestParms) (*domain.Pagination[[]domain.User], error) {
	users, err := usecase.repository.Fetch(paginationRequest)

	if err != nil {
		return nil, err
	}

	return users, nil
}