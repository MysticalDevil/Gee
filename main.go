// Package main Application main entry
package main

import (
	"gout"
	"log"
	"net/http"
	"time"
)

func main() {
	r := gout.New()
	r.Use(gout.Logger())

	r.GET("/index", func(c *gout.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gout.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gout</h1>")
		})

		v1.GET("/hello", func(c *gout.Context) {
			// expect /hello?name="gout"
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}

	v2 := r.Group("/v2")
	v2.Use(onlyForV2())
	{
		v2.GET("/hello/:name", func(c *gout.Context) {
			// expect /hello/gout
			c.String(http.StatusOK, "Hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *gout.Context) {
			c.JSON(http.StatusOK, gout.H{
				"username": c.PostFrom("username"),
				"password": c.PostFrom("password"),
			})
		})
	}

	r.Run(":8080")
}

func onlyForV2() gout.HandlerFunc {
	return func(c *gout.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(http.StatusInternalServerError, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}
