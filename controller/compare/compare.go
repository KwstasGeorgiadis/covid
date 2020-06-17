package comparectl

/*
	Controller used for the endpoints:
		/compare
		/compare/firstdeath
		/compare/perday
*/

import (
	"encoding/json"
	"fmt"
	"time"

	applogger "github.com/junkd0g/covid/lib/applogger"
	curve "github.com/junkd0g/covid/lib/curve"
	structs "github.com/junkd0g/covid/lib/structs"

	"io/ioutil"
	"net/http"
)

//Request used for the https request's body
type Request struct {
	NameOne string `json:"countryOne"`
	NameTwo string `json:"countryTwo"`
}

// Handle POST request to /api/compare/all endpoint
/*
	Request:

	{
		"countryOne" : "Spain",
		"countryTwo" : "Italy"
	}

	Response

	{
    "countryOne": {
        "country": "Spain",
        "dataDeaths": [
            28,
            27888,
            27940,
            28628,
            28678,
            28752
        ],
        "dataDeathsFromFirst": [
            28,
            27888,
            27940,
            28628,
            28678,
            28752
        ],
        "dataDeathsPerDay": [
            110,
            52,
            688,
            50,
            74
        ],
        "dataRecoverd": [
            150376,
            150376,
            150376,
            150376,
            150376
        ],
        "dataCases": [
            15,
            32,
            45,
            84,
            120,
            165,
            222,
            259,
        ],
        "dataCasesFromFirst": [
            159,
            294,
            394,
            334,
            318,
            332,
            240
        ]
    },
    "countryTwo": {
        "country": "Spain",
        "dataDeaths": [
            0,
            0,
            0,
            0,
            0,
        ],
        "dataDeathsFromFirst": [
            33530,
            33601,
            33689,
            33774,
            33846,
            33899
        ],
        "dataDeathsPerDay": [
            1,
            1,
            1,
            4,
            3,
            2,
        ],
        "dataRecoverd": [
            0,
            1,
            1,
            1,
            2,
            3,
            45,
            46,
            46,
        ],
        "dataCases": [
            232664,
            232997,
            233197,
            233515,
            233836,
            234013,
            234531,
            234801,
            234998
        ],
        "dataCasesFromFirst": [
            200,
            318,
            321,
            177,
            518,
            270,
            197
        ]
    }
}
}
*/
func Handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := perform(r)
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "compare", "Handle",
		"Endpoint /compare/percent called with response JSON body "+string(jsonBody), status, elapsed)
}

func perform(r *http.Request) ([]byte, int) {
	var compareRequest Request

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		applogger.Log("ERROR", "compare", "perform", errIoutilReadAll.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	unmarshallError := json.Unmarshal(b, &compareRequest)
	if unmarshallError != nil {
		applogger.Log("ERROR", "compare", "perform", unmarshallError.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: unmarshallError.Error(), Code: 400})
		return statsErrJSONBody, 400
	}

	applogger.Log("INFO", "compare", "Perform",
		fmt.Sprintf("Getting this request %v", compareRequest))

	compareDeathsCountries, compareDeathsCountriesErr := curve.CompareDeathsCountries(compareRequest.NameOne, compareRequest.NameTwo)
	if compareDeathsCountriesErr != nil {
		applogger.Log("ERROR", "compare", "perform", compareDeathsCountriesErr.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: compareDeathsCountriesErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	compareRecoveryCountries, compareRecoveryCountriesErr := curve.CompareRecoveryCountries(compareRequest.NameOne, compareRequest.NameTwo)
	if compareDeathsCountriesErr != nil {
		applogger.Log("ERROR", "compare", "perform", compareRecoveryCountriesErr.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: compareRecoveryCountriesErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	compareCasesCountries, compareCasesCountriesErr := curve.CompareCasesCountries(compareRequest.NameOne, compareRequest.NameTwo)
	if compareDeathsCountriesErr != nil {
		applogger.Log("ERROR", "compare", "perform", compareCasesCountriesErr.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: compareCasesCountriesErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	comparePerDayCasesCountries, comparePerDayCasesCountriesErr := curve.ComparePerDayCasesCountries(compareRequest.NameOne, compareRequest.NameTwo)
	if compareDeathsCountriesErr != nil {
		applogger.Log("ERROR", "compare", "perform", comparePerDayCasesCountriesErr.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: comparePerDayCasesCountriesErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	comparePerDayDeathsCountries, comparePerDayDeathsCountriesErr := curve.ComparePerDayDeathsCountries(compareRequest.NameOne, compareRequest.NameTwo)
	if compareDeathsCountriesErr != nil {
		applogger.Log("ERROR", "compare", "perform", comparePerDayDeathsCountriesErr.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: comparePerDayDeathsCountriesErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	compareDeathsFromFirstDeathCountries, compareDeathsFromFirstDeathCountriesErr := curve.CompareDeathsFromFirstDeathCountries(compareRequest.NameOne, compareRequest.NameTwo)
	if compareDeathsCountriesErr != nil {
		applogger.Log("ERROR", "compare", "perform", compareDeathsFromFirstDeathCountriesErr.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: compareDeathsFromFirstDeathCountriesErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}
	var countryOneAllData structs.CompareAllData
	var countryTwoAllData structs.CompareAllData

	countryOneAllData.Country = compareRequest.NameOne
	countryOneAllData.DataDeaths = compareDeathsCountries.CountryOne.Data
	countryOneAllData.DataDeathsFromFirst = compareDeathsFromFirstDeathCountries.CountryOne.Data
	countryOneAllData.DataDeathsPerDay = comparePerDayDeathsCountries.CountryOne.Data
	countryOneAllData.DataRecovered = compareRecoveryCountries.CountryOne.Data
	countryOneAllData.DataCases = compareCasesCountries.CountryOne.Data
	countryOneAllData.DataCasesFromFist = comparePerDayCasesCountries.CountryOne.Data

	countryTwoAllData.Country = compareRequest.NameTwo
	countryTwoAllData.DataDeaths = compareDeathsCountries.CountryTwo.Data
	countryTwoAllData.DataDeathsFromFirst = compareDeathsFromFirstDeathCountries.CountryTwo.Data
	countryTwoAllData.DataDeathsPerDay = comparePerDayDeathsCountries.CountryTwo.Data
	countryTwoAllData.DataRecovered = compareRecoveryCountries.CountryTwo.Data
	countryTwoAllData.DataCases = compareCasesCountries.CountryTwo.Data
	countryTwoAllData.DataCasesFromFist = comparePerDayCasesCountries.CountryTwo.Data

	jsonBody, jsonBodyErr := json.Marshal(structs.CompareAll{countryOneAllData, countryTwoAllData})
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "compare", "perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "compare", "perform",
		"Returning status: 200 with JSONbody "+string(jsonBody))
	return jsonBody, 200
}
