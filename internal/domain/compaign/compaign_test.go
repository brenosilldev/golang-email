package compaign_test

import (
	"emailn/internal/domain/compaign"

	"testing"

	"github.com/stretchr/testify/assert"
)

var ( // COLOCAR VARIAVEIS GLOBAIS PARA TESTES -- Constantes
	name     = "Test Compaign"
	content  = "This is a test compaign"
	contacts = []string{
		"test@example.com",
		"test2@example.com",
		"test3@example.com",
	}
)

func TestNewCompaign(t *testing.T) { // t *ttsting.T = Teste para função NewCompaign
	assert := assert.New(t) // assert = assert.New(t) = cria um novo assert para o teste

	compaign, err := compaign.NewCompaign(name, content, contacts)

	assert.Equal(compaign.Name, name)
	assert.Equal(compaign.Content, content)
	assert.Equal(len(compaign.Contacts), len(contacts))
	assert.NoError(err)

}

func Test_NewCompaignIDIsNotNil(t *testing.T) { // t *ttsting.T = Teste para função NewCompaignIDIsNotNil
	assert := assert.New(t)
	compaign, _ := compaign.NewCompaign("Test Compaign", "This is a test compaign", []string{
		"test@example.com",
		"test2@example.com",
		"test3@example.com",
	})

	assert.NotNil(compaign.CreatedAt)
}
