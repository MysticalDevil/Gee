// Package main Application main entry
package main

import (
	"fmt"
	"gout"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gout.Default()

	r.GET("/", func(c *gout.Context) {
		c.String(http.StatusOK, "Hello Gout\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gout.Context) {
		names := []string{"gout"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":8080")
}
