package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/payment/logs"
	"github.com/payment/consts"
	"github.com/payment/dto"
	"github.com/payment/entity"
	"github.com/payment/repository"
)

type PaymentService interface {
	HandlerRequest(w http.ResponseWriter, r *http.Request)
	SavePayments(paymentData dto.PaymentData) (*dto.PaymentResponse, error)
	HandlerGetPaymentId(w http.ResponseWriter, r *http.Request)
	HandlerGetPayment(w http.ResponseWriter, r *http.Request)
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
		logs.LogToFile("logs/error.log", err.Error())
		return
	}

	if paymentData.Amount <= 0 || paymentData.Partner_id <= 0 || paymentData.Consumer.Name == "" || paymentData.Consumer.National_id == "" {
		http.Error(w, "Invalid data", http.StatusBadRequest)
		logs.LogToFile("logs/error.log", "Invalid data")
		return
	}

	paymentData.Amount = float64(int(paymentData.Amount*100)) / 100
	if len(paymentData.Consumer.National_id) != 11 {
		http.Error(w, "Invalid National Id", http.StatusBadRequest)
		logs.LogToFile("logs/error.log", "Invalid National Id")
		return
	}

	payment, err := ps.SavePayments(paymentData)
	if err != nil {
		if err.Error() == "partner not found" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			logs.LogToFile("logs/error.log", err.Error())
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			logs.LogToFile("logs/error.log", err.Error())
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(payment)
}

func (ps *paymentService) CalculateForeingAmount(partner_id int32, amount float64) float64 {
	url := fmt.Sprintf("http://localhost:3001/api/v1/partners/%d", partner_id)

	r, err := http.Get(url)
	if err != nil {
		return 0
	}
	defer r.Body.Close()

	partner := dto.PartnerData{}
	err = json.NewDecoder(r.Body).Decode(&partner)
	if err != nil {
		return 0
	}

	if partner.Currency == "USD" {
		return amount * consts.USD
	} else if partner.Currency == "EUR" {
		return amount * consts.EUR
	} else if partner.Currency == "GBP" {
		return amount * consts.GBP
	} else {
		return 0
	}
}

func (ps *paymentService) SavePayments(paymentData dto.PaymentData) (*dto.PaymentResponse, error) {

	consumer, _ := ps.paymentRepository.FindConsumerByNationalId(context.Background(), paymentData.Consumer.National_id)
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

	foreingAmount := ps.CalculateForeingAmount(paymentData.Partner_id, paymentData.Amount)
	dtoPaymentResponse := dto.PaymentResponse{
		ID:             newPayment.ID,
		Partner_id:     newPayment.Partner_id,
		Amount:         newPayment.Amount,
		Foreing_amount: foreingAmount,
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

func (ps *paymentService) HandlerGetPaymentId(w http.ResponseWriter, r *http.Request) {

	paymentID := mux.Vars(r)["id"]

	if paymentID != "" {
		entityPayment, err := ps.paymentRepository.FindPaymentById(context.Background(), paymentID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			logs.LogToFile("logs/error.log", err.Error())
			return
		}
		consumer, err := ps.paymentRepository.FindConsumerById(context.Background(), entityPayment.Consumer_id)
		if *consumer == (entity.Consumer{}) {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Consumer not found"))
			logs.LogToFile("logs/error.log", "Consumer not found")
			return
		}

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			logs.LogToFile("logs/error.log", err.Error())
			return
		}

		foreingAmount := ps.CalculateForeingAmount(entityPayment.Partner_id, entityPayment.Amount)
		w.WriteHeader(http.StatusOK)
		dtoPaymentResponse := dto.PaymentResponse{
			ID:             entityPayment.ID,
			Partner_id:     entityPayment.Partner_id,
			Amount:         entityPayment.Amount,
			Foreing_amount: foreingAmount,
			Consumer: dto.ConsumerData{
				ID:          consumer.ID,
				Name:        consumer.Name,
				National_id: consumer.National_id,
			},

			Created_at: entityPayment.Created_at.Format("2006-01-02 15:04:05"),
			Updated_at: entityPayment.Updated_at.Format("2006-01-02 15:04:05"),
		}

		json.NewEncoder(w).Encode(dtoPaymentResponse)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte("Bad Request"))
}

func (ps *paymentService) HandlerGetPayment(w http.ResponseWriter, r *http.Request) {
	offsetStr := r.URL.Query().Get("offset")
    limitStr := r.URL.Query().Get("limit")

    offset, err := strconv.Atoi(offsetStr)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid offset parameter"))
		logs.LogToFile("logs/error.log", err.Error())
        return
    }
    limit, err := strconv.Atoi(limitStr)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write([]byte("Invalid limit parameter"))
		logs.LogToFile("logs/error.log", err.Error())
        return
    }

	payments, err := ps.paymentRepository.GetPayments(context.Background(), offset, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		logs.LogToFile("logs/error.log", err.Error())
	}

	var paymentResponses []dto.PaymentResponse

	for _, payment := range payments {
		consumer, err := ps.paymentRepository.FindConsumerById(context.Background(), payment.Consumer_id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal Server Error"))
			logs.LogToFile("logs/error.log", err.Error())
			return
		}

		paymentResponse := dto.PaymentResponse{
			ID:         payment.ID,
			Partner_id: payment.Partner_id,
			Amount:     payment.Amount,
			Consumer: dto.ConsumerData{
				ID:          consumer.ID,
				Name:        consumer.Name,
				National_id: consumer.National_id,
			},
			Created_at: payment.Created_at.Format("2006-01-02 15:04:05"),
			Updated_at: payment.Updated_at.Format("2006-01-02 15:04:05"),
		}
		paymentResponses = append(paymentResponses, paymentResponse)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(paymentResponses)
}
