package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"

	"practica1/controllers"
)

func main() {

	r := httprouter.New()

	uc := controllers.NewCarController(getSession())

	r.GET("/car/:id", uc.GetCar)
	r.POST("/car", uc.CreateCar)
	r.DELETE("/car/:id", uc.DeleteCar )
	r.PUT("/car/:id", uc.UpdateCar )

	http.ListenAndServe("localhost:9000", r)

}

func getSession() *mgo.Session {
	// Mongo DB connection
	s, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	return s
}
