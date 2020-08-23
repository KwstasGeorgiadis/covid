package cssectl

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	applogger "github.com/junkd0g/covid/lib/applogger"
	csse "github.com/junkd0g/covid/lib/csse"
	merror "github.com/junkd0g/neji"
)

/*
	Get request to /api/csse/{country}

	{
    	"country": "US",
    	"data": [
        	{
            	"county": "<nil>",
            	"province": "Diamond Princess",
            	"cases": 49,
            	"deaths": 0,
            	"recovered": 0
        	},
        	{
            	"county": "<nil>",
            	"province": "Grand Princess",
            	"cases": 103,
            	"deaths": 3,
            	"recovered": 0
			}
		]
	}
*/
func Handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	jsonBody, status := perform(vars["country"])
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "cssectl", "Handle",
		"Endpoint /api/world called with response JSON body "+string(jsonBody), status, elapsed)
}

//Perform used in the /api/csse/{country} endpoint's handle to return
//	@return array of bytes of the json object
//	@return int http code status
func perform(country string) ([]byte, int) {

	csseData, err := csse.GetCSSECountryData(country)
	if err != nil {
		applogger.Log("ERROR", "cssectl", "perform", err.Error())
		statsErrJSONBody, _ := merror.SimpeErrorResponseWithStatus(500, err)
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(csseData)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "cssectl", "perform", jsonBodyErr.Error())
		errorJSONBody, _ := merror.SimpeErrorResponseWithStatus(500, jsonBodyErr)
		return errorJSONBody, 500
	}

	return jsonBody, 200
}
