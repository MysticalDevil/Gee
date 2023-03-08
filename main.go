// Package main Application main entry
package main

import (
	"fmt"
	"gout"
	"html/template"
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

	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": FormatAsDate,
	})
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./static")

	stu1 := &student{Name: "Gout", Age: 20}
	stu2 := &student{Name: "Alpha", Age: 18}
	r.GET("/", func(c *gout.Context) {
		c.HTML(http.StatusOK, "css.tmpl", nil)
	})
	r.GET("/students", func(c *gout.Context) {
		c.HTML(http.StatusOK, "arr.tmpl", gout.H{
			"title":  "gout",
			"stuArr": [2]*student{stu1, stu2},
		})
	})

	r.GET("/date", func(c *gout.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gout.H{
			"title": "gout",
			"now":   time.Date(2022, 3, 9, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":8080")
}
