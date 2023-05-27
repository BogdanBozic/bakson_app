package main

import (
	"fmt"
	"net/http"
	time2 "time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/hello", HelloHandler)
	r.GET("/time", CurrentTime)

	err := r.Run(":3000")
	if err != nil {
		fmt.Println("Something went wrong. Debug me, please.")
	}
}

// HelloHandler : A simple test endpoint to get familiarized with gin and Go.
func HelloHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

// CurrentTime : Calculates current time on the server it is running on and returns it as string.
func CurrentTime(c *gin.Context) {

	time := time2.Now()
	fmt.Println("Time: ", time)
	timeString := time.Format("2006-01-02 15:04:05")
	c.String(http.StatusOK, timeString)
}
