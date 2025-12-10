package compaign

import (
	"testing"

	"emailn/internal/contract"

	"emailn/internal/internalerros"
	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct { // repositoryMock = repositoryMock = cria um novo repositório mock(teste de unidade)
	mock.Mock // Mock = Mock = cria um novo mock
}

var (
	newCampaign = contract.NewCampaignDto{ // newCampaign = newCampaign = cria um novo compaign
		Name:    "Test Compaign",
		Content: "This is a test compaign",
		Contacts: []string{ // Contacts = Contacts = cria um novo slice de contatos
			"test@example.com",
			"test2@example.com",
			"test3@example.com", // test3@example.com = test3@example.com = cria um novo contato
		},
	}
	service = CompaignService{}
)

//MOCK DO REPOSITÓRIO

func (m *repositoryMock) Save(compaign *Compaign) error { // Save = salva um compaign

	args := m.Called(compaign) // args = args.Called(compaign) = chama o método Save do repositório
	return args.Error(0)       // args.Error(0) = retorna o erro do método Save do repositório
}

func Test_Create_NewCompaignService(t *testing.T) {
	assert := assert.New(t)

	id, err := service.Create(newCampaign)
	assert.NotNil(id)

	assert.Nil(err)
}

func Test_Create_ValidateDomainErrors(t *testing.T) { // Teste para validar erros de domínio
	assert := assert.New(t)
	newCampaign.Name = ""

	_, err := service.Create(newCampaign)
	assert.NotNil(err)

	assert.Equal("name is required", err.Error())
}

func Test_Create_SaveCompaignService(t *testing.T) { // Teste para salvar um compaign
	// assert := assert.New(t)
	repositoryMock := new(repositoryMock) // repositoryMock = repositoryMock = cria um novo repositório mock

	repositoryMock.On("Save", mock.MatchedBy(func(compaign *Compaign) bool { // On = On = cria um novo evento de save
		if compaign.Name != newCampaign.Name { // compaign.Name != newCampaign.Name = compaign.Name != newCampaign.Name = verifica se o nome do compaign é igual ao nome do compaign
			return false
		}
		if compaign.Content != newCampaign.Content { // compaign.Content != newCampaign.Content = compaign.Content != newCampaign.Content = verifica se o conteúdo do compaign é igual ao conteúdo do compaign
			return false
		}
		if len(compaign.Contacts) != len(newCampaign.Contacts) { // len(compaign.Contacts) != len(newCampaign.Contacts) = len(compaign.Contacts) != len(newCampaign.Contacts) = verifica se o número de contatos do compaign é igual ao número de contatos do compaign
			return false
		}
		return true
	})).Return(nil) // Return = Return = retorna um nil

	service.repository = repositoryMock

	service.Create(newCampaign) // Create = Create = cria um novo compaign

	repositoryMock.AssertExpectations(t)
}

func Test_Create_ValidateDataBaseErrors(t *testing.T) { // Teste para validar erros de dados
	assert := assert.New(t)
	repositoryMock := new(repositoryMock) // repositoryMock = repositoryMock = cria um novo repositório mock

	repositoryMock.On("Save", mock.Anything).Return(errors.New("error saving campaign"))

	service.repository = repositoryMock // repository = repository = cria um novo repositório

	_, err := service.Create(newCampaign) // Create = Create = cria um novo compaign

	assert.True(errors.Is(err, internalerros.ErrInternalServerError))
	// repositoryMock.AssertExpectations(t)
}
