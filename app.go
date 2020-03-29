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
	//reads the config and creates a AppConf struct
	serverConf = pconf.GetAppConfig("./config/covid.json")
)

/*
	POST request to /country wit no parameters

	Request:

	{
		"country" : "Greece"
	}

	Response

		{
		    "country": "Greece",
    		"cases": 1061,
    		"todayCases": 0,
    		"deaths": 37,
    		"todayDeaths": 5,
    		"recovered": 52,
    		"active": 972,
    		"critical": 66,
    		"casesPerOneMillion": 102
		}

*/
func country(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := countryCon.Perform(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
}

/*
	Get request to /countries wit no parameters

	Response:

	{
    	"data": [
        	{
            	"country": "Zimbabwe",
            	"cases": 7,
            	"todayCases": 0,
            	"deaths": 1,
            	"todayDeaths": 0,
            	"recovered": 0,
            	"active": 6,
            	"critical": 0,
            	"casesPerOneMillion": 0.5
        	},
        	{
            	"country": "Zambia",
            	"cases": 29,
            	"todayCases": 1,
            	"deaths": 0,
            	"todayDeaths": 0,
            	"recovered": 0,
            	"active": 29,
            	"critical": 0,
            	"casesPerOneMillion": 2
			}
		]
	}

*/
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
