package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"practica1/models"

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

func (uc CarController) GetCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	u := models.Car{}

	if err := uc.session.DB("so1-practica1").C("cars").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
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
	u := models.Car{}

	json.NewDecoder(r.Body).Decode(&u)
	u.Id = bson.NewObjectId()
	uc.session.DB("so1-practica1").C("cars").Insert(u)
	uj, err := json.Marshal(u)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc CarController) DeleteCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	if err := uc.session.DB("so1-practica1").C("cars").RemoveId(oid); err != nil {
		w.WriteHeader(404)
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted car", oid, "\n")

}

func (uc CarController) UpdateCar(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	id := p.ByName("id")

	newCar := models.Car{}

	json.NewDecoder(r.Body).Decode(&newCar)

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)


	if err := uc.session.DB("so1-practica1").C("cars").Update(bson.M{"_id": oid}, newCar); err != nil {
		fmt.Printf("update fail %v\n", err)
		w.WriteHeader(404)
		return
	}

	uj, err := json.Marshal(newCar)

	if err != nil {
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)


}
