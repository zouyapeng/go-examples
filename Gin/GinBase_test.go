package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRootRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "Hello, I'm Gin\n", w.Body.String())
}

func TestJsonRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/json", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	body := map[string]interface{}{}
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, map[string]interface{}{"message": "Hello, I'm Gin"}, body)
}

func TestGETRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/api/v1/test", nil)
	req.Header.Set("token", "test-token")
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	body := map[string]interface{}{}
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, map[string]interface{}{"status": "ok"}, body)
}

func TestPOSTRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	jsonValue, _ := json.Marshal(map[string]interface{}{"name": "test-1111", "tests": []string{"111", "222"}, "count": 10, "create_time": time.Now().Add(-10 * time.Minute), "update_time": time.Now()})
	req, _ := http.NewRequest("POST", "/api/v1/test", bytes.NewBuffer(jsonValue))
	req.Header.Set("token", "test-token")
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	body := map[string]interface{}{}
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, map[string]interface{}{"status": "ok"}, body)
}

func TestPUTRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/api/v1/test/1", nil)
	req.Header.Set("token", "test-token")
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	body := map[string]interface{}{}
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, map[string]interface{}{"status": "ok"}, body)
}

func TestDeleteRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/api/v1/test/1", nil)
	req.Header.Set("token", "test-token")
	req.Header.Set("Content-Type", "application/json")
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	body := map[string]interface{}{}
	_ = json.Unmarshal(w.Body.Bytes(), &body)
	assert.Equal(t, map[string]interface{}{"status": "ok"}, body)
}
