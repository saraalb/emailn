package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewCampaign(t *testing.T) {
	//Arrange
	assert := assert.New(t)
	name := "Campaing x"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}
	nome := "nome"
	conteudo := "conteudo"
	contato := []string{"contato1@e.com", "contato2@e.com"}

	//Action
	campaign := NewCampaign(name, content, contacts)
	campaign2 :=  NewCampaign(nome, conteudo, contato)
	
	//Assert
	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))
	assert.Equal(campaign2.ID, "1")
	assert.Equal(campaign2.Name, nome)
	assert.Equal(campaign2.Content, conteudo)
	assert.Equal(len(campaign2.Contacts), len(contato))

}
