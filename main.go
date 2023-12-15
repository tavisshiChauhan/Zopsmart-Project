package main

import (
	"Go-Lang-Zopsmart/datastore"
	"Go-Lang-Zopsmart/handler"

	"gofr.dev/pkg/gofr"
)

func main() {
	
	app := gofr.New()

	s := datastore.New()
	h := handler.New(*s)

	app.GET("/students/{ID}", h.GetByID)   
	app.POST("/students", h.Create)
	app.PUT("/students/{ID}", h.Update)    
	app.DELETE("/students/{ID}", h.Delete) 

	app.Server.HTTP.Port = 9092
	app.Start()
}
