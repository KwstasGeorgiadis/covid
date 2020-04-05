package allcountries

import (
	"encoding/json"

	stats "../../lib/stats"
	structs "../../lib/structs"
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
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(totalStats)
	if jsonBodyErr != nil {
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}
