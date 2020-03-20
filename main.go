package main

import (
	"fmt"
	"gee/gee"
	"html/template"
	"net/http"
	"time"
)

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := gee.Default()
	r.Use(gee.Logger()) // global middleware
	r.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	r.GET("/", func(c *gee.Context) {
		c.String(http.StatusOK, "Hello Geek\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *gee.Context) {
		names := []string{"geek"}
		c.String(http.StatusOK, names[100])
	})
	r.GET("/date", func(c *gee.Context) {
		c.HTML(http.StatusOK, "custom_func.tmpl", gee.H{
			"title": "gee",
			"now":   time.Date(2020, 3, 20, 0, 0, 0, 0, time.UTC),
		})
	})

	r.Run(":9999")
}
