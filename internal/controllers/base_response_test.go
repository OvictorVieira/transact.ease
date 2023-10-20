package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewSuccessResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/success", func(c *gin.Context) {
		NewSuccessResponse(c, 200, "OK", gin.H{"foo": "bar"})
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/success", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, `{"status":200,"message":"OK","data":{"foo":"bar"}}`, w.Body.String())
}

func TestNewErrorResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/error", func(c *gin.Context) {
		NewErrorResponse(c, 500, "Internal Server Error")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/error", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 500, w.Code)
	assert.JSONEq(t, `{"status":500,"message":"Internal Server Error"}`, w.Body.String())
}

func TestNewAbortResponse(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	router.GET("/abort", func(c *gin.Context) {
		NewAbortResponse(c, "Unauthorized")
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/abort", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 401, w.Code)
	assert.JSONEq(t, `{"status":401,"message":"Unauthorized"}`, w.Body.String())
}
