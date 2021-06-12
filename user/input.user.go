package user

type RegisterUserInput struct{
	Name string	`json:"name" binding:"required"`
	Ocuption string	`json:"ocuption" binding:"required"`
	Email string	`json:"email" binding:"required,email"`
	Password string	`json:"password" binding:"required"`
}

type LoginInput struct{
	Email string 	`json:"email" binding:"required,email"`
	Password string	`json:"password" binding:"required"`
}

type CheckEmailInput struct{
	Email string `json:"email" binding:"required,email"`
}