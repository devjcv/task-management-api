package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type task struct {
	ID          int
	Summary     string
	Description string
}

func test(c *gin.Context) {
	var tasks = []task{
		{ID: 1, Summary: "Blue Train", Description: "John Coltrane"},
	}
	c.IndentedJSON(http.StatusOK, tasks)
}

func main() {
	router := gin.Default()
	router.GET("/test", test)

	router.Run("localhost:8000")
}
