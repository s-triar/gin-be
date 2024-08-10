package tool

import (
	"fmt"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"gin-be/internal/model"
)

func GenerateJWTToken(user *model.User) (string, error) {
	envApp := NewEnv(nil)
	claims := jwt.MapClaims{}
	claims["sub"] = user.ID

	claims["exp"] = time.Now().Add(time.Hour * time.Duration(envApp.JWT_LIFETIME)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(envApp.JWT_SECRET))

}

func ExtractToken(c *gin.Context) (*model.User, error) {
	envApp := NewEnv(nil)
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
		return nil, fmt.Errorf("authorization header required")
	}
	tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(envApp.JWT_SECRET), nil
	})
	if err != nil {
		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return nil, fmt.Errorf(err.Error())
	}
	if !token.Valid {
		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return nil, fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return nil, fmt.Errorf("invalid token claims")
	}

	uuidUser, err := uuid.Parse(claims["sub"].(string))
	if err != nil {
		// c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid ID user claims"})
		return nil, fmt.Errorf("invalid ID user claims")
	}

	user := model.User{
		ID: uuidUser,
	}

	return &user, nil
}
