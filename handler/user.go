package handler

import (
	"github.com/makarimachmad/makarimwebappbwa/auth"
	"github.com/makarimachmad/makarimwebappbwa/helper"
	"github.com/makarimachmad/makarimwebappbwa/user"

	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)


type userHandler struct {
	userService user.Service
	authservice auth.Service
}

func NewUserHandler(userService user.Service, authservice auth.Service) *userHandler{
	return &userHandler{userService, authservice}
}

func (h *userHandler) RegisterUser(c *gin.Context){
	//tangkap input dari user
	//memetakan dari user ke struct registerinput
	//struct di atas disimpan menjadi service

	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil{

		errors := helper.FormatValidationError(err) 
		errorsMessage := gin.H{"errors": errors}
		
		response := helper.APIResponse("Gagal membuat akun", http.StatusUnprocessableEntity, "Gagal", errorsMessage)
		c.JSON(http.StatusBadRequest, response)
		return //supaya gk eksekusi ke bawah
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil{
		response := helper.APIResponse("Gagal membuat akun", http.StatusBadRequest, "Gagal", err)
		c.JSON(http.StatusBadRequest, response)
		return //supaya gk eksekusi ke bawah
	}

	token, err := h.authservice.GenarateToken(newUser.ID)
	if err != nil{
		response := helper.APIResponse("Gagal membuat akun", http.StatusBadRequest, "Gagal", err)
		c.JSON(http.StatusBadRequest, response)
		return //supaya gk eksekusi ke bawah
	}	

	formatter := user.FormatUser(newUser, token)

	response := helper.APIResponse("Akun berhasil terdaftar", http.StatusOK, "Sukses", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) Login(c *gin.Context){
	//user masukkan email dan sandi
	//input ditangkap handler
	//pemetaan input user ke struct user
	//input struct passing service
	//service mencari dg bantuan repo dengan email
	//mencocokan password

	var input user.LoginInput

	// cek  validasi format
	err := c.ShouldBindJSON(&input)
	if err != nil{
		errors := helper.FormatValidationError(err) 
		errorsMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Gagal masuk", http.StatusUnprocessableEntity, "Gagal", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return //kalau gagal tidak akan diteruskan
	}

	loggedinUser, err := h.userService.Login(input)
	if err != nil{
		errorsMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Gagal masuk", http.StatusUnprocessableEntity, "Gagal", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	token, err := h.authservice.GenarateToken(loggedinUser.ID)
	if err != nil{
		response := helper.APIResponse("Gagal login", http.StatusBadRequest, "Gagal", err)
		c.JSON(http.StatusBadRequest, response)
		return //supaya gk eksekusi ke bawah
	}	

	formatter := user.FormatUser(loggedinUser, token)
	response := helper.APIResponse("Berhasil masuk", http.StatusOK, "Sukses", formatter)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) CheckEmail(c *gin.Context){
	//cek email apakah sudah pernah terdaftar atau belum
	//ada input email dari user
	//input email dimapping ke struct input
	//struct input dipassing ke service
	//service akan memanggil repository apakah ada
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil{
		errors := helper.FormatValidationError(err) 
		errorsMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email sudah ada", http.StatusUnprocessableEntity, "Gagal", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil{
		errorsMessage := gin.H{"errors": "server eror"}
		response := helper.APIResponse("Email sudah ada", http.StatusUnprocessableEntity, "Gagal", errorsMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available" : isEmailAvailable,
	}
	metaMessage :="emil terdaftar"
	if isEmailAvailable{
		metaMessage="emil tersedia"
	}
	response := helper.APIResponse(metaMessage, http.StatusOK, "Berhasil", data)
	c.JSON(http.StatusOK, response)
}

func (h *userHandler) UploadAvatar(c *gin.Context){
	//input dari pengguna
	//simpan gambar di folder images/
	//di service panggil repository untuk menentukan user siapa yg bisa akses
	//JWT hardcode seakan2 user login dengan ID=1
	//di db itu menyimpan path gambarnya

	//c.SaveUploadedFile(file, )
	file, err := c.FormFile("avatar")
	if err != nil{
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("gagal upload gambar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil{
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("gagal upload gambar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil{
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("gagal upload gambar", http.StatusBadRequest, "error", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("berhasil upload gambar", http.StatusOK, "sukses", data)
	c.JSON(http.StatusOK, response)	
}

func (h *userHandler) FetchUser(c *gin.Context) {

	currentUser := c.MustGet("currentUser").(user.User)

	formatter := user.FormatUser(currentUser, "")

	response := helper.APIResponse("Successfuly fetch user data", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}