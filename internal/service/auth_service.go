package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"gin-be/internal/ent"
	"gin-be/internal/ent/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"gin-be/internal/model"
)

func ExtractUserLoggedIn(ctx *gin.Context) (*uuid.UUID, error) {
	user_id, err := ctx.Get("user_id")

	if !err {
		return nil, fmt.Errorf("User is not exist in context key")
	}

	uuidUser, errr := uuid.Parse(user_id.(string))

	if errr != nil {
		return nil, errr
	}

	return &uuidUser, nil
}

func RegisterUserByEmail(
	ctx context.Context,
	clientTx *ent.Tx,
	fullname string,
	email string,
	phone string,
	password string,
) (*ent.User, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	hashedtoken, err := bcrypt.GenerateFromPassword([]byte(email+password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	entity, err := clientTx.User.
		Create().
		SetFullname(strings.TrimSpace(fullname)).
		SetEmail(strings.TrimSpace(email)).
		SetPhone(strings.TrimSpace(phone)).
		SetPassword(string(hashedPassword)).
		SetTokenAuth(string(hashedtoken)).
		SetIsEmailConfirmed(false).
		SetIsPhoneConfirmed(false).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	return entity, nil

}

func LoginUserByEmail(
	ctx context.Context,
	client *ent.Client,
	email string,
	password string,

) (*model.User, error) {

	entity, err := client.User.Query().
		Where(user.Email(email)).
		First(ctx)

	if err != nil {
		log.Printf("service_auth.go|LoginUserByEmail|%s", err.Error())
		return nil, fmt.Errorf("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(entity.Password), []byte(password))
	if err != nil {
		return nil, fmt.Errorf("password does not match")
	}

	user := model.User{
		ID:       entity.ID,
		Fullname: entity.Fullname,
		Email:    entity.Email,
		Phone:    entity.Phone,
	}
	return &user, nil
}

func GetUserById(
	ctx context.Context,
	client *ent.Client,
	id uuid.UUID,
) (*model.User, error) {

	entity, err := client.User.Query().
		Where(user.ID(id)).
		First(ctx)

	if err != nil {
		log.Printf("service_auth.go|GetUserById|%s", err.Error())
		return nil, fmt.Errorf("User not found")
	}

	user := model.User{
		ID:       entity.ID,
		Fullname: entity.Fullname,
		Email:    entity.Email,
		Phone:    entity.Phone,
	}
	return &user, nil

}

func CheckExistingPhone(
	ctx context.Context,
	client *ent.Client,
	phone string,
) (bool, error) {

	entity, err := client.User.Query().
		Where(user.Phone(phone)).
		First(ctx)

	if err != nil {
		return false, err
	}

	return entity != nil, nil

}

func CheckExistingEmail(
	ctx context.Context,
	client *ent.Client,
	email string,
) (bool, error) {

	entity, err := client.User.Query().
		Where(user.EmailEqualFold(strings.ToLower(email))).
		First(ctx)

	if err != nil {
		return false, err
	}

	return entity != nil, nil

}
