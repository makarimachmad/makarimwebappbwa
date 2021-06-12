package user

import "time"


type User struct {
	ID           int
	Name         string
	Role         string
	Ocuption     string
	Email        string
	Passwordhash string
	Avatar       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}