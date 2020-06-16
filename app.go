package main

/*
	Author : Iordanis Paschalidis
	Date   : 29/03/2020
*/

import (
	"fmt"
	"net/http"

	"github.com/junkd0g/covid/controller/allcountries"
	comparectl "github.com/junkd0g/covid/controller/compare"
	continentctl "github.com/junkd0g/covid/controller/continent"
	countriescon "github.com/junkd0g/covid/controller/countries"
	countrycon "github.com/junkd0g/covid/controller/country"
	hotspot "github.com/junkd0g/covid/controller/hotspot"
	crnews "github.com/junkd0g/covid/controller/news"
	sortcon "github.com/junkd0g/covid/controller/sort"
	totalcon "github.com/junkd0g/covid/controller/totalcon"
	worldct "github.com/junkd0g/covid/controller/world"

	"github.com/gorilla/mux"
	pconf "github.com/junkd0g/covid/lib/config"
	"github.com/rs/cors"
)

var (
	//reads the config and creates a AppConf struct
	serverConf = pconf.GetAppConfig()
)

/*
	Running the server in port 9080 (getting the value from ./config/covid.json )

	"server" : {
        "port" : ":9080"
    },

	Endpoints:
		GET:
			/api/hotspot
            /api/world
            /api/continent
			/api/total
			/api/countries
			/api/countries/all
			/api/news
			/api/news/all
			/api/news/vaccine
			/api/news/treatment
		POST
			/api/country
			/api/sort
			/api/stats
			/api/compare/all

*/

func main() {
	router := mux.NewRouter().StrictSlash(true)
	port := serverConf.Server.Port
	fmt.Println("server running at port " + port)

	router.HandleFunc("/api/hotspot/{days}", hotspot.Handle).Methods("GET")
	router.HandleFunc("/api/world", worldct.Handle).Methods("GET")
	router.HandleFunc("/api/continent", continentctl.Handle).Methods("GET")
	router.HandleFunc("/api/news/all", crnews.NewsAllHandle).Methods("GET")
	router.HandleFunc("/api/country", countrycon.Handle).Methods("POST")
	router.HandleFunc("/api/countries", countriescon.Handle).Methods("GET")
	router.HandleFunc("/api/countries/all", allcountries.Handle).Methods("GET")
	router.HandleFunc("/api/sort", sortcon.Handle).Methods("POST")
	router.HandleFunc("/api/total", totalcon.Handle).Methods("GET")
	router.HandleFunc("/api/compare/all", comparectl.Handle).Methods("POST")

	c := cors.New(cors.Options{
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	http.ListenAndServe(port, handler)
}
