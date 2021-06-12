package transaction

import (
	"github.com/makarimachmad/makarimwebappbwa/campaign"
	"github.com/makarimachmad/makarimwebappbwa/payment"

	"errors"
	"fmt"
	"strconv"
	"time"
)


type service struct {
	repository 		   Repository
	campaignRepository campaign.Repository
	paymentService	   payment.Service
}

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error)
	GetTransactionByUserID(userID int) ([]Transaction, error)
	CreateTransactionInput(input CreateTransactionInput) (Transaction, error)
	ProcessPayment(input TransactionNotificationInput) error
}

func NewService(repository Repository, campaignRepository campaign.Repository, paymentService payment.Service) *service {
	return &service{repository, campaignRepository, paymentService}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsInput) ([]Transaction, error) {

	campaign, err := s.campaignRepository.FindByID(input.ID)
	if err != nil{
		return []Transaction{}, errors.New("bukan pemilik data")
	}
	if campaign.UserID != input.User.ID{
		return []Transaction{}, errors.New("bukan pemilik data")
	}

	transactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}

// user transaction
func (s *service) GetTransactionByUserID(userID int) ([]Transaction, error){
	transactions, err := s.repository.GetByUserID(userID)
	if err != nil{
		return transactions, err
	}
	return transactions, nil
}

func (s *service) CreateTransactionInput(input CreateTransactionInput) (Transaction, error){
	transaction := Transaction{}
	transaction.CampaignID = input.CampaignID
	transaction.Amount = input.Amount
	transaction.UserID = input.User.ID
	transaction.Status = "pending"

	stringcode := fmt.Sprintf("%d", input.User.ID)
	stringcode = stringcode + time.Now().Format("09-07-2017")
	transaction.Code = stringcode

	newTransaction, err := s.repository.Save(transaction)
	if err != nil{
		return newTransaction, err
	}

	paymentTransaction := payment.Transaction{
		ID: newTransaction.ID,
		Amount: newTransaction.Amount,
	}

	paymentURL, err := s.paymentService.GetPaymentURL(paymentTransaction, input.User)
	if err != nil{
		return newTransaction, err
	}

	newTransaction.PaymentUrl = paymentURL
	newTransaction, err = s.repository.Update(newTransaction)
	if err != nil{
		return newTransaction, err
	}

	return newTransaction, nil
}

//notifikasi midtrans

func (s *service) ProcessPayment(input TransactionNotificationInput) error{
	transactionID,_ := strconv.Atoi(input.OrderID)

	transaction, err := s.repository.GetByTransactionUserID(transactionID)
	if err != nil{
		return err
	}
	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept"{
		transaction.Status = "paid"
	}else if input.PaymentType == "settlement"{
		transaction.Status = "paid"
	}else if input.PaymentType == "deny" || input.TransactionStatus == "expire" || input.TransactionStatus == " cancel"{
		transaction.Status = "cancled"
	}

	updatedTransaction, err := s.repository.Update(transaction)
	if err != nil{
		return err
	}
	campaign, err := s.campaignRepository.FindByID(updatedTransaction.CampaignID)
	if err != nil{
		return err
	}
	if updatedTransaction.Status == "status"{
		campaign.BackerCount = campaign.BackerCount + 1
		campaign.CurrentAmount = campaign.CurrentAmount + updatedTransaction.Amount

		_, err := s.campaignRepository.Update(campaign)
		if err != nil{
			return err
		}
	}
	return nil
}