package main

import (
	"crown/crown"
	"net/http"
)

func main() {
	r := crown.New()
	
	r.GET("/", func(c *crown.Context) {
		c.JSON(http.StatusOK, crown.H{
			"msg": "context",
		})
	})

	r.Run(":8000")
}
