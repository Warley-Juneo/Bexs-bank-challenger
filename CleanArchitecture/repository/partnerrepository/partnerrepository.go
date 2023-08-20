package partnerrepository

import (
	"context"

	"github.com/wjuneo/bexs/postgres"
	"github.com/wjuneo/bexs/entity"
)

type PartnerRepository interface {
	SavePartners(ctx context.Context, partner entity.Partner) (*entity.Partner, error)
}

type partnerRepository struct {
	db postgres.PoolInterface
}

func NewPartnerRepository(conn postgres.PoolInterface) PartnerRepository {
	return &partnerRepository{
		db: conn,
	}
}

func (repo *partnerRepository) SavePartners(ctx context.Context, partner entity.Partner) (*entity.Partner, error) {
	savedPartner := entity.Partner{}

	err := repo.db.QueryRow(
		ctx,
		"INSERT INTO partner (trading_name, document, currency) VALUES ($1, $2, $3) returning *",
		partner.Trading_name,
		partner.Document,
		partner.Currency,
	).Scan(
		&savedPartner.ID,
		&savedPartner.Trading_name,
		&savedPartner.Document,
		&savedPartner.Currency,
	)

	if err != nil {
		return nil, err
	}

	return &savedPartner, nil
}