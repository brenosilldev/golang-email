package compaign

type CompaignRepository interface {
	Save(compaign *Compaign) error
}
