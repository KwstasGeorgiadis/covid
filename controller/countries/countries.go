package countriescon

import (
	"encoding/json"
	"net/http"
	"time"

	applogger "github.com/junkd0g/covid/lib/applogger"
	merror "github.com/junkd0g/covid/lib/model/error"
	stats "github.com/junkd0g/covid/lib/stats"
)

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
func Handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "countriescon", "Handle",
		"Endpoint /api/countries called with response JSON body "+string(jsonBody), status, elapsed)
}

//Perform used in the /countries endpoint's handle to return
//	the structs.Countries struct as a json response by calling
//	stats.GetAllCountries() which returns grobal statistics
//
//	Array of all countries' object data. Country string value of country name,
//	cases integer in total confirm cases of the country, todayCases int contains
//  today's cases in the country, todayDeaths int contains today's deaths
//  in the country, deaths integer of total confirm deaths in the country,
//	active integer of total confirm active cases in the country,
//	critical integer of total confirm in critical conditions cases in the country,
//  casesPerOneMillion float cases per one millions
//
//	In this JSON format
//  {
//		"data": [
//			{
//				"country": "Zimbabwe",
//				"cases": 7,
//				"todayCases": 0,
//				"deaths": 1,
//				"todayDeaths": 0,
//				"recovered": 0,
//				"active": 6,
//				"critical": 0,
//				"casesPerOneMillion": 0.5
//			}
//		]
//	}
//
//	@return array of bytes of the json object
//	@return int http code status
func perform() ([]byte, int) {

	countries, err := stats.GetAllCountries()
	if err != nil {
		applogger.Log("ERROR", "countriescon", "perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(merror.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(countries)
	if err != nil {
		applogger.Log("ERROR", "countriescon", "perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(merror.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}
