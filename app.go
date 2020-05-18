package main

/*
	Author : Iordanis Paschalidis
	Date   : 29/03/2020
*/

import (
	"fmt"
	"net/http"
	"time"

	statisticscon "github.com/junkd0g/covid/controller/statistics"
	"github.com/junkd0g/covid/lib/applogger"
	"github.com/junkd0g/covid/lib/news"

	allcountries "github.com/junkd0g/covid/controller/allcountries"
	compare "github.com/junkd0g/covid/controller/compare"
	countriescon "github.com/junkd0g/covid/controller/countries"
	countrycon "github.com/junkd0g/covid/controller/country"
	totalcon "github.com/junkd0g/covid/controller/totalcon"

	"github.com/gorilla/mux"
	sortcon "github.com/junkd0g/covid/controller/sort"
	pconf "github.com/junkd0g/covid/lib/config"
	"github.com/rs/cors"
)

var (
	//reads the config and creates a AppConf struct
	serverConf = pconf.GetAppConfig()
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
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := countrycon.Perform(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "country",
		"Endpoint /country called with response JSON body "+string(jsonBody), status, elapsed)
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
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := countriescon.Perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "countries",
		"Endpoint /countries called with response JSON body "+string(jsonBody), status, elapsed)
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
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := sortcon.Perform(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "sort",
		"Endpoint /sort called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	CHECK THIS ENDPOINT LOOKS THAT IT IS MISSING

*/
func statistics(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := statisticscon.Perform(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "statistics",
		"Endpoint /stats called with response JSON body "+string(jsonBody), status, elapsed)
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
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := totalcon.Perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "totalStatistics",
		"Endpoint /total called with response JSON body "+string(jsonBody), status, elapsed)
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
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := allcountries.Perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "allCountriesHandle",
		"Endpoint /countries called with response JSON body "+string(jsonBody), status, elapsed)
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
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := compare.Perform(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "compareHandle",
		"Endpoint /compare called with response JSON body "+string(jsonBody), status, elapsed)
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
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := compare.PerformFromFirstDeath(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "compareFromFirstDeathHandle",
		"Endpoint /compare/firstdeath called with response JSON body "+string(jsonBody), status, elapsed)

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
func comparePerDayDeathHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := compare.PerformPerDayDeath(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "comparePerDayDeathHandle",
		"Endpoint /compare/perday called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	POST request to /compare/percent endpoint

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
func comparePercantagePerDayDeathHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := compare.PerformPercentangePerDayDeath(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "comparePercantagePerDayDeathHandle",
		"Endpoint /compare/percent called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	POST request to /compare/recovery endpoint

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
func compareRecoveryHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := compare.PerformCompareRecorey(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "comparePercantagePerDayDeathHandle",
		"Endpoint /compare/percent called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	POST request to /compare/cases endpoint

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
func compareCasesHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := compare.PerformCompareCases(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "comparePercantagePerDayDeathHandle",
		"Endpoint /compare/percent called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	POST request to /compare/cases/unique endpoint

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
func compareUniqueCasesHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := compare.PerformCompareUniquePerDayCases(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "comparePercantagePerDayDeathHandle",
		"Endpoint /compare/percent called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	Running the server in port 9080 (getting the value from ./config/covid.json )

	"server" : {
                "port" : ":9080"
    },

	Endpoints:
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
			/compare/recovery
			/compare/cases
			/compare/unique
*/

func main() {
	news.News()
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
	router.HandleFunc("/compare/perday", comparePerDayDeathHandle).Methods("POST")
	router.HandleFunc("/compare/percent", comparePercantagePerDayDeathHandle).Methods("POST")
	router.HandleFunc("/compare/recovery", compareRecoveryHandle).Methods("POST")
	router.HandleFunc("/compare/cases", compareCasesHandle).Methods("POST")
	router.HandleFunc("/compare/cases/unique", compareUniqueCasesHandle).Methods("POST")

	c := cors.New(cors.Options{
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	http.ListenAndServe(port, handler)
}
