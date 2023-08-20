package repository

import (
	"context"

	"github.com/payment/entity"
	"github.com/payment/postgres"
)

type PaymentRepository interface {
	SavePayment(ctx context.Context, payment entity.Payment) (*entity.Payment, error)
	FindConsumer(ctx context.Context, national_id string) (*entity.Consumer, error)
	FindPayment(ctx context.Context, payment_id int32) (*entity.Payment, error)
}

type paymentRepository struct {
	db postgres.PoolInterface
}

func NewPaymentRepository(conn postgres.PoolInterface) PaymentRepository {
	return &paymentRepository{
		db: conn,
	}
}

func (repo *paymentRepository) SavePayment(ctx context.Context, payment entity.Payment) (*entity.Payment, error) {
	savedPayment := entity.Payment{}

	err := repo.db.QueryRow(
		ctx,
		"INSERT INTO payment (partner_id, amount, consumer) VALUES ($1, $2, $3) returning *",
		payment.Partner_id,
		payment.Amount,
		payment.Consumer,
	).Scan(
		&savedPayment.ID,
		&savedPayment.Partner_id,
		&savedPayment.Amount,
		&savedPayment.Consumer,
		&savedPayment.Created_at,
		&savedPayment.Updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &savedPayment, nil
}

func (repo *paymentRepository) FindConsumer(ctx context.Context, national_id string) (*entity.Consumer, error) {
	consumer := entity.Consumer{}

	err := repo.db.QueryRow(
		ctx,
		"SELECT * FROM consumer WHERE national_id = $1",
		national_id,
	).Scan(
		&consumer.ID,
		&consumer.Name,
		&consumer.National_id,
		&consumer.Created_at,
		&consumer.Updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &consumer, nil
}

func (repo *paymentRepository) FindPayment(ctx context.Context, payment_id int32) (*entity.Payment, error) {
	payment := entity.Payment{}

	err := repo.db.QueryRow(
		ctx,
		"SELECT * FROM payment WHERE id = $1",
		payment_id,
	).Scan(
		&payment.ID,
		&payment.Partner_id,
		&payment.Amount,
		&payment.Consumer,
		&payment.Created_at,
		&payment.Updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &payment, nil
}
