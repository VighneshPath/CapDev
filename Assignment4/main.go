package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func getData(c *gin.Context) {
	if os.Getenv("TATKAL_BUSY_TIME") == "true" {
		c.IndentedJSON(http.StatusPreconditionFailed, "busy")
		return
	}
	c.IndentedJSON(http.StatusOK, "hello")

}

func main() {
	router := gin.Default()

	router.GET("/", getData)

	router.Run(":8080")
}
