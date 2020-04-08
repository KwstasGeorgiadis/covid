package sortcon

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	stats "../../lib/stats"
	structs "../../lib/structs"
)

//SortRequest used for the https request's body
type SortRequest struct {
	Type string `json:"type"`
}

//Perform used in the /sort endpoint's handle to return
//	the structs.Countries struct as a json response by calling
//	stats.SortByDeaths() or tats.GetAllCountries() or stats.GetAllCountries()
//  or stats.SortByCasesPerOneMillion() or stats.SortByCritical() or
//  stats.SortByActive() or stats.SortByRecovered() or stats.SortByTodayDeaths()
//  or stats.SortByTodayCases() or stats.SortByCases()
//  which get and return sorted by field data: array
//
//	CompareRequest used as the struct for the request
//		example:
//			{
//				"type" : "deaths"
//			}
//
//	In this JSON format
//	{
//		"data": [{
//			"country": "Italy",
//			"cases": 124632,
//			"todayCases": 4805,
//			"deaths": 15362,
//			"todayDeaths": 681,
//			"recovered": 20996,
//			"active": 88274,
//			"critical": 3994,
//			"casesPerOneMillion": 2061
//		},
//		{
//			"country": "Spain",
//			"cases": 124736,
//			"todayCases": 5537,
//			"deaths": 11744,
//			"todayDeaths": 546,
//			"recovered": 34219,
//			"active": 78773,
//			"critical": 6416,
//			"casesPerOneMillion": 2668
//		}]
//	}
//
//
//	@param r *http.Request used to get http request's body
//
//	@return array of bytes of the json object
//	@return int http code status
func Perform(r *http.Request) ([]byte, int) {
	var sortRequest SortRequest

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)

	unmarshallError := json.Unmarshal(b, &sortRequest)
	if unmarshallError != nil {
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: unmarshallError.Error(), Code: 400})
		return statsErrJSONBody, 400
	}

	sortType := sortRequest.Type
	var countries structs.Countries
	var countriesError error

	switch sortType {
	case "deaths":
		countries, countriesError = stats.SortByDeaths()
		if countriesError != nil {
			statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: countriesError.Error(), Code: 500})
			return statsErrJSONBody, 500
		}
	case "cases":
		countries, countriesError = stats.SortByCases()
		if errIoutilReadAll != nil {
			statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: countriesError.Error(), Code: 500})
			return statsErrJSONBody, 500
		}
	case "todayCases":
		countries, countriesError = stats.SortByTodayCases()
		if countriesError != nil {
			statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: countriesError.Error(), Code: 500})
			return statsErrJSONBody, 500
		}
	case "todayDeaths":
		countries, countriesError = stats.SortByTodayDeaths()
		if countriesError != nil {
			statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: countriesError.Error(), Code: 500})
			return statsErrJSONBody, 500
		}
	case "recovered":
		countries, countriesError = stats.SortByRecovered()
		if errIoutilReadAll != nil {
			statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: countriesError.Error(), Code: 500})
			return statsErrJSONBody, 500
		}
	case "active":
		countries, countriesError = stats.SortByActive()
		if errIoutilReadAll != nil {
			statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: countriesError.Error(), Code: 500})
			return statsErrJSONBody, 500
		}
	case "critical":
		countries, countriesError = stats.SortByCritical()
		if countriesError != nil {
			statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: countriesError.Error(), Code: 500})
			return statsErrJSONBody, 500
		}
	case "casesPerOneMillion":
		countries, countriesError = stats.SortByCasesPerOneMillion()
		if countriesError != nil {
			statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: countriesError.Error(), Code: 500})
			return statsErrJSONBody, 500
		}
	default:
		countries, countriesError = stats.GetAllCountries()
		if countriesError != nil {
			statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: countriesError.Error(), Code: 500})
			return statsErrJSONBody, 500
		}
	}

	jsonBody, err := json.Marshal(countries)
	if err != nil {
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return errorJSONBody, 500
	}
	return jsonBody, 200
}
