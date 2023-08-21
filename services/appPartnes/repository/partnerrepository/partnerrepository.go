package partnerrepository

import (
	"context"

	"github.com/wjuneo/bexs/entity"
	"github.com/wjuneo/bexs/postgres"
)

type PartnerRepository interface {
	SavePartners(ctx context.Context, partner entity.Partner) (*entity.Partner, error)
	FindPartnerByID(ctx context.Context, partner_id string) (*entity.Partner, error)
	FindPartnerByDocument(ctx context.Context, document string) (*entity.Partner, error)
	GetPartners(ctx context.Context) ([]entity.Partner, error)
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
		&savedPartner.Created_at,
		&savedPartner.Updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &savedPartner, nil
}

func (repo *partnerRepository) FindPartnerByDocument(ctx context.Context, document string) (*entity.Partner, error) {
	partner := entity.Partner{}

	err := repo.db.QueryRow(
		ctx,
		"SELECT * FROM partner WHERE document = $1",
		document,
	).Scan(
		&partner.ID,
		&partner.Trading_name,
		&partner.Document,
		&partner.Currency,
		&partner.Created_at,
		&partner.Updated_at,
	)

	if err != nil {
		return nil, nil
	}
	return &partner, nil
}

func (repo *partnerRepository) FindPartnerByID(ctx context.Context, partner_id string) (*entity.Partner, error) {
	partner := entity.Partner{}

	err := repo.db.QueryRow(
		ctx,
		"SELECT * FROM partner WHERE id = $1",
		partner_id,
	).Scan(
		&partner.ID,
		&partner.Trading_name,
		&partner.Document,
		&partner.Currency,
		&partner.Created_at,
		&partner.Updated_at,
	)

	if err != nil {
		return nil, nil
	}
	return &partner, nil
}

func (repo *partnerRepository) GetPartners(ctx context.Context) ([]entity.Partner, error) {
	partners := []entity.Partner{}

	rows, err := repo.db.Query(
		ctx,
		"SELECT * FROM partner",
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		partner := entity.Partner{}
		err := rows.Scan(
			&partner.ID,
			&partner.Trading_name,
			&partner.Document,
			&partner.Currency,
			&partner.Created_at,
			&partner.Updated_at,
		)

		if err != nil {
			return nil, err
		}
		partners = append(partners, partner)
	}

	return partners, nil
}
