package userrepository

import (
	"context"

	"githut.com/warley-juneo/bexs-bank-challenger/core/domain"
	"githut.com/warley-juneo/bexs-bank-challenger/core/dto"
)

func (repository repository) Create(userRequest *dto.CreateUserRequest) (*domain.User, error) {
	ctx := context.Background()
	user := domain.User{}

	err := repository.db.QueryRow (
		ctx,
		"INSERT INTO product (trading_name, document, currency) VALUES ($1, $2, $3) returning *",
		userRequest.Trading_name,
		userRequest.Document,
		userRequest.Currency,
	).Scan (
		&user.ID,
		&user.Trading_name,
		&user.Document,
		&user.Currency,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}