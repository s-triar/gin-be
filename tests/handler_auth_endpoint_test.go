package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gin-be/internal/controller"
	"gin-be/internal/database"
	"gin-be/internal/ent/enttest"
	"gin-be/internal/ent/migrate"
	"gin-be/internal/server"
	"gin-be/internal/tool"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

// Test for checking auth endpoint
func TestAutEndpointhHandler(t *testing.T) {
	w := httptest.NewRecorder()

	path := "../.env.test"
	_ = tool.NewEnv(&path)
	ctx, _ := gin.CreateTestContext(w)
	router := server.NewServer()

	client := enttest.Open(t, "sqlite3", "file:ent1?cache=shared&mode=memory&_fk=1")
	defer client.Close()
	// Create an SQLite memory database and generate the schema.
	errr := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(false))
	require.NoError(t, errr)

	database.GetDB().SwapDBEntClient(client)

	initDataUser(ctx, client)

	tests := []struct {
		name    string // name of the test
		purpose string
		method  string
		expect  func(*gin.Engine, string, string) // actual exceptions
	}{
		{
			name:    "/api/v1/auth/register",
			method:  "POST",
			purpose: "check duplicate phone number",
			expect: func(router *gin.Engine, method string, path string) {
				w := httptest.NewRecorder()
				body := controller.UserRegister{
					Fullname:        "Rob",
					Email:           "s.triarjo@gmail.com",
					Phone:           "085755519123",
					Password:        "qweasd",
					ConfirmPassword: "qweasd",
				}
				marshalled, _ := json.Marshal(body)
				req, _ := http.NewRequest(method, path, bytes.NewReader(marshalled))
				router.ServeHTTP(w, req)
				assert.Equal(t, 400, w.Code)
				result := make(map[string]interface{})
				json.Unmarshal([]byte(w.Body.String()), &result)
				expected_response := "ent: constraint failed: UNIQUE constraint failed: users.phone"
				assert.Equal(t, expected_response, result["error"])

			},
		},
		{
			name:    "/api/v1/auth/register",
			method:  "POST",
			purpose: "check duplicate email",
			expect: func(router *gin.Engine, method string, path string) {
				w := httptest.NewRecorder()
				body := controller.UserRegister{
					Fullname:        "Rob",
					Email:           "s.triarjo@gmail.com",
					Phone:           "085755519120",
					Password:        "qweasd",
					ConfirmPassword: "qweasd",
				}
				marshalled, _ := json.Marshal(body)
				req, _ := http.NewRequest(method, path, bytes.NewReader(marshalled))
				router.ServeHTTP(w, req)
				assert.Equal(t, 400, w.Code)
				result := make(map[string]interface{})
				json.Unmarshal([]byte(w.Body.String()), &result)
				expected_response := "ent: constraint failed: UNIQUE constraint failed: users.email"
				assert.Equal(t, expected_response, result["error"])

			},
		},
		{
			name:    "/api/v1/auth/register",
			method:  "POST",
			purpose: "check register success",
			expect: func(router *gin.Engine, method string, path string) {
				w := httptest.NewRecorder()
				body := controller.UserRegister{
					Fullname:        "Rob",
					Email:           "s.triarjo@live.com",
					Phone:           "085755519124",
					Password:        "qweasd",
					ConfirmPassword: "qweasd",
				}
				marshalled, _ := json.Marshal(body)
				req, _ := http.NewRequest(method, path, bytes.NewReader(marshalled))
				router.ServeHTTP(w, req)
				assert.Equal(t, 201, w.Code)
				result := make(map[string]interface{})
				json.Unmarshal([]byte(w.Body.String()), &result)
				expected_response := "User is created successfully"
				assert.Equal(t, expected_response, result["message"])

			},
		},
		{
			name:    "/api/v1/auth/register",
			method:  "POST",
			purpose: "check not match password and confirm password",
			expect: func(router *gin.Engine, method string, path string) {
				w := httptest.NewRecorder()
				body := controller.UserRegister{
					Fullname:        "Rob",
					Email:           "s.triarjo@gmail.com",
					Phone:           "085755519123",
					Password:        "qweasd",
					ConfirmPassword: "qweasde",
				}
				marshalled, _ := json.Marshal(body)
				req, _ := http.NewRequest(method, path, bytes.NewReader(marshalled))
				router.ServeHTTP(w, req)
				assert.Equal(t, 400, w.Code)
				result := make(map[string]interface{})
				json.Unmarshal([]byte(w.Body.String()), &result)
				expected_response := "Confirm Password does not match Password"
				assert.Equal(t, expected_response, result["error"])

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name+" - "+tt.purpose, func(t *testing.T) {

			tt.expect(router, tt.method, tt.name)

		})
	}

	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}
