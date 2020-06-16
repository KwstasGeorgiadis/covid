package continentctl

import (
	"encoding/json"
	"net/http"
	"time"

	applogger "github.com/junkd0g/covid/lib/applogger"
	continent "github.com/junkd0g/covid/lib/continent"
	structs "github.com/junkd0g/covid/lib/structs"
)

/*
	Get request to /api/continent with no parameters

	Response:

{[
    {
        "updated": 1591969378141,
        "cases": 1313866,
        "todayCases": 884,
        "deaths": 56394,
        "todayDeaths": 21,
        "recovered": 678380,
        "todayRecovered": 111,
        "active": 579092,
        "critical": 11988,
        "casesPerOneMillion": 3051.48,
        "deathsPerOneMillion": 130.98,
        "tests": 5424046,
        "testsPerOneMillion": 12597.45,
        "population": 430566996,
        "continent": "South America",
        "activePerOneMillion": 1344.95,
        "recoveredPerOneMillion": 1575.55,
        "criticalPerOneMillion": 27.84,
        "countries": [
            "Argentina",
            "Bolivia",
            "Brazil",
            "Chile",
            "Colombia",
            "Ecuador",
            "Falkland Islands (Malvinas)",
            "French Guiana",
            "Guyana",
            "Paraguay",
            "Peru",
            "Suriname",
            "Uruguay",
            "Venezuela"
        ]
    },
    {
        "updated": 1591969378149,
        "cases": 8901,
        "todayCases": 5,
        "deaths": 124,
        "todayDeaths": 0,
        "recovered": 8371,
        "todayRecovered": 22,
        "active": 406,
        "critical": 2,
        "casesPerOneMillion": 217.71,
        "deathsPerOneMillion": 3.03,
        "tests": 2070918,
        "testsPerOneMillion": 50652.24,
        "population": 40885025,
        "continent": "Australia/Oceania",
        "activePerOneMillion": 9.93,
        "recoveredPerOneMillion": 204.74,
        "criticalPerOneMillion": 0.05,
        "countries": [
            "Australia",
            "Fiji",
            "French Polynesia",
            "New Caledonia",
            "New Zealand",
            "Papua New Guinea"
        ]
    }
]
*/
func Handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "continentct", "Handle",
		"Endpoint /api/world called with response JSON body "+string(jsonBody), status, elapsed)
}

//Perform used in the /api/continent endpoint's handle to return
//	@return array of bytes of the json object
//	@return int http code status
func perform() ([]byte, int) {

	continentData, err := continent.GetContinentData()
	if err != nil {
		applogger.Log("ERROR", "continentct", "perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(continentData)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "continentct", "perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}
