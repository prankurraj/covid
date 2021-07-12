package main

import (
	"github.com/labstack/echo"
	"handler"
	"gopkg.in/mgo.v2"
)

//main function
func main() {
	// create a new echo instance
	e := echo.New()
	e.Logger.Fatal(e.Start(":8000"))

	db, err := mgo.Dial("localhost")
	if err != nil {
		e.Logger.Fatal(err)
	}

	h := &handler.Handler{DB: db}

	// Routes
	e.GET("/feed", h.Covid)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}