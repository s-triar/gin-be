package controller

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	pvalidator "github.com/go-passwd/validator"
	"github.com/go-playground/validator/v10"

	"gin-be/internal/database"
	"gin-be/internal/model"
	"gin-be/internal/service"
	"gin-be/internal/tool"
)

var validate *validator.Validate
var emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

func init() {
	validate = validator.New()
	validate.RegisterValidation("email-validator", EmailValidator)
	validate.RegisterValidation("password-validator", PasswordValidator)
}

func PasswordValidator(fl validator.FieldLevel) bool {
	passwordValidator := pvalidator.New(pvalidator.MinLength(5, fmt.Errorf("password must be at least 5 characters")))
	err := passwordValidator.Validate(fl.Field().String())
	return err == nil
}

func EmailValidator(fl validator.FieldLevel) bool {
	return emailRegex.MatchString(fl.Field().String())
}

type UserRegister struct {
	Fullname        string `json:"Fullname" binding:"required"`
	Email           string `json:"Email" binding:"required" validate:"email-validator"`
	Phone           string `json:"Phone" binding:"required"`
	Password        string `json:"Password" binding:"required" validate:"password-validator"`
	ConfirmPassword string `json:"ConfirmPassword" binding:"required"`
}
type RegisterUserResponse struct {
	model.Response[string]
}
type UserLogin struct {
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
type LoginResponse struct {
	model.Response[string]
}
type UserDTO struct { // User logged In
	ID       string `json:"id"`
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// @BasePath /api
// AuthR godoc
// @Summary Register user with email
// @Schemes
// @Description Register user with email
// @Tags v1/auth
// @Accept json
// @Produce json
// @Param user body UserRegister true "User Register object"
// @Success 201 {object} RegisterUserResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /v1/auth/register [post]
func RegisterControllerMethod(c *gin.Context) {

	var input UserRegister
	resp := RegisterUserResponse{
		Response: model.Response[string]{
			Message: "",
			Data:    "",
			Error:   "",
		},
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if err := validate.Struct(input); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	if input.Password != input.ConfirmPassword {
		resp.Error = "Confirm Password does not match Password"
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	clientTx, errr := database.GetDB().GetDBClientEntTx(c)

	if errr != nil {
		resp.Error = errr.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	_, err := service.RegisterUserByEmail(
		c,
		clientTx,
		input.Fullname,
		input.Email,
		input.Phone,
		input.Password,
	)

	if err != nil {
		resp.Error = err.Error()
		err := database.GetDB().RollbackTransaction(clientTx, err)
		if err != nil {
			resp.Error = err.Error()
		}
		c.JSON(http.StatusBadRequest, resp)
		return
	} else {
		clientTx.Commit()
	}
	resp.Message = "User is created successfully"
	c.JSON(http.StatusCreated, resp)
}

// @BasePath /api
// AuthR godoc
// @Summary Login user with email
// @Schemes
// @Description Login user with email
// @Tags v1/auth
// @Accept json
// @Produce json
// @Param user body UserLogin true "User Login object"
// @Success 201 {object} LoginResponse
// @Failure 400 {object} model.ErrorResponse
// @Router /v1/auth/login [post]
func LoginControllerMethod(c *gin.Context) {
	var input UserLogin
	resp := LoginResponse{
		Response: model.Response[string]{
			Message: "",
			Data:    "",
			Error:   "",
		},
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	client := database.GetDB().GetDBClientEnt()
	user, err := service.LoginUserByEmail(c, client, input.Email, input.Password)
	if err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	token, err := tool.GenerateJWTToken(user)
	if err != nil {
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	resp.Data = token
	c.JSON(http.StatusOK, resp)
}

// @BasePath /api
// AuthR godoc
// @Summary Check existing email
// @Schemes
// @Description Check existing email
// @Tags v1/auth
// @Accept json
// @Produce json
// @Param email query string true "Email to be compared"
// @Success 200 {object} bool
// @Failure 400 {object} model.ErrorResponse
// @Router /v1/auth/check_email [get]
func CheckEmailControllerMethod(c *gin.Context) {
	var email = c.Query("email")
	resp := model.ErrorResponse{
		Response: model.Response[string]{
			Message: "",
			Data:    "",
			Error:   "",
		},
	}
	client := database.GetDB().GetDBClientEnt()
	result, err := service.CheckExistingEmail(c, client, email)
	if err != nil {
		if err.Error() == "ent: user not found" {
			c.JSON(http.StatusOK, false)
			return
		}
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @BasePath /api
// AuthR godoc
// @Summary Check existing phone
// @Schemes
// @Description Check existing phone
// @Tags v1/auth
// @Accept json
// @Produce json
// @Param phone query string true "Phone to be compared"
// @Success 200 {object} bool
// @Failure 400 {object} model.ErrorResponse
// @Router /v1/auth/check_phone [get]
func CheckPhoneControllerMethod(c *gin.Context) {
	var phone = c.Query("phone")
	resp := model.ErrorResponse{
		Response: model.Response[string]{
			Message: "",
			Data:    "",
			Error:   "",
		},
	}
	client := database.GetDB().GetDBClientEnt()
	result, err := service.CheckExistingPhone(c, client, phone)
	if err != nil {
		if err.Error() == "ent: user not found" {
			c.JSON(http.StatusOK, false)
			return
		}
		resp.Error = err.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	c.JSON(http.StatusOK, result)
}

// @BasePath /api
// AuthR godoc
// @Security BearerAuth
// @Summary Get user profile
// @Schemes
// @Description Get user profile
// @Tags v1/auth
// @Accept json
// @Produce json
// @Success 200 {object} UserDTO
// @Failure 400 {object} model.ErrorResponse
// @Router /v1/auth/user [get]
func UserProfileControllerMethod(c *gin.Context) {
	resp := model.ErrorResponse{
		Response: model.Response[string]{
			Message: "",
			Data:    "",
			Error:   "",
		},
	}

	uuidUser, errr := service.ExtractUserLoggedIn(c)

	if errr != nil {
		resp.Error = errr.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	client := database.GetDB().GetDBClientEnt()

	if errr != nil {
		resp.Error = errr.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	result, errr := service.GetUserById(c, client, *uuidUser)

	if errr != nil {
		resp.Error = errr.Error()
		c.JSON(http.StatusBadRequest, resp)
		return
	}

	data := UserDTO{
		ID:       result.ID.String(),
		Fullname: result.Fullname,
		Email:    result.Email,
		Phone:    result.Phone,
	}

	c.JSON(http.StatusOK, data)
}
