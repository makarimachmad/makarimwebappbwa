package campaign

import (
	"github.com/makarimachmad/makarimwebappbwa/user"

	"time")


type Campaign struct {
	ID               int
	UserID           int
	Name             string
	ShortDescription string
	Description      string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Perks		 string
	Slug             string
	CreatedAt        time.Time
	UpdatedAt	 time.Time
	GambarCampaigns	 []GambarCampaign
	User		 user.User
}

type GambarCampaign struct{
	ID		int
	CampaignID	int
	FileName	string
	IsPrimary	int
	CreatedAt	time.Time
	UpdatedAt	time.Time	
}