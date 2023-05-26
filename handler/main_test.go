package main

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	assert2 "github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"
)

func TestHelloHandler(t *testing.T) {

	router := gin.Default()

	router.GET("/hello", HelloHandler)
	req, err := http.NewRequest("GET", "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	expectedResponse := `{"message":"Hello World!"}`

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, expectedResponse, recorder.Body.String())
}

func TestCurrentTime(t *testing.T) {
	router := gin.Default()

	router.GET("/time", CurrentTime)

	req, err := http.NewRequest("GET", "/time", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()
	router.ServeHTTP(recorder, req)

	body := recorder.Body.String()
	match, _ := regexp.MatchString(`\d{4}-\d{2}-\d{2} \d{2}:\d{2}:\d{2}`, body)

	assert2.True(t, match, "Time not returned in proper format.")
}
