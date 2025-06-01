package campaign

import "testing"

func TestNewCampaign(t *testing.T) {
	name := "Campaing x"
	content := "Body"
	contacts := []string{"email1@e.com", "email2@e.com"}

	campaign := NewCampaign(name, content, contacts)

	if campaign.ID != "1" {
		t.Errorf("Expected campaign id to be 1, but got %s", campaign.ID)
	} else if campaign.Name != name {
		t.Errorf("Expected campaign name to be %s, but got %s", name, campaign.Name)
	} else if campaign.Content != content {
		t.Errorf("Expected campaign content to be %s, but got %s", content, campaign.Content)
	} else if len(campaign.Contacts) != len(contacts) {
		t.Errorf("Expected campaign contacts to be %d, but got %d", len(contacts), len(campaign.Contacts))
	}

}
