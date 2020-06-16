package allcountries

import (
	"encoding/json"
	"net/http"
	"time"

	applogger "github.com/junkd0g/covid/lib/applogger"
	stats "github.com/junkd0g/covid/lib/stats"
	structs "github.com/junkd0g/covid/lib/structs"
)

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
func Handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "allcountries", "AllCountriesHandle",
		"Endpoint /api/countries/all called with response JSON body "+string(jsonBody), status, elapsed)
}

//Perform used in the /countries/all endpoint's handle to return
//	the AllCountriesName struct as a json response by calling
//	stats.GetAllCountriesName which get and return grobal statistics
//
//	Array of all countries' name that we have covid-19 stats
//
//	In this JSON format
//  {
//  	"countries": [
//   		"Zimbabwe",
//        	"Zambia",
//       	"Western Sahara"
//		]
//	}
//
//	@return array of bytes of the json object
//	@return int http code status
func perform() ([]byte, int) {

	totalStats, err := stats.GetAllCountriesName()
	if err != nil {
		applogger.Log("ERROR", "allcountries", "perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(totalStats)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "allcountries", "perform", err.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}
