// Package main Application main entry
package main

import (
	"gout"
	"net/http"
)

func main() {
	r := gout.New()

	r.GET("/", func(c *gout.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gout</h1>")
	})

	r.GET("/hello", func(c *gout.Context) {
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gout.Context) {
		c.JSON(http.StatusOK, gout.H{
			"username": c.PostFrom("username"),
			"password": c.PostFrom("password"),
		})
	})

	r.Run(":8080")
}
