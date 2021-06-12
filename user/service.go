package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// mapping struct input ke struct user
// simpan struct user melalui repository

type Service interface{
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, filelocation string) (User, error)
	GetUserByID(ID int) (User, error)
}

type service struct{
	repository Repository
}

func NewService(repository Repository) *service{
	return &service{repository}
}

func(s *service) RegisterUser(input RegisterUserInput) (User, error){
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Ocuption = input.Ocuption
	Passwordhash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)

	if err != nil{
		return user, err
	}
	user.Passwordhash = string(Passwordhash)
	user.Role = "boss"
	
	newUser, err := s.repository.Save(user)
	if err != nil{
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginInput) (User, error){
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil{
		return user, err
	}
	if user.ID == 0{
		return user, errors.New("pengguna tidak ditemukan")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Passwordhash), []byte(password))
	if err != nil{
		return user, err
	}
	return user, nil
}

func (s *service) IsEmailAvailable(input CheckEmailInput) (bool, error){
	email := input.Email
	
	user, err := s.repository.FindByEmail(email)
	if err != nil{
		return false, err
	}
	if user.ID == 0{
		return true, nil
	}
	return false, nil
}

func (s* service) SaveAvatar(ID int, filelocation string) (User, error){
	user, err := s.repository.CariById(ID)
	if err != nil{
		return user, err
	}

	user.Avatar = filelocation
	updatedUser, err := s.repository.Update(user)
	if err != nil{
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) GetUserByID(ID int) (User, error){
	user, err := s.repository.CariById(ID)
	if err != nil{
		return user, err
	}

	if user.ID == 0{
		return user, errors.New("berdasarkan email tersebut pengguna tidak ditemukan")
	}
	return user, nil
}