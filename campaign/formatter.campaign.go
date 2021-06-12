package campaign

import "strings"

type CampaignFormatter struct {
	ID               int    `json:"id"`
	UserID           int    `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
	BackerCount	 int	`json:"backer_count"`
}

type CampaignDetailFormatter struct {
	ID               int      						`json:"id"`
	Name             string   						`json:"name"`
	ShortDescription string   						`json:"short_description"`
	Description      string   						`json:"description"`
	ImageUrl         string   						`json:"image_url"`
	GoalAmount       int      						`json:"goal_amount"`
	UserID           int      						`json:"user_id"`
	CurrentAmount    int      						`json:"current_ammount"`
	BackerCount		 int							`json:"backer_count"`
	Slug             string   						`json:"slug"`
	Perks            []string 						`json:"perks"`
	User			 UserCampaignDetailFormatter	`json:"user"`
	Gambars			 []GambarCampaignDetailFormatter`json:"gambar"`
}

type UserCampaignDetailFormatter struct{
	Name		string 	`json:"name"`
	ImageUrl	string	`json:"image_url"`
}

type GambarCampaignDetailFormatter struct{
	ImageUrl	string	`json:"image_url"`
	IsPrimary	bool	`json:"is_primary"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{}
	campaignFormatter.ID = campaign.ID
	campaignFormatter.UserID = campaign.UserID
	campaignFormatter.Name = campaign.Name
	campaignFormatter.ShortDescription = campaign.ShortDescription
	campaignFormatter.GoalAmount = campaign.GoalAmount
	campaignFormatter.CurrentAmount = campaign.CurrentAmount
	campaignFormatter.BackerCount = campaign.BackerCount
	campaignFormatter.ImageUrl = ""
	campaignFormatter.Slug = campaign.Slug

	if len(campaign.GambarCampaigns) > 0 {
		campaignFormatter.ImageUrl = campaign.GambarCampaigns[0].FileName
	}
	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {

	campaignsFormatter := []CampaignFormatter{} //nilai default

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}
	return campaignsFormatter
}

func FormatCampaignDetail(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{}
	campaignDetailFormatter.ID = campaign.ID
	campaignDetailFormatter.Name = campaign.Name
	campaignDetailFormatter.ShortDescription = campaign.ShortDescription
	campaignDetailFormatter.Description = campaign.Description
	campaignDetailFormatter.GoalAmount = campaign.GoalAmount
	campaignDetailFormatter.CurrentAmount = campaign.CurrentAmount
	campaignDetailFormatter.UserID = campaign.UserID
	campaignDetailFormatter.ImageUrl = ""
	campaignDetailFormatter.Slug = campaign.Slug

	if len(campaign.GambarCampaigns) > 0 {
		campaignDetailFormatter.ImageUrl = campaign.GambarCampaigns[0].FileName
	}

	var perks []string
	for _, perk := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(perk))
	}

	campaignDetailFormatter.Perks = perks
	
	user := campaign.User
	userCampaignDetailFormatter := UserCampaignDetailFormatter{}
	userCampaignDetailFormatter.Name = user.Name
	userCampaignDetailFormatter.ImageUrl = user.Avatar

	campaignDetailFormatter.User = userCampaignDetailFormatter

	gambars := []GambarCampaignDetailFormatter{}
	
	for _, gambar := range campaign.GambarCampaigns{
		gambarCampaignDetailFormatter := GambarCampaignDetailFormatter{}
		gambarCampaignDetailFormatter.ImageUrl = gambar.FileName

		isPrimary := false
		if gambar.IsPrimary == 1{
			isPrimary = true
		}
		gambarCampaignDetailFormatter.IsPrimary = isPrimary
		gambars = append(gambars, gambarCampaignDetailFormatter)
	}
	campaignDetailFormatter.Gambars = gambars	

	return campaignDetailFormatter
}