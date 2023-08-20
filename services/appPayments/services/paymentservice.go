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
	SavePayments(paymentData dto.PaymentData) (*entity.Payment, error)
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

	payment, err := ps.SavePayments(paymentData)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}

func (ps *paymentService) SavePayments(paymentData dto.PaymentData) (*entity.Payment, error) {
	
	payment := entity.Payment{
		Partner_id: paymentData.Partner_id,
		Amount:     paymentData.Amount,
		Consumer: entity.Consumer{
			Name:        paymentData.Consumer.Name,
			National_id: paymentData.Consumer.National_id,
		},
	}

	newPayment, err := ps.paymentRepository.SavePayment(context.Background(), payment)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return newPayment, nil
}