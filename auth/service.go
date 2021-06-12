package auth

import (
	"errors"

	"github.com/dgrijalva/jwt-go")


type Service interface {
	GenarateToken(userID int) (string, error)
	ValidateToken(token string) (*jwt.Token, error)
}

type jwtService struct {
}

var katakunci = []byte("m4k4r1mp4nc3n0y3")
//katakunci ditaro di env

func NewService() *jwtService{
	return &jwtService{}
}

func (s *jwtService) GenarateToken(userID int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userID
	//tambahin expired date, nama, 
	//midlleware(tambahin kalo sekarang > expired date maka toke expired)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	signedToken, err := token.SignedString(katakunci)
	if err != nil{
		return signedToken, err
	}
	return signedToken, nil
}

func (s *jwtService) ValidateToken(encodedtoken string) (*jwt.Token, error){
	token, err := jwt.Parse(encodedtoken, func(token *jwt.Token)(interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Token salah")
		}
		return []byte(katakunci), nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}