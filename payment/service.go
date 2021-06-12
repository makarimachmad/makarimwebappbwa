package payment

import (
	"github.com/makarimachmad/makarimwebappbwa/user"

	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"github.com/veritrans/go-midtrans"
)


type Service interface {
	GetPaymentURL(transaction Transaction, user user.User) (string, error)
}

type service struct {
}

func NewService() *service{
	return &service{}
}

func (s *service) GetPaymentURL(transaction Transaction, user user.User) (string, error){
	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file. Make sure .env file is exists!")
	}
	
	midclient := midtrans.NewClient()
    midclient.ServerKey = os.Getenv("SERVER_KEY")
    midclient.ClientKey = os.Getenv("CLIENT_KEY")
    midclient.APIEnvType = midtrans.Sandbox

    snapGateway := midtrans.SnapGateway{
        Client: midclient,
    }

	snapReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			Email: user.Email,
			FName: user.Name,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID: strconv.Itoa(transaction.ID),
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapGateway.GetToken(snapReq)
	log.Println("GetToken:")
    snapTokenResp, err := snapGateway.GetToken(snapReq)
	if err != nil{
		return "", err
	}
	return snapTokenResp.RedirectURL, nil
}