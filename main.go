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
		// expect /hello?name=gout
		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.GET("/hello/:name", func(c *gout.Context) {
		// expect /hello/gout
		c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *gout.Context) {
		c.JSON(http.StatusOK, gout.H{
			"filepath": c.Param("filepath"),
		})
	})

	r.Run(":8080")
}
