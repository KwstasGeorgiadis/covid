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

	allcountries "github.com/junkd0g/covid/controller/allcountries"
	compare "github.com/junkd0g/covid/controller/compare"
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
	POST request to /api/country
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
			"casesPerOneMillion": 102,
			"tests": 21298974,
    		"testsPerOneMillion": 64371
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
		"Endpoint /api/country called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	Get request to /api/countries with no parameters

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
				"casesPerOneMillion": 0.5,
				"tests": 48305,
            	"testsPerOneMillion": 1243
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
				"casesPerOneMillion": 2,
				"tests": 48305,
            	"testsPerOneMillion": 1243
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
		"Endpoint /api/countries called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	POST request to /api/sort endpoint

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
			"casesPerOneMillion": 2061,
			"tests": 21298974,
            "testsPerOneMillion": 64371
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
			"casesPerOneMillion": 2668,
			"tests": 21298974,
            "testsPerOneMillion": 64371
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
		"Endpoint /api/sort called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	Get request to /api/total with no parameters

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
		"Endpoint /api/total called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	Get request to /api/countries/all with no parameters

	Response:

	{
    	"countries": [
        	"Afghanistan",
        	"Albania",
        	"Algeria",
        	"Andorra",
        	"Angola",
        	"Anguilla",
       		"Antigua and Barbuda",
        	"Argentina",
        	"Armenia",
        	"Aruba",
        	"Australia",
        	"Austria",
        	"Azerbaijan",
        	"Bahamas",
        	"Bahrain",
        	"Bangladesh",
        	"Barbados",
        	"Belarus",
        	"Belgium",
        	"Belize",
        	"Benin",
			"Bermuda"
		]
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
		"Endpoint /api/countries/all called with response JSON body "+string(jsonBody), status, elapsed)
}

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
	applogger.LogHTTP("INFO", "main", "compareUniqueCasesHandle",
		"Endpoint /compare/cases/unique called with response JSON body "+string(jsonBody), status, elapsed)
}

//----------------------------------------------------------------------------------------
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
	Get request to /news with no parameters

	Response:

	{
    "data": [
        {
            "title": "UGC ने जारी किए नंबर, पूछ सकते हैं एडमिशन से लेकर एग्जाम तक अपने सवाल",
            "description": "कोरोना संक्रमण के चलते देश भर में  लॉकडाउन को एक बार फिर बढ़ा दिया गया है.",
            "url": "https://aajtak.intoday.in/education/story/ugc-direct-numbers-for-queries-helpline-for-ug-pg-students-tedu-1-1192009.html",
            "urlToImage": "https://smedia2.intoday.in/aajtak/images/stories/092019/3_1589811689_618x347.jpeg",
            "publishedAt": "2020-05-18T15:14:14Z",
            "content": "UGC helpline:"
        },
        {
            "title": "Karen who can't believe she has to wear a mask to enter a supermarket confronts store manager",
            "description": "A woman who called herself Shelley Lewis acted rude and arrogant toward Gelson's supermarket employees",
            "url": "https://boingboing.net/2020/05/18/karen-who-cant-believe-she-h.html",
            "urlToImage": "https://i1.wp.com/media.boingboing.net/wp-content/uploads/2020/05/mask-1.jpg?fit=700%2C503&ssl=1",
            "publishedAt": "2020-05-18T15:12:19Z",
            "content": ""
        }
		]
	}
*/
func newsHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := crnews.Perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "newsHandle",
		"Endpoint /compare/percent called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	Get request to /news/vaccine with no parameters

	Response:

	{
    "data": [
        {
            "title": "UGC ने जारी किए नंबर, पूछ सकते हैं एडमिशन से लेकर एग्जाम तक अपने सवाल",
            "description": "कोरोना संक्रमण के चलते देश भर में  लॉकडाउन को एक बार फिर बढ़ा दिया गया है.",
            "url": "https://aajtak.intoday.in/education/story/ugc-direct-numbers-for-queries-helpline-for-ug-pg-students-tedu-1-1192009.html",
            "urlToImage": "https://smedia2.intoday.in/aajtak/images/stories/092019/3_1589811689_618x347.jpeg",
            "publishedAt": "2020-05-18T15:14:14Z",
            "content": "UGC helpline:"
        },
        {
            "title": "Karen who can't believe she has to wear a mask to enter a supermarket confronts store manager",
            "description": "A woman who called herself Shelley Lewis acted rude and arrogant toward Gelson's supermarket employees",
            "url": "https://boingboing.net/2020/05/18/karen-who-cant-believe-she-h.html",
            "urlToImage": "https://i1.wp.com/media.boingboing.net/wp-content/uploads/2020/05/mask-1.jpg?fit=700%2C503&ssl=1",
            "publishedAt": "2020-05-18T15:12:19Z",
            "content": ""
        }
		]
	}
*/
func newsVaccineHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := crnews.PerformVaccineNews()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "newsVaccineHandle",
		"Endpoint /compare/percent called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	Get request to /news/treatment with no parameters

	Response:

	{
    "data": [
        {
            "title": "UGC ने जारी किए नंबर, पूछ सकते हैं एडमिशन से लेकर एग्जाम तक अपने सवाल",
            "description": "कोरोना संक्रमण के चलते देश भर में  लॉकडाउन को एक बार फिर बढ़ा दिया गया है.",
            "url": "https://aajtak.intoday.in/education/story/ugc-direct-numbers-for-queries-helpline-for-ug-pg-students-tedu-1-1192009.html",
            "urlToImage": "https://smedia2.intoday.in/aajtak/images/stories/092019/3_1589811689_618x347.jpeg",
            "publishedAt": "2020-05-18T15:14:14Z",
            "content": "UGC helpline:"
        },
        {
            "title": "Karen who can't believe she has to wear a mask to enter a supermarket confronts store manager",
            "description": "A woman who called herself Shelley Lewis acted rude and arrogant toward Gelson's supermarket employees",
            "url": "https://boingboing.net/2020/05/18/karen-who-cant-believe-she-h.html",
            "urlToImage": "https://i1.wp.com/media.boingboing.net/wp-content/uploads/2020/05/mask-1.jpg?fit=700%2C503&ssl=1",
            "publishedAt": "2020-05-18T15:12:19Z",
            "content": ""
        }
		]
	}
*/
func newsTreatmentHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := crnews.PerformTreatmentNews()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "newsTreatmentHandle",
		"Endpoint /compare/percent called with response JSON body "+string(jsonBody), status, elapsed)
}

func newsAllHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := crnews.PerformAll()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "newsAllHandle",
		"Endpoint /compare/percent called with response JSON body "+string(jsonBody), status, elapsed)
}

func worldHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := worldct.Perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "worldHandle",
		"Endpoint /compare/percent called with response JSON body "+string(jsonBody), status, elapsed)
}

func hotspotHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := hotspot.Perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "worldHandle",
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

	router.HandleFunc("/api/hotspot", hotspotHandle).Methods("GET")
	router.HandleFunc("/api/world", worldHandle).Methods("GET")
	router.HandleFunc("/api/news", newsHandle).Methods("GET")
	router.HandleFunc("/api/news/all", newsAllHandle).Methods("GET")
	router.HandleFunc("/api/news/vaccine", newsVaccineHandle).Methods("GET")
	router.HandleFunc("/api/news/treatment", newsTreatmentHandle).Methods("GET")
	router.HandleFunc("/api/country", country).Methods("POST")
	router.HandleFunc("/api/countries", countries).Methods("GET")
	router.HandleFunc("/api/countries/all", allCountriesHandle).Methods("GET")
	router.HandleFunc("/api/sort", sort).Methods("POST")
	router.HandleFunc("/api/total", totalStatistics).Methods("GET")
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
