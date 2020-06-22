package hotspot

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	analytics "github.com/junkd0g/covid/lib/analytics"
	applogger "github.com/junkd0g/covid/lib/applogger"
	merror "github.com/junkd0g/covid/lib/model/error"
)

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
func Handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	jsonBody, status := perform(vars["days"])
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "hotspot", "Handle",
		"Endpoint /api/hotspot called with response JSON body "+string(jsonBody), status, elapsed)
}

//Perform used in the /api/hotspot endpoint's handle to return
//	@return array of bytes of the json object
//	@return int http code status
func perform(days string) ([]byte, int) {
	i, errAtoi := strconv.Atoi(days)
	if errAtoi != nil {
		applogger.Log("ERROR", "hotspot", "perform", errAtoi.Error())
		statsErrJSONBody, _ := json.Marshal(merror.ErrorMessage{ErrorMessage: errAtoi.Error(), Code: 400})
		return statsErrJSONBody, 400
	}

	worldData, err := analytics.MostCasesDeathsNearPast(i)
	if err != nil {
		applogger.Log("ERROR", "hotspot", "perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(merror.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(worldData)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "hotspot", "perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(merror.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}
	return jsonBody, 200
}
