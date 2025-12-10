package compaign

import (
	"emailn/internal/contract"
	"emailn/internal/internalerros"
)

type CompaignService struct { // CompaignService = serviço de compaigns
	repository CompaignRepository // repository = repository de compaigns
}

func NewCompaignService(repository CompaignRepository) *CompaignService { // NewCompaignService = cria um novo serviço de compaigns
	return &CompaignService{repository: repository}
}

func (s *CompaignService) Create(newCampaignDto contract.NewCampaignDto) (string, error) { // Save = salva um compaign

	compaign, err := NewCompaign(newCampaignDto.Name, newCampaignDto.Content, newCampaignDto.Contacts) // NewCompaign = cria um novo compaign
	if err != nil {
		return "", err
	}
	err = s.repository.Save(compaign) // Save = salva um compaign
	if err != nil {
		return "", internalerros.ErrInternalServerError
	}
	return compaign.ID, nil // ID = ID = retorna o ID do compaign
}
