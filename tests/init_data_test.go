package tests

import (
	"gin-be/internal/ent"
	"gin-be/internal/service"

	"github.com/gin-gonic/gin"
)

// TDODO: insert data dummy for testing
func initDataUser(ctx *gin.Context, client *ent.Client) {
	clientTx, _ := client.Tx(ctx)
	service.RegisterUserByEmail(ctx, clientTx, "aku", "s.triarjo@gmail.com", "085755519123", "qweasd")
	clientTx.Commit()
}
