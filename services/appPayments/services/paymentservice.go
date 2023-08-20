package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/payment/dto"
	"github.com/payment/entity"
	"github.com/payment/repository"
)

type PaymentService interface {
	HandlerRequest(w http.ResponseWriter, r *http.Request)
	SavePayments(paymentData dto.PaymentData) (*dto.PaymentResponse, error)
}

type paymentService struct {
	paymentRepository repository.PaymentRepository
}

func NewPaymentService(paymentRepository repository.PaymentRepository) PaymentService {
	return &paymentService{
		paymentRepository: paymentRepository,
	}
}

func (ps *paymentService) HandlerRequest(w http.ResponseWriter, r *http.Request) {
	var paymentData dto.PaymentData

	if err := json.NewDecoder(r.Body).Decode(&paymentData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if paymentData.Amount <= 0 || paymentData.Partner_id <= 0 || paymentData.Consumer.Name == "" || paymentData.Consumer.National_id == "" {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		return
	}

	paymentData.Amount = float64(int(paymentData.Amount*100)) / 100

	payment, err := ps.SavePayments(paymentData)
	if err != nil {
		if err.Error() == "partner not found" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}

func (ps *paymentService) SavePayments(paymentData dto.PaymentData) (*dto.PaymentResponse, error) {

	consumer, _ := ps.paymentRepository.FindConsumer(context.Background(), paymentData.Consumer.National_id)
	if consumer == nil {
		entityConsumer := entity.Consumer{
			Name:        paymentData.Consumer.Name,
			National_id: paymentData.Consumer.National_id,
		}

		var err error
		consumer, err = ps.paymentRepository.SaveConsumer(context.Background(), entityConsumer)
		if err != nil {
			return nil, fmt.Errorf(err.Error())
		}
	}

	payment := entity.Payment{
		Partner_id:  paymentData.Partner_id,
		Amount:      paymentData.Amount,
		Consumer_id: consumer.ID,
	}

	newPayment, err := ps.paymentRepository.SavePayment(context.Background(), payment)
	if err != nil {
		if err.Error() == "ERROR: insert or update on table \"payment\" violates foreign key constraint \"payment_partner_id_fkey\" (SQLSTATE 23503)" {
			return nil, fmt.Errorf("partner not found")
		}
		return nil, fmt.Errorf(err.Error())
	}

	dtoPaymentResponse := dto.PaymentResponse{
		ID:         newPayment.ID,
		Partner_id: newPayment.Partner_id,
		Amount:     newPayment.Amount,
		Consumer: dto.ConsumerData{
			ID:          consumer.ID,
			Name:        consumer.Name,
			National_id: consumer.National_id,
		},

		Created_at: newPayment.Created_at.Format("2006-01-02 15:04:05"),
		Updated_at: newPayment.Updated_at.Format("2006-01-02 15:04:05"),
	}

	return &dtoPaymentResponse, nil
}
