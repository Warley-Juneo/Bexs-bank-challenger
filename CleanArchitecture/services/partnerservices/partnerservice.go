package partnerservices

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/wjuneo/bexs/dto/partnerdto"
	"github.com/wjuneo/bexs/entity"
	"github.com/wjuneo/bexs/repository/partnerrepository"
)

type PartnerService interface {
	SavePartners(w http.ResponseWriter, r *http.Request)
}

type partnerService struct {
	partnerRepository partnerrepository.PartnerRepository
}

func NewPartnerService(partnerRepository partnerrepository.PartnerRepository) PartnerService {
	return &partnerService{
		partnerRepository: partnerRepository,
	}
}

func (ps partnerService) SavePartners(w http.ResponseWriter, r *http.Request) {

	var partnerdto partnerdto.PartnerData
	if err := json.NewDecoder(r.Body).Decode(&partnerdto); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	entity := entity.Partner{
		Trading_name: partnerdto.TradingName,
		Document:     partnerdto.Document,
		Currency:     partnerdto.Currency,
	}

	ctx := context.Background()
	_, err := ps.partnerRepository.SavePartners(ctx, entity)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Partner created successfully"))
}
