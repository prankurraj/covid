package main

import (
	"github.com/labstack/echo"
	"github.com/covid/src/handler"
	"fmt"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo"
	"context"
	"log"
)

//main function
func main() {
	fmt.Printf("hello")
	e := echo.New()
	e.Use(middleware.Logger())

	clientOptions := options.Client().
		ApplyURI("mongodb+srv://covid:covid@cluster0.v7qux.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	client.Database("covid")

	h := &handler.Handler{DB: client.Database("covid")}

	// Routes
	e.GET("/refresh", h.Covid)
	e.GET("/getCases", h.GetCasesFromGeoLocation)
	e.GET("/getCasesFromState", h.GetCasesFromStateCode)


	e.Logger.Fatal(e.Start(":8000"))

}