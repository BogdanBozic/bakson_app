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

	r.Run(":3000")
}

func HelloHandler(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World!",
	})
}

func CurrentTime(c *gin.Context) {

	time := time2.Now()
	fmt.Println("Time: ", time)
	timeString := time.Format("2006-01-02 15:04:05")
	c.String(http.StatusOK, timeString)
}
