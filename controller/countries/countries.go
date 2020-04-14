package countriescon

import (
	"encoding/json"

	applogger "../../lib/applogger"
	stats "../../lib/stats"
	structs "../../lib/structs"
)

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
func Perform() ([]byte, int) {

	countries, err := stats.GetAllCountries()
	if err != nil {
		applogger.Log("ERROR", "comcountriesconpare", "Perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(countries)
	if err != nil {
		applogger.Log("ERROR", "comcountriesconpare", "Perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "compare", "Perform",
		"Returning status: 200 with JSONbody "+string(jsonBody))
	return jsonBody, 200
}
