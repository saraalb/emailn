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

	//Action
	campaign := NewCampaign(name, content, contacts)
	
	//Assert
	println(campaign.ID)	
	assert.Equal(campaign.Name, name)
	assert.Equal(campaign.Content, content)
	assert.Equal(len(campaign.Contacts), len(contacts))

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