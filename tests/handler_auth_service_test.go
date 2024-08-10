package tests

import (
	"fmt"
	"gin-be/internal/ent"
	"gin-be/internal/ent/enttest"
	"gin-be/internal/ent/migrate"
	"gin-be/internal/service"
	"gin-be/internal/tool"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

// Test for checking each services
func TestAuthServiceHandler(t *testing.T) {
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)
	path := "../.env.test"
	_ = tool.NewEnv(&path)

	// opts := []enttest.Option{
	// 	enttest.WithOptions(ent.Log(t.Log)),
	// 	enttest.WithMigrateOptions(migrate.WithGlobalUniqueID(true)),
	// }
	client := enttest.Open(t, "sqlite3", "file:ent1?cache=shared&mode=memory&_fk=1")
	defer client.Close()

	// clientTx, err := client.Tx(ctx)

	// require.NoError(t, err)

	// Create an SQLite memory database and generate the schema.
	errr := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(false))
	require.NoError(t, errr)

	initDataUser(ctx, client)

	tests := []struct {
		name   string                                               // name of the test
		expect func(*gin.Context, *testing.T, *ent.Client, *ent.Tx) // actual exceptions
	}{
		{
			name: "auth|RegisterUserByEmail",
			expect: func(ctx *gin.Context, t *testing.T, client *ent.Client, clientTx *ent.Tx) {
				_, err3 := service.RegisterUserByEmail(ctx, clientTx, "aku", "aku2@email.com", "085755519122", "qweasd")

				require.NoError(t, err3)

				_, err := service.RegisterUserByEmail(ctx, clientTx, "aku", "as.triarjo@gmail.com", "085755519126", "qweasd")
				if err != nil {
					assert.Equal(t, "ent: constraint failed: UNIQUE constraint failed: users.email", err.Error())
				}
				_, err2 := service.RegisterUserByEmail(ctx, clientTx, "aku", "aku3@email.com", "085755519123", "qweasd")
				if err2 != nil {
					assert.Equal(t, "ent: constraint failed: UNIQUE constraint failed: users.phone", err2.Error())
				}
				clientTx.Commit()
			},
		},
		{
			name: "auth|LoginUserByEmail",
			expect: func(ctx *gin.Context, t *testing.T, client *ent.Client, clientTx *ent.Tx) {
				_, err := service.LoginUserByEmail(ctx, client, "s.triarjo@gmail.com", "qweasd")
				require.NoError(t, err)
				_, err2 := service.LoginUserByEmail(ctx, client, "s.triarjo@gmail.com", "qweasde")
				if err2 != nil {
					assert.Equal(t, "password does not match", err2.Error())
				}
				_, err3 := service.LoginUserByEmail(ctx, client, "akuakuq@email.com", "qweasde")
				if err3 != nil {
					assert.Equal(t, "user not found", err3.Error())
				}

			},
		},
		{
			name: "auth|GetUserById",
			expect: func(ctx *gin.Context, t *testing.T, client *ent.Client, clientTx *ent.Tx) {
				entity, err := service.LoginUserByEmail(ctx, client, "s.triarjo@gmail.com", "qweasd")
				require.NoError(t, err)
				user, err := service.GetUserById(ctx, client, entity.ID)
				require.NoError(t, err)
				if err == nil {
					assert.Equal(t, entity.ID, user.ID)
					assert.Equal(t, entity.Email, user.Email)
				}
				_, err2 := service.GetUserById(ctx, client, uuid.New())
				if err2 != nil {
					assert.Equal(t, "User not found", err2.Error())
				}
			},
		},
		{
			name: "auth|CheckExistingPhone",
			expect: func(ctx *gin.Context, t *testing.T, client *ent.Client, clientTx *ent.Tx) {
				entity, err := service.CheckExistingPhone(ctx, client, "085755519123")
				require.NoError(t, err)
				assert.Equal(t, true, entity)
				check, err := service.CheckExistingPhone(ctx, client, "+6285755519123")
				assert.Equal(t, false, check)
			},
		},
		{
			name: "auth|CheckExistingEmail",
			expect: func(ctx *gin.Context, t *testing.T, client *ent.Client, clientTx *ent.Tx) {
				entity, err := service.CheckExistingEmail(ctx, client, "s.triarjo@gmail.com")
				require.NoError(t, err)
				assert.Equal(t, true, entity)
				check, err := service.CheckExistingEmail(ctx, client, "s.triar@live.com")
				assert.Equal(t, false, check)
			},
		},
		// more tests ...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clientTx2, err := client.Tx(ctx)
			require.NoError(t, err)
			// Run the actual exceptions.
			tt.expect(ctx, t, client, clientTx2)

		})
	}
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}
