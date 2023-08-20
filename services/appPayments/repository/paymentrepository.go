package repository

import (
	"context"

	"github.com/payment/entity"
	"github.com/payment/postgres"
)

type PaymentRepository interface {
	SavePayment(ctx context.Context, payment entity.Payment) (*entity.Payment, error)
	SaveConsumer(ctx context.Context, consumer entity.Consumer) (*entity.Consumer, error)
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

func (repo *paymentRepository) SaveConsumer(ctx context.Context, consumer entity.Consumer) (*entity.Consumer, error) {
	savedConsumer := entity.Consumer{}

	err := repo.db.QueryRow(
		ctx,
		"INSERT INTO consumer (name, national_id) VALUES ($1, $2) returning *",
		consumer.Name,
		consumer.National_id,
	).Scan(
		&savedConsumer.ID,
		&savedConsumer.Name,
		&savedConsumer.National_id,
		&savedConsumer.Created_at,
		&savedConsumer.Updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &savedConsumer, nil
}

func (repo *paymentRepository) SavePayment(ctx context.Context, payment entity.Payment) (*entity.Payment, error) {

	savedPayment := entity.Payment{}

	err := repo.db.QueryRow(
		ctx,
		"INSERT INTO payment (partner_id, amount, consumer_id) VALUES ($1, $2, $3) returning *",
		payment.Partner_id,
		payment.Amount,
		payment.Consumer_id,
	).Scan(
		&savedPayment.ID,
		&savedPayment.Partner_id,
		&savedPayment.Amount,
		&savedPayment.Consumer_id,
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
		if err.Error() == "sql: no rows in result set" {
			return nil, nil
		}
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
		&payment.Consumer_id,
		&payment.Created_at,
		&payment.Updated_at,
	)

	if err != nil {
		return nil, err
	}

	return &payment, nil
}
