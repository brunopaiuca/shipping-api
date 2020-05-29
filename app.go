// app.go

package main

import (
	"database/sql"
	"fmt"
	"log"

	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

var a App

func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

	var err error
	a.DB, err = sql.Open("postgres", connectionString)
	if err != nil {
		log.Fatal(err)
	}

	//router := mux.NewRouter()
	//a.Router = router

	a.Router = mux.NewRouter()
	a.initializeRoutes()

}

func (a *App) Run(addr string) {}

func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func (a *App) ShippingQuotation(w http.ResponseWriter, r *http.Request) {

	var body map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	ship_id := body["ship_id"].(string)
	destination_zipcode := body["destination_zipcode"].(string)
	weight_in_kilogramas := body["weight_in_kilogramas"].(string)

	s := Shipping{Weight_in_kilogramas: weight_in_kilogramas, Ship_id: ship_id, Destination_zipcode: destination_zipcode}

	if err := s.getShippingDataForQuotation(a.DB); err != nil {
		switch err {
		case sql.ErrNoRows:
			//respondWithError(w, http.StatusNotFound, err.Error())
			respondWithError(w, http.StatusNotFound, "Either this type of shipping or no one of them do not cover the destination Zipcode.")
		default:
			respondWithError(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	respondWithJSON(w, http.StatusOK, s)
}

func (a *App) ShippingFullQuotation(w http.ResponseWriter, r *http.Request) {

	var body map[string]interface{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	defer r.Body.Close()

	weight_in_kilogramas := body["weight_in_kilogramas"].(string)
	destination_zipcode := body["destination_zipcode"].(string)

	shippings, err := getShippingDataForFullQuotation(a.DB, weight_in_kilogramas, destination_zipcode)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		//respondWithError(w, http.StatusInternalServerError, "a")
		return
	}

	//var response []byte
	//response, err := json.Marshal(shippings)
	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(200)
	//w.Write(response)
	//return
	respondWithJSON(w, http.StatusOK, shippings)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/shipping_fullquotation", a.ShippingFullQuotation).Methods("POST")
	a.Router.HandleFunc("/shipping_quotation", a.ShippingQuotation).Methods("POST")
	a.Router.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8010", a.Router))
}
