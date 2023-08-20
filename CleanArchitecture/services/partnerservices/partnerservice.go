package partnerservices

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/wjuneo/bexs/dto/partnerdto"
	"github.com/wjuneo/bexs/entity"
	"github.com/wjuneo/bexs/logs"
	"github.com/wjuneo/bexs/repository/partnerrepository"
)

type PartnerService interface {
	HandlerRequest(w http.ResponseWriter, r *http.Request)
	SavePartners(partnerData partnerdto.PartnerData) (*entity.Partner, error)
}

type partnerService struct {
	partnerRepository partnerrepository.PartnerRepository
}

func NewPartnerService(partnerRepository partnerrepository.PartnerRepository) PartnerService {
	return &partnerService{
		partnerRepository: partnerRepository,
	}
}

func (ps *partnerService) ValidatePartners(partnerData partnerdto.PartnerData) error {
	partner, _ := ps.partnerRepository.FindPartnerByDocument(context.Background(), partnerData.Document)
	if partner != nil {
		logs.LogToFile("logs/error.log", "a partner with that document already exists")
		return fmt.Errorf("a partner with that document already exists")
	}
	return nil
}

func (ps *partnerService) HandlerRequest(w http.ResponseWriter, r *http.Request) {
	var partnerData partnerdto.PartnerData
	if err := json.NewDecoder(r.Body).Decode(&partnerData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entity, err := ps.SavePartners(partnerData)
	if err != nil {
		if err.Error() == "a partner with that document already exists" {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("a partner with that document already exists"))
			return
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Internal server error"))
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

	ctx := context.Background()
	newEntity, err := ps.partnerRepository.SavePartners(ctx, entity)
	if err != nil {
		logs.LogToFile("logs/error.log", "failed to save partner in database")
		return nil, fmt.Errorf("failed to save partner")
	}

	return newEntity, nil
}
