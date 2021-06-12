package transaction

import (
	"github.com/makarimachmad/makarimwebappbwa/campaign"
	"github.com/makarimachmad/makarimwebappbwa/user"
	
	"time"
)


type Transaction struct {
	ID         int
	CampaignID int
	UserID     int
	Amount     int
	Status     string
	Code       string
	Campaign   campaign.Campaign
	User	   user.User
	PaymentUrl string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type TransactionNotificationInput struct{
	TransactionStatus	string	`json:"transaction_status"`
	OrderID				string	`json:"order_id"`
	PaymentType			string	`json:"payment_type"`
	FraudStatus			string	`json:"fraud_status"`
}