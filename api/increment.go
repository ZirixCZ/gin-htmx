package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// create class count that will be incremented
var count int

func Increment(c *gin.Context) {
	count++
	c.String(200, strconv.Itoa(count))
}
