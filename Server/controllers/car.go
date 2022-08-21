package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practica1/models"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type CarController struct {
	session *mgo.Session
}

func NewCarController(s *mgo.Session) *CarController {
	return &CarController{s}
}

func (uc CarController) GetAllCars(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	enableCors(&w)
	var cars []bson.M
	if err := uc.session.DB("so1p1").C("cars").Find(bson.M{}).All(&cars); err != nil {
		w.WriteHeader(404)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json, _ := json.Marshal(cars)
	fmt.Fprintf(w, "%s\n", json)
}

func (uc CarController) GetCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	enableCors(&w)
	id := p.ByName("carId")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.Car{}

	if err := uc.session.DB("so1p1").C("cars").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		// Inserting log
		uc.generateLog(fmt.Sprintf("Error, car with ID: %v doesn't exist.", id))
		return
	}

	uj, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "%s\n", uj)
}
func (uc CarController) CreateCar(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	enableCors(&w)

	u := models.Car{}

	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	uc.session.DB("so1p1").C("cars").Insert(u)
	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
	uc.generateLog("Car data succesfully saved.")
}

func (uc CarController) DeleteCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	enableCors(&w)

	id := p.ByName("carId")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("so1p1").C("cars").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted car", oid, "\n")
	uc.generateLog(fmt.Sprintf("Car with ID: %v succesfully deleted.", id))

}

func (uc CarController) UpdateCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	enableCors(&w)

	id := p.ByName("id")

	newCar := models.Car{}

	json.NewDecoder(r.Body).Decode(&newCar)

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)


	if err := uc.session.DB("so1p1").C("cars").Update(bson.M{"_id": oid}, newCar); err != nil {
		fmt.Printf("update fail %v\n", err)
		w.WriteHeader(404)
		uc.generateLog(fmt.Sprintf("Error, car with ID: %v doesn't exist.", id))
		return
	}

	uj, err := json.Marshal(newCar)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
	uc.generateLog(fmt.Sprintf("Car with ID: %v successfully updated.", id))


}

func (uc CarController) generateLog(description string) {
	newLog := models.Log{}
	newLog.Id = bson.NewObjectId()
	newLog.Description = description
	dt := time.Now()
	newLog.LogDate = dt.String()
	uc.session.DB("so1p1").C("logs").Insert(newLog)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
