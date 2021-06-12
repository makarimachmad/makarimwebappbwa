package transaction

import "time"

type CampaignTransactionsFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type UserTransactionsFormatter struct{
	ID				  int				`json:"id"`
	Amount			  int				`json:"mount"`
	Status			  string			`json:"status"`
	CreatedAt		  time.Time			`json:"created_at"`
	Campaign		  CampaignFormatter	`json:"campaign"`
}

type CampaignFormatter struct{
	Name		string	`json:"name"`
	ImageUrl	string	`json:"image_url"`
}

type PaymentFormatter struct{
	ID			int		  `json:"id"`
	CampaignID	int	      `json:"campaign_id"`
	UserID		int	   	  `json:"user_id"`
	Amount		int		  `json:"amount"`
	Status		string	  `json:"status"`
	Code		string	  `json:"code"`
	PaymentURL  string	  `json:"payment_url"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionsFormatter {
	formatter := CampaignTransactionsFormatter{}
	formatter.ID = transaction.ID
	formatter.Name = transaction.User.Name
	formatter.Amount = transaction.Amount
	formatter.CreatedAt = transaction.CreatedAt
	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionsFormatter {
	if len(transactions) == 0 {
		return []CampaignTransactionsFormatter{}
	}

	var transactionsformatter []CampaignTransactionsFormatter

	for _, transaction := range transactions {
		formatter := FormatCampaignTransaction(transaction)
		transactionsformatter = append(transactionsformatter, formatter)
	}

	return transactionsformatter
}

//user transaction
func FormatUserTransaction(transaction Transaction) UserTransactionsFormatter{
	formatter := UserTransactionsFormatter{}
	formatter.ID = transaction.ID
	formatter.Amount = transaction.Amount
	formatter.Status = transaction.Status
	formatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Name = transaction.Campaign.Name
	
	campaignFormatter.ImageUrl = ""
	if len(transaction.Campaign.GambarCampaigns) > 0 {
		campaignFormatter.ImageUrl = transaction.Campaign.GambarCampaigns[0].FileName
	}
	formatter.Campaign = campaignFormatter
	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionsFormatter {
	if len(transactions) == 0 {
		return []UserTransactionsFormatter{}
	}

	var transactionsformatter []UserTransactionsFormatter

	for _, transaction := range transactions {
		formatter := FormatUserTransaction(transaction)
		transactionsformatter = append(transactionsformatter, formatter)
	}

	return transactionsformatter
}

func FormatPaymentTransaction(transaction Transaction) PaymentFormatter {
	formatter := PaymentFormatter{}
	formatter.ID = transaction.ID
	formatter.CampaignID = transaction.CampaignID
	formatter.Amount = transaction.Amount
	formatter.UserID = transaction.UserID
	formatter.Code = transaction.Code
	formatter.Status = transaction.Status
	formatter.PaymentURL = transaction.PaymentUrl
	return formatter
}