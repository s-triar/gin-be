package tests

import (
	"encoding/json"
	"fmt"
	"gin-be/internal/server"
	"gin-be/internal/tool"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"

	_ "github.com/mattn/go-sqlite3"
)

func TestHomeHandler(t *testing.T) {

	path := "../.env.test"
	_ = tool.NewEnv(&path)

	router := server.NewServer()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	expected_response := "{\"message\":\"Welcome to POS API\"}"
	assert.Equal(t, expected_response, w.Body.String())
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}

func TestHealthHandler(t *testing.T) {

	path := "../.env.test"
	_ = tool.NewEnv(&path)

	router := server.NewServer()
	// database.GetDB().Migrate()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)
	result := make(map[string]interface{})
	json.Unmarshal([]byte(w.Body.String()), &result)
	assert.Equal(t, "up", result["status"])
	fmt.Println("+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++")
}
