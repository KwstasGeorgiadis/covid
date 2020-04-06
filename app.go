package main

/*
	Author : Iordanis Paschalidis
	Date   : 29/03/2020
*/

import (
	"fmt"
	"net/http"

	statisticsCon "./controller/statistics"

	allcountries "./controller/allcountries"
	compare "./controller/compare"
	countriescon "./controller/countries"
	countryCon "./controller/country"
	totalcon "./controller/totalcon"

	sortCon "./controller/sort"
	pconf "./lib/config"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	//reads the config and creates a AppConf struct
	serverConf = pconf.GetAppConfig("./config/covid.json")
)

/*
	POST request to /country
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
	Get request to /countries with no parameters

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
	jsonBody, status := countriescon.Perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
}

/*
	POST request to /sort endpoint

	Request:

	{
		"type" : "deaths"
	}

	Response

	{
    	"data": [{
        	"country": "Italy",
            "cases": 124632,
            "todayCases": 4805,
            "deaths": 15362,
            "todayDeaths": 681,
            "recovered": 20996,
            "active": 88274,
            "critical": 3994,
            "casesPerOneMillion": 2061
        },
        {
            "country": "Spain",
            "cases": 124736,
            "todayCases": 5537,
            "deaths": 11744,
            "todayDeaths": 546,
            "recovered": 34219,
            "active": 78773,
            "critical": 6416,
            "casesPerOneMillion": 2668
		}]
	}

*/
func sort(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := sortCon.Perform(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
}

/*
	CHECK THIS ENDPOINT LOOKS THAT IT IS MISSING

*/
func statistics(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := statisticsCon.Perform(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
}

/*
	Get request to /total with no parameters

	Response:

	{
    	"todayPerCentOfTotalCases": 7,
    	"todayPerCentOfTotalDeaths": 6,
    	"totalCases": 1188489,
    	"totalDeaths": 64103,
    	"todayTotalCases": 71846,
    	"todayTotalDeaths": 4933
	}
*/
func totalStatistics(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := totalcon.Perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
}

/*
	Get request to /countries with no parameters

	Response:

	{
    	"data": [{
            "country": "Zimbabwe",
            "cases": 9,
            "todayCases": 0,
            "deaths": 1,
            "todayDeaths": 0,
            "recovered": 0,
            "active": 8,
            "critical": 0,
            "casesPerOneMillion": 0.6
        },
        {
            "country": "Zambia",
            "cases": 39,
            "todayCases": 0,
            "deaths": 1,
            "todayDeaths": 0,
            "recovered": 2,
            "active": 36,
            "critical": 0,
            "casesPerOneMillion": 2
		}]
	}

*/
func allCountriesHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := allcountries.Perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
}

/*
	POST request to /compare endpoint

	Request:

	{
		"countryOne" : "Spain",
		"countryTwo" : "Italy"
	}

	Response

	{
    "countryOne": {
        "country": "Spain",
        "data": [
            1,
            2,
            3,
            7,
            12428,
            13155,
            13915,
            14681
        ]
    },
    "countryTwo": {
        "country": "Italy",
        "data": [
            1,
            2,
            3,
            7,
            12428,
            13155,
            13915,
            14681
        ]
    }
}

*/
func compareHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := compare.Perform(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
}

/*
	POST request to /compare/firstdeath endpoint

	Request:

	{
		"countryOne" : "Spain",
		"countryTwo" : "Italy"
	}

	Response

	{
    "countryOne": {
        "country": "Spain",
        "data": [
            1,
            2,
            3,
            7,
            12428,
            13155,
            13915,
            14681
        ]
    },
    "countryTwo": {
        "country": "Italy",
        "data": [
            1,
            2,
            3,
            7,
            12428,
            13155,
            13915,
            14681
        ]
    }
}

*/
func compareFromFirstDeathHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := compare.PerformFromFirstDeath(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
}

/*
	POST request to /compare/perday endpoint

	Request:

	{
		"countryOne" : "Spain",
		"countryTwo" : "Italy"
	}

	Response

	{
    "countryOne": {
        "country": "Spain",
        "data": [
            1,
            2,
            3,
            7,
            12428,
            13155,
            13915,
            14681
        ]
    },
    "countryTwo": {
        "country": "Italy",
        "data": [
            1,
            2,
            3,
            7,
            12428,
            13155,
            13915,
            14681
        ]
    }
}

*/
func comparPerDayDeathHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := compare.PerformPerDayDeath(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
}

/*
	Running the server in port 9080 (getting the value from ./config/covid.json )

	"server" : {
                "port" : ":9080"
    },

	Endpoint:
		GET:
			/total
			/countries
			/countries/all
		POST
			/country
			/sort
			/stats
			/compare
			/compare/firstdeath
			/compare/perday
*/

func main() {
	router := mux.NewRouter().StrictSlash(true)
	port := serverConf.Server.Port

	fmt.Println("server running at port " + port)

	router.HandleFunc("/country", country).Methods("POST")
	router.HandleFunc("/countries", countries).Methods("GET")
	router.HandleFunc("/countries/all", allCountriesHandle).Methods("GET")
	router.HandleFunc("/sort", sort).Methods("POST")
	router.HandleFunc("/stats", statistics).Methods("POST")
	router.HandleFunc("/total", totalStatistics).Methods("GET")
	router.HandleFunc("/compare", compareHandle).Methods("POST")
	router.HandleFunc("/compare/firstdeath", compareFromFirstDeathHandle).Methods("POST")
	router.HandleFunc("/compare/perday", comparPerDayDeathHandle).Methods("POST")

	c := cors.New(cors.Options{
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	http.ListenAndServe(port, handler)
}
