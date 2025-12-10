package contract

type NewCampaignDto struct { // NewCampaignDto = DTO para criar um novo compaign
	Name     string
	Content  string
	Contacts []string
}
