package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_NewCampaign_CreateCampaign(t *testing.T) {
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
	println(campaign.ID)	
	assert.Equal(campaign.ID, "1")
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

	assert.Equal(campaign2.ID, "1")
	assert.Equal(campaign2.Name, nome)
	assert.Equal(campaign2.Content, conteudo)
	assert.Equal(len(campaign2.Contacts), len(contato))

}

func Test_NewCampaign_IDisNotNill(t *testing.T) {
	assert := assert.New(t)
	name := "Campaing x"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	assert.NotNil(campaign.ID)
}

func Test_NewCampaign_CreatedOnIsNotNill(t *testing.T) {
	assert := assert.New(t)
	name := "Campaing x"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}
	now := time.Now().Add(-time.Minute)

	campaign := NewCampaign(name, content, contacts)

	assert.Greater(campaign.CreatedOn, now)
}