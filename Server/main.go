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
	r.POST("/cars/filter", uc.GetFilteredCars )
	r.POST("/cars/update/:carId", uc.UpdateCar )
	r.GET("/cars/car/:carId", uc.GetCar)
	r.GET("/cars", uc.GetAllCars)
	r.POST("/cars", uc.CreateCar)
	handler := cors.Default().Handler(r)
	http.ListenAndServe(":9001", handler)

}

func getSession() *mgo.Session {
	// Mongo DB connection
	s, err := mgo.Dial(os.Getenv("MONGO_URI"))
	// s, err := mgo.Dial(os.Getenv("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	return s
}
