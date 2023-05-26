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

/*
What I am doing here is starting up a new instance of gin for each function. This is rather bad practice
and there must be a way to start up a single instance for all the tests. Furthermore, tests like these
should be used with mocking methods, not directly on an instance like this. I opted for this approach
because it was the simplest one for me to get started with.
*/

// Test if the endpoint will return the expected JSON and 200 as response.
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

// Test if the format we are expecting will match the one that we will get. I wouldn't test the time itself as
// it can differ, so I thought this would be a better approach.
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
