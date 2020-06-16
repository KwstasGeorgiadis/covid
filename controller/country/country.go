package countrycon

import (
	"encoding/json"
	"time"

	applogger "github.com/junkd0g/covid/lib/applogger"
	stats "github.com/junkd0g/covid/lib/stats"
	structs "github.com/junkd0g/covid/lib/structs"

	"io/ioutil"
	"net/http"
)

//CountryRequest used for the https request's body
type CountryRequest struct {
	Name string `json:"country"`
}

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
func Country(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := Perform(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "main", "country",
		"Endpoint /api/country called with response JSON body "+string(jsonBody), status, elapsed)
}

//Perform used in the /country endpoint's handle to return
//	the Country struct as a json response by calling
//	stats.GetCountry which returns
//
//	CountryRequest used as the struct for the request
//		example:
//			{
//				"country" : "Greece",
//			}
//
//	Country string value of country name,
//	cases integer in total confirm cases of the country, todayCases int contains
//  today's cases in the country, todayDeaths int contains today's deaths
//  in the country, deaths integer of total confirm deaths in the country,
//	active integer of total confirm active cases in the country,
//	critical integer of total confirm in critical conditions cases in the country,
//  casesPerOneMillion float cases per one millions
//
//	In this JSON format
//	{
//		"country": "Greece",
//		"cases": 1061,
//		"todayCases": 0,
//		"deaths": 37,
//		"todayDeaths": 5,
//		"recovered": 52,
//		"active": 972,
//		"critical": 66,
//		"casesPerOneMillion": 102
//	}
//
//	@param r *http.Request used to get http request's body
//
//	@return array of bytes of the json object
//	@return int http code status
func Perform(r *http.Request) ([]byte, int) {
	var countryRequest CountryRequest

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		applogger.Log("ERROR", "countrycon", "Perform", errIoutilReadAll.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	json.Unmarshal(b, &countryRequest)

	country, err := stats.GetCountry(countryRequest.Name)
	if err != nil {
		applogger.Log("ERROR", "countrycon", "Perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(country)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "countrycon", "Perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}
