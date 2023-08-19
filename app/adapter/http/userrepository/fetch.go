package userrepository

import (
	"context"

	"githut.com/warley-juneo/bexs-bank-challenger/core/domain"
	"githut.com/warley-juneo/bexs-bank-challenger/core/dto"
	"github.com/booscaaa/go-paginate/paginate"
)

func (repository repository) Fetch(pagination *dto.PaginationRequestParms) (*domain.Pagination[[]domain.User], error) {
	ctx := context.Background()
	users := []domain.User{}
	total := int32(0)

	pagin := paginate.Instance(pagination)
	query, queryCount := pagin.
	Query("SELECT * FROM user").
	Sort(pagination.Sort).
	Page(pagination.Page).
	Desc(pagination.Descending).
	RowsPerPage(pagination.ItemsPerPage).
	SearchBy(pagination.Search, "trading_name", "document", "currency").
	Select()
	
	{
		rows, err := repository.db.Query(
			ctx,
			*query,
		)

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			user := domain.User{}

			rows.Scan(
				&user.ID,
				&user.Trading_name,
				&user.Document,
				&user.Currency,
			)

			users = append(users, user)
			
		}
	}

	{
		err := repository.db.QueryRow(ctx, *queryCount).Scan(&total)

		if err != nil {
			return nil, err
		}
	}

	return &domain.Pagination[[]domain.User] {
		Objects: users,
		Total: total,
	}, nil
}

