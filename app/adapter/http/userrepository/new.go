package userrepository

import (
	"githut.com/warley-juneo/bexs-bank-challenger/adapter/postgres"
	"githut.com/warley-juneo/bexs-bank-challenger/core/domain"
)

type repository struct {
	db postgres.PoolInterface
}

// New returns contract implementation of UserRepository
func New(db postgres.PoolInterface) domain.UserRepository {
	return &repository {
		db: db,
	}
}