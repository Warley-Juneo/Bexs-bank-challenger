package userusecase

import "githut.com/warley-juneo/bexs-bank-challenger/core/domain"

type usecase struct {
	repository domain.UserRepository
}

//new returns contract implementation of UserUseCase
func New(repository domain.UserRepository) domain.UserUseCase {
	return &usecase {
		repository: repository,
	}
}