package partnerrepository

import (
	"context"

	"github.com/wjuneo/bexs/entity"
	"github.com/wjuneo/bexs/postgres"
)

type PartnerRepository interface {
	SavePartners(ctx context.Context, partner entity.Partner) (*entity.Partner, error)
	FindPartnerByDocument(ctx context.Context, document string) (*entity.Partner, error)
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

func (repo *partnerRepository) FindPartnerByDocument(ctx context.Context, document string) (*entity.Partner, error) {
	var partner entity.Partner

	err := repo.db.QueryRow(
		ctx,
		"SELECT * FROM partner WHERE document = $1",
		document,
	).Scan(
		&partner.ID,
		&partner.Trading_name,
		&partner.Document,
		&partner.Currency,
	)

	if err != nil {
		return nil, nil
	}
	return &partner, nil
}