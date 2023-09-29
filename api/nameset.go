package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Nameset(c *gin.Context) {
	name := c.PostForm("name")

	if name == "" {
		name = "anon"
	}

	fmt.Println("here", name)
	c.HTML(200, "layout.html", gin.H{
		"name": name,
	})
	c.HTML(200, "orbit.html", gin.H{
		"name": name,
	})
}
