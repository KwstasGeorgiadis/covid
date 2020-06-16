package main

/*
	Author : Iordanis Paschalidis
	Date   : 29/03/2020
*/

import (
	"fmt"
	"net/http"
	"time"

	"github.com/junkd0g/covid/lib/applogger"

	"github.com/junkd0g/covid/controller/allcountries"
	compare "github.com/junkd0g/covid/controller/compare"
	continentct "github.com/junkd0g/covid/controller/continent"
	countriescon "github.com/junkd0g/covid/controller/countries"
	countrycon "github.com/junkd0g/covid/controller/country"
	hotspot "github.com/junkd0g/covid/controller/hotspot"
	crnews "github.com/junkd0g/covid/controller/news"
	totalcon "github.com/junkd0g/covid/controller/totalcon"
	worldct "github.com/junkd0g/covid/controller/world"

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
	POST request to /api/compare endpoint

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
		"Endpoint /api/compare called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	POST request to /api/compare/firstdeath endpoint

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
		"Endpoint /api/compare/firstdeath called with response JSON body "+string(jsonBody), status, elapsed)

}

/*
	POST request to /api/compare/perday endpoint

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
		"Endpoint /api/compare/perday called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	POST request to /api/compare/recovery endpoint

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
	applogger.LogHTTP("INFO", "main", "compareRecoveryHandle",
		"Endpoint /api/compare/recovery called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	POST request to /api/compare/cases endpoint

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
	applogger.LogHTTP("INFO", "main", "compareCasesHandle",
		"Endpoint /api/compare/cases called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	POST request to /api/compare/cases/unique endpoint

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
	applogger.LogHTTP("INFO", "main", "compareUniqueCasesHandle",
		"Endpoint /compare/cases/unique called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	POST request to /api/compare/all endpoint

	Request:

	{
		"countryOne" : "Spain",
		"countryTwo" : "Italy"
	}

	Response

	{
    "countryOne": {
        "country": "Spain",
        "dataDeaths": [
            28,
            27888,
            27940,
            28628,
            28678,
            28752
        ],
        "dataDeathsFromFirst": [
            28,
            27888,
            27940,
            28628,
            28678,
            28752
        ],
        "dataDeathsPerDay": [
            110,
            52,
            688,
            50,
            74
        ],
        "dataRecoverd": [
            150376,
            150376,
            150376,
            150376,
            150376
        ],
        "dataCases": [
            15,
            32,
            45,
            84,
            120,
            165,
            222,
            259,
        ],
        "dataCasesFromFirst": [
            159,
            294,
            394,
            334,
            318,
            332,
            240
        ]
    },
    "countryTwo": {
        "country": "Spain",
        "dataDeaths": [
            0,
            0,
            0,
            0,
            0,
        ],
        "dataDeathsFromFirst": [
            33530,
            33601,
            33689,
            33774,
            33846,
            33899
        ],
        "dataDeathsPerDay": [
            1,
            1,
            1,
            4,
            3,
            2,
        ],
        "dataRecoverd": [
            0,
            1,
            1,
            1,
            2,
            3,
            45,
            46,
            46,
        ],
        "dataCases": [
            232664,
            232997,
            233197,
            233515,
            233836,
            234013,
            234531,
            234801,
            234998
        ],
        "dataCasesFromFirst": [
            200,
            318,
            321,
            177,
            518,
            270,
            197
        ]
    }
}
}
*/
func compareAllHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := compare.PerformAll(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "compareAllHandle",
		"Endpoint /compare/percent called with response JSON body "+string(jsonBody), status, elapsed)
}

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
			/api/compare
			/api/compare/firstdeath
			/api/compare/perday
			/api/compare/recovery
			/api/compare/cases
			/api/compare/cases/unique
			/api/compare/all

*/

func main() {
	router := mux.NewRouter().StrictSlash(true)
	port := serverConf.Server.Port
	fmt.Println("server running at port " + port)

	router.HandleFunc("/api/hotspot/{days}", hotspot.Handle).Methods("GET")
	router.HandleFunc("/api/world", worldct.Handle).Methods("GET")
	router.HandleFunc("/api/continent", continentct.Handle).Methods("GET")
	router.HandleFunc("/api/news", crnews.NewsHandle).Methods("GET")
	router.HandleFunc("/api/news/all", crnews.NewsAllHandle).Methods("GET")
	router.HandleFunc("/api/news/vaccine", crnews.NewsVaccineHandle).Methods("GET")
	router.HandleFunc("/api/news/treatment", crnews.NewsTreatmentHandle).Methods("GET")
	router.HandleFunc("/api/country", countrycon.Handle).Methods("POST")
	router.HandleFunc("/api/countries", countriescon.Handle).Methods("GET")
	router.HandleFunc("/api/countries/all", allcountries.Handle).Methods("GET")
	router.HandleFunc("/api/sort", sortcon.Handle).Methods("POST")
	router.HandleFunc("/api/total", totalcon.Handle).Methods("GET")
	router.HandleFunc("/api/compare", compareHandle).Methods("POST")
	router.HandleFunc("/api/compare/firstdeath", compareFromFirstDeathHandle).Methods("POST")
	router.HandleFunc("/api/compare/perday", comparePerDayDeathHandle).Methods("POST")
	router.HandleFunc("/api/compare/recovery", compareRecoveryHandle).Methods("POST")
	router.HandleFunc("/api/compare/cases", compareCasesHandle).Methods("POST")
	router.HandleFunc("/api/compare/cases/unique", compareUniqueCasesHandle).Methods("POST")
	router.HandleFunc("/api/compare/all", compareAllHandle).Methods("POST")

	c := cors.New(cors.Options{
		AllowCredentials: true,
	})

	handler := c.Handler(router)
	http.ListenAndServe(port, handler)
}
