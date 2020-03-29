package main

/*
	Author : Iordanis Paschalidis
	Date   : 29/03/2020
*/

import (
	"fmt"
	"net/http"

	countriesCon "./controller/countries"
	countryCon "./controller/country"
	pconf "./lib/config"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	serverConf = pconf.GetAppConfig("./config/covid.json")
)

func country(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := countryCon.Perform(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
}

func countries(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := countriesCon.Perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)

	fmt.Println("server running at port " + serverConf.Server.Port)

	router.HandleFunc("/country", country).Methods("POST")
	router.HandleFunc("/countries", countries).Methods("GET")

	c := cors.New(cors.Options{
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	http.ListenAndServe(serverConf.Server.Port, handler)
}
