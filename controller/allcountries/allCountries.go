package allcountries

import (
	"encoding/json"

	applogger "github.com/junkd0g/covid/lib/applogger"
	stats "github.com/junkd0g/covid/lib/stats"
	structs "github.com/junkd0g/covid/lib/structs"
)

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
func Perform() ([]byte, int) {

	totalStats, err := stats.GetAllCountriesName()
	if err != nil {
		applogger.Log("ERROR", "allcountries", "Perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(totalStats)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "allcountries", "Perform", err.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "allcountries", "Perform",
		"Returning status: 200 with JSONbody "+string(jsonBody))

	return jsonBody, 200
}
