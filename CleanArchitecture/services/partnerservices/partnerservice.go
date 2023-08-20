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

func (ps *partnerService) SavePartners(partnerData partnerdto.PartnerData) (*entity.Partner, error) {
    entity := entity.Partner{
        Trading_name: partnerData.TradingName,
        Document:     partnerData.Document,
        Currency:     partnerData.Currency,
    }

    ctx := context.Background()
    newEntity, err := ps.partnerRepository.SavePartners(ctx, entity)
    if err != nil {
        return	nil, err
    }

    return newEntity, nil
}


func (ps *partnerService) HandlerRequest(w http.ResponseWriter, r *http.Request) {
    var partnerData partnerdto.PartnerData
    if err := json.NewDecoder(r.Body).Decode(&partnerData); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    _, err := ps.SavePartners(partnerData)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("Partner created successfully"))
}
