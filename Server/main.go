package main

import (
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
	"gopkg.in/mgo.v2"

	"practica1/controllers"
)

func main() {

	r := httprouter.New()

	uc := controllers.NewCarController(getSession())

	r.GET("/cars/delete/:carId", uc.DeleteCar )
	r.GET("/cars/car/:carId", uc.GetCar)
	r.GET("/cars", uc.GetAllCars)
	r.POST("/cars", uc.CreateCar)
	r.POST("/cars/:id", uc.UpdateCar )
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":9000", handler)

}

func getSession() *mgo.Session {
	// Mongo DB connection
	s, err := mgo.Dial(os.Getenv("MONGO_URI"))
	if err != nil {
		panic(err)
	}
	return s
}
