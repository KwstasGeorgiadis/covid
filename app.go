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
	Get request to /api/news with no parameters

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
		"Endpoint /api/news called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	Get request to /api/news/vaccine with no parameters

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
		"Endpoint /api/news/vaccine called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	Get request to /api/news/treatment with no parameters

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
		"Endpoint /api/news/treatment called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	Get request to /api/news/all with no parameters

	Response:

	{
    "vaccine": {
        "data": [
            {
                "title": "Here's where we stand on getting a coronavirus vaccine - CNN",
                "description": "<ol><li><a href=\"https://www.cnn.com/2020/06/08/health/covid-19-vaccine-latest/index.html\" target=\"_blank\">Here's where we stand on getting a coronavirus vaccine</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">CNN</font></li><li><a href=\"https://www.thelancet.com/journals/lancet/article/PIIS0140-6736(20)31252-6/fulltext\" target=\"_blank\">COVID-19 vaccine development pipeline gears up</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">The Lancet</font></li><li><a href=\"https://www.healthline.com/health-news/why-companies-are-making-billions-of-covid-19-vaccine-doses-that-may-not-work\" target=\"_blank\">Why Companies Are Making Billions of COVID-19 Vaccine Doses</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">Healthline</font></li><li><a href=\"https://www.weforum.org/agenda/2020/06/astrazeneca-covid19-vaccine-gates-foundation\" target=\"_blank\">Pharmaceutical company pledges 2 billion COVID-19 vaccine doses</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">World Economic Forum</font></li><li><a href=\"https://www.usnews.com/news/health-news/articles/2020-06-08/experts-optimistic-in-search-for-covid-19-vaccine\" target=\"_blank\">Experts Optimistic in Search for COVID-19 Vaccine | Health News</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">U.S. News & World Report</font></li><li><strong><a href=\"https://news.google.com/stories/CAAqOQgKIjNDQklTSURvSmMzUnZjbmt0TXpZd1NoTUtFUWpyNnZMMmo0QU1FWTRkbWhyVmgyeWxLQUFQAQ?oc=5\" target=\"_blank\">View Full Coverage on Google News</a></strong></li></ol>",
                "url": "https://www.cnn.com/2020/06/08/health/covid-19-vaccine-latest/index.html",
                "urlToImage": "",
                "publishedAt": "Mon, 08 Jun 2020 12:12:50 GMT",
                "source": "CNN",
                "sourceURL": "https://www.cnn.com"
            },
            {
                "title": "COVID-19 vaccine trials bring hope for many but come too late for this family - NBC News",
                "description": "<a href=\"https://www.nbcnews.com/health/health-news/covid-19-vaccine-trials-bring-hope-many-come-too-late-n1226716\" target=\"_blank\">COVID-19 vaccine trials bring hope for many but come too late for this family</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">NBC News</font>",
                "url": "https://www.nbcnews.com/health/health-news/covid-19-vaccine-trials-bring-hope-many-come-too-late-n1226716",
                "urlToImage": "",
                "publishedAt": "Sun, 07 Jun 2020 16:02:58 GMT",
                "source": "NBC News",
                "sourceURL": "https://www.nbcnews.com"
            }
        ]
    },
    "treament": {
        "data": [
            {
                "title": "U.S. government’s supply of COVID-19 treatment drug, remdesivir, will run out at the end of the month - KTLA",
                "description": "<a href=\"https://ktla.com/news/coronavirus/u-s-governments-supply-of-covid-19-treatment-drug-remdesivir-will-run-out-at-the-end-of-the-month/\" target=\"_blank\">U.S. government’s supply of COVID-19 treatment drug, remdesivir, will run out at the end of the month</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">KTLA</font>",
                "url": "https://ktla.com/news/coronavirus/u-s-governments-supply-of-covid-19-treatment-drug-remdesivir-will-run-out-at-the-end-of-the-month/",
                "urlToImage": "",
                "publishedAt": "Mon, 08 Jun 2020 02:39:00 GMT",
                "source": "KTLA",
                "sourceURL": "https://ktla.com"
            },
            {
                "title": "Drug Trial Planned For Synthetic Cannabinoid COVID-19 Treatment - The Fresh Toast",
                "description": "<a href=\"https://thefreshtoast.com/cannabis/drug-trial-planned-for-synthetic-cannabinoid-covid-19-treatment/\" target=\"_blank\">Drug Trial Planned For Synthetic Cannabinoid COVID-19 Treatment</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">The Fresh Toast</font>",
                "url": "https://thefreshtoast.com/cannabis/drug-trial-planned-for-synthetic-cannabinoid-covid-19-treatment/",
                "urlToImage": "",
                "publishedAt": "Mon, 08 Jun 2020 13:36:15 GMT",
                "source": "The Fresh Toast",
                "sourceURL": "https://thefreshtoast.com"
            }
        ]
    },
    "news": {
        "data": [
            {
                "title": "What you need to know about the COVID-19 pandemic on 8 June - World Economic Forum",
                "description": "<a href=\"https://www.weforum.org/agenda/2020/06/covid-19-what-you-need-to-know-about-the-coronavirus-pandemic-on-8-june/\" target=\"_blank\">What you need to know about the COVID-19 pandemic on 8 June</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">World Economic Forum</font>",
                "url": "https://www.weforum.org/agenda/2020/06/covid-19-what-you-need-to-know-about-the-coronavirus-pandemic-on-8-june/",
                "urlToImage": "",
                "publishedAt": "Mon, 08 Jun 2020 08:51:16 GMT",
                "source": "World Economic Forum",
                "sourceURL": "https://www.weforum.org"
            },
            {
                "title": "How Reskilling Can Soften the Economic Blow of Covid-19 - Harvard Business Review",
                "description": "<a href=\"https://hbr.org/2020/06/how-reskilling-can-soften-the-economic-blow-of-covid-19\" target=\"_blank\">How Reskilling Can Soften the Economic Blow of Covid-19</a>&nbsp;&nbsp;<font color=\"#6f6f6f\">Harvard Business Review</font>",
                "url": "https://hbr.org/2020/06/how-reskilling-can-soften-the-economic-blow-of-covid-19",
                "urlToImage": "",
                "publishedAt": "Mon, 08 Jun 2020 12:10:46 GMT",
                "source": "Harvard Business Review",
                "sourceURL": "https://hbr.org"
			}
        ]
    }
}
*/
func newsAllHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := crnews.PerformAll()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "newsAllHandle",
		"Endpoint /api/news/all called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	Get request to /api/world with no parameters

	Response:

{
    "cases": [
        555,
        654,
        941,
        1434,

    ],
    "deaths": [
        17,
        18,
        26,
        42,
        56,
        82,
    ],
    "recovered": [
        28,
        30,
        36,
        39,
        52,
    ],
    "casesDaily": [
        99,
        287,
        493,
        684,
        809,
    ],
    "deathsDaily": [
        1,
        8,
        16,
        14,
        26,
        49,
    ],
    "recoveredDaily": [
        2,
        6,
        3,
        13,
        9,
        46,
        19,
        17,
    ]
}
}
*/
func worldHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := worldct.Perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "worldHandle",
		"Endpoint /api/world called with response JSON body "+string(jsonBody), status, elapsed)
}

/*
	Get request to /api/hotspot with no parameters

	Response:

{
    "mostCases": {
        "country": "Brazil",
        "data": [
            20599,
            26417,
            26928,
            33274,
            16409,
            11598,
            28936,
            28633,
            30925,
            30830,
            27075
        ]
    },
    "secondCases": {
        "country": "USA",
        "data": [
            18263,
            22577,
            24266,
            24146,
            20007,
            20848,
            20801,
            19699,
            21140,
            24720,
            22681
        ]
    },
    "thirdCases": {
        "country": "Russia",
        "data": [
            8338,
            8371,
            8572,
            8952,
            9268,
            8485,
            8858,
            8529,
            8823,
            8718,
            8846
        ]
    },
    "mostDeaths": {
        "country": "USA",
        "data": [
            1505,
            1199,
            1193,
            967,
            605,
            768,
            1031,
            995,
            1036,
            921,
            670
        ]
    },
    "secondDeaths": {
        "country": "Brazil",
        "data": [
            1086,
            1156,
            1124,
            956,
            480,
            623,
            1262,
            1349,
            1473,
            1005,
            904
        ]
    },
    "thirdDeaths": {
        "country": "Mexico",
        "data": [
            463,
            447,
            371,
            364,
            151,
            237,
            470,
            1092,
            816,
            625,
            341
        ]
    }
}
*/
func hotspotHandle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := hotspot.Perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "hotspotHandle",
		"Endpoint /api/hotspot called with response JSON body "+string(jsonBody), status, elapsed)
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
