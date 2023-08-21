package partnerservices

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/wjuneo/bexs/dto/partnerdto"
	"github.com/wjuneo/bexs/entity"
	"github.com/wjuneo/bexs/errors"
	"github.com/wjuneo/bexs/logs"
	"github.com/wjuneo/bexs/repository/partnerrepository"
)

type PartnerService interface {
	HandlerRequest(w http.ResponseWriter, r *http.Request)
	SavePartners(partnerData partnerdto.PartnerData) (*entity.Partner, error)
	GetPartners(w http.ResponseWriter, r *http.Request)
}

type partnerService struct {
	partnerRepository partnerrepository.PartnerRepository
}

func NewPartnerService(partnerRepository partnerrepository.PartnerRepository) PartnerService {
	return &partnerService{
		partnerRepository: partnerRepository,
	}
}

func ValidatedCurrency(currency string) bool {
	allowedCurrencies := []string{"GBP", "EUR", "USD"}

	for _, allowedCurrency := range allowedCurrencies {
		if allowedCurrency == currency {
			return true
		}
	}
	return false
}

func (ps *partnerService) ValidatePartners(partnerData partnerdto.PartnerData) error {

	partner, _ := ps.partnerRepository.FindPartnerByDocument(context.Background(), partnerData.Document)
	if partner != nil {
		logs.LogToFile("logs/error.log", errors.ErrPartnerAlreadyExists)
		return fmt.Errorf(errors.ErrPartnerAlreadyExists)
	}

	if !ValidatedCurrency(partnerData.Currency) {
		logs.LogToFile("logs/error.log", errors.ErrCurrencyNotAllowed)
		return fmt.Errorf(errors.ErrCurrencyNotAllowed)
	}
	return nil
}

func (ps *partnerService) HandlerRequest(w http.ResponseWriter, r *http.Request) {
	var partnerData partnerdto.PartnerData
	if err := json.NewDecoder(r.Body).Decode(&partnerData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if partnerData.TradingName == "" || partnerData.Document == "" || partnerData.Currency == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Invalid data"))
		return
	}

	entity, err := ps.SavePartners(partnerData)
	if err != nil {
		if err.Error() == errors.ErrPartnerAlreadyExists {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(errors.ErrPartnerAlreadyExists))
			return
		} else if err.Error() == errors.ErrCurrencyNotAllowed {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(errors.ErrCurrencyNotAllowed))
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errors.ErrInternalServerError))
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(entity)
}

func (ps *partnerService) SavePartners(partnerData partnerdto.PartnerData) (*entity.Partner, error) {

	err := ps.ValidatePartners(partnerData)
	if err != nil {
		return nil, err
	}

	entity := entity.Partner{
		Trading_name: partnerData.TradingName,
		Document:     partnerData.Document,
		Currency:     partnerData.Currency,
	}

	newEntity, err := ps.partnerRepository.SavePartners(context.Background(), entity)
	if err != nil {
		logs.LogToFile("logs/error.log", errors.ErrInternalServerError)
		return nil, fmt.Errorf(errors.ErrInternalServerError)
	}

	return newEntity, nil
}

func (ps *partnerService) GetPartners(w http.ResponseWriter, r *http.Request) {
	partnerID := mux.Vars(r)["id"]
	if partnerID != "" {
		partner, err := ps.partnerRepository.FindPartnerByID(context.Background(), partnerID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(errors.ErrInternalServerError))
			return
		}

		if partner == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(errors.ErrPartnerNotFound))
			return
		}

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(partner)
		return
	}

	partners, err := ps.partnerRepository.GetPartners(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errors.ErrInternalServerError))
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(partners)
	return
}
