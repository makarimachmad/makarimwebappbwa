package user

type UserFormatter struct{
	ID int `json:"id"`
	Name string `json:"name"`
	Ocuption string `json:"ocuption"`
	Email string `json:"email"`
	Token string `json:"token"`
	Image_URL	string	`json:"image_url"`
}

func FormatUser(user User, token string) UserFormatter{
	formatter := UserFormatter{
		ID: user.ID,
		Name: user.Name,
		Ocuption: user.Ocuption,
		Email: user.Email,
		Token: token,
		Image_URL: user.Avatar,
	}
	return formatter
}
