package compare

/*
	Controller used for the endpoints:
		/compare
		/compare/firstdeath
		/compare/perday
*/

import (
	"encoding/json"
	"fmt"

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

//Perform used in the /compare endpoint's handle to return
//	the Compare struct as a json response by calling
//	curve.CompareDeathsCountries which get and return grobal statistics
//
//	Request used as the struct for the request
//		example:
//			{
//				"countryOne" : "Italy",
//				"countryTwo" : "Spain"
//			}
//
//	Data structure that returns for two countries the names
//  and an array that contains deaths per day. It is sorted
//  and the first element is for the date 22/01/2020
//
//	In this JSON format
//  {
//    "countryOne": {
//        "country": "Spain",
//        "data": [
//            5,
//            10,
//		   	  17
//		   ]
//		},
//		"countryTwo": {
//      	"country": "Italy",
//       	"data": [
//            	197,
//            	233,
//				366
//			]
//		}
//	}
//
//	@param r *http.Request used to get http request's body
//
//	@return array of bytes of the json object
//	@return int http code status
func Perform(r *http.Request) ([]byte, int) {
	var compareRequest Request

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		applogger.Log("ERROR", "compare", "Perform", errIoutilReadAll.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	unmarshallError := json.Unmarshal(b, &compareRequest)
	if unmarshallError != nil {
		applogger.Log("ERROR", "compare", "Perform", unmarshallError.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: unmarshallError.Error(), Code: 400})
		return statsErrJSONBody, 400
	}

	applogger.Log("INFO", "compare", "Perform",
		fmt.Sprintf("Getting this request %v", compareRequest))

	country, err := curve.CompareDeathsCountries(compareRequest.NameOne, compareRequest.NameTwo, structs.CountryCurve{}, structs.CountryCurve{})
	if err != nil {
		applogger.Log("ERROR", "compare", "Perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(country)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "compare", "Perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "compare", "Perform",
		"Returning status: 200 with JSONbody "+string(jsonBody))
	return jsonBody, 200
}

//PerformFromFirstDeath used in the /compare/firstdeath endpoint's handle to return
//	the Compare struct as a json response by calling
//	curve.CompareDeathsFromFirstDeathCountries which get
//	and return grobal statistics
//
//	Request used as the struct for the request
//		example:
//			{
//				"countryOne" : "Italy",
//				"countryTwo" : "Spain"
//			}
//
//	Data structure that returns for two countries the names
//  and an array that contains total deaths per day. It is sorted
//  and the first element is for the date when the country
//  had their first death
//
//	In this JSON format
//  {
//    "countryOne": {
//        "country": "Spain",
//        "data": [
//            5,
//            10,
//		   	  17
//		   ]
//		},
//		"countryTwo": {
//      	"country": "Italy",
//       	"data": [
//            	197,
//            	233,
//				366
//			]
//		}
//	}
//
//	@param r *http.Request used to get http request's body
//
//	@return array of bytes of the json object
//	@return int http code status
func PerformFromFirstDeath(r *http.Request) ([]byte, int) {

	var compareRequest Request
	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		applogger.Log("ERROR", "compare", "PerformFromFirstDeath", errIoutilReadAll.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	unmarshallError := json.Unmarshal(b, &compareRequest)
	if unmarshallError != nil {
		applogger.Log("ERROR", "compare", "PerformFromFirstDeath", unmarshallError.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: unmarshallError.Error(), Code: 400})
		return statsErrJSONBody, 400
	}

	applogger.Log("INFO", "compare", "PerformFromFirstDeath",
		fmt.Sprintf("Getting this request %v", compareRequest))

	country, err := curve.CompareDeathsFromFirstDeathCountries(compareRequest.NameOne, compareRequest.NameTwo, structs.CountryCurve{}, structs.CountryCurve{})
	if err != nil {
		applogger.Log("ERROR", "compare", "PerformFromFirstDeath", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(country)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "compare", "PerformFromFirstDeath", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "compare", "PerformFromFirstDeath",
		"Returning status: 200 with JSONbody "+string(jsonBody))
	return jsonBody, 200
}

//PerformPerDayDeath used in the /compare/firstdeath endpoint's handle to return
//	the Compare struct as a json response by calling
//	curve.ComparePerDayDeathsCountries which get
//	and return grobal statistics
//
//	Request used as the struct for the request
//		example:
//			{
//				"countryOne" : "Italy",
//				"countryTwo" : "Spain"
//			}
//
//	Data structure that returns for two countries the names
//  and an array that contains deaths per day. It is sorted
//  and the first element is for the date when the country
//  had their first death
//
//	In this JSON format
//  {
//    "countryOne": {
//        "country": "Spain",
//        "data": [
//            5,
//            10,
//		   	  17
//		   ]
//		},
//		"countryTwo": {
//      	"country": "Italy",
//       	"data": [
//            	197,
//            	233,
//				366
//			]
//		}
//	}
//
//	@param r *http.Request used to get http request's body
//
//	@return array of bytes of the json object
//	@return int http code status
func PerformPerDayDeath(r *http.Request) ([]byte, int) {
	var compareRequest Request

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		applogger.Log("ERROR", "compare", "PerformFromFirstDeath", errIoutilReadAll.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	unmarshallError := json.Unmarshal(b, &compareRequest)
	if unmarshallError != nil {
		applogger.Log("ERROR", "compare", "PerformPerDayDeath", unmarshallError.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: unmarshallError.Error(), Code: 400})
		return statsErrJSONBody, 400
	}

	applogger.Log("INFO", "compare", "PerformPerDayDeath",
		fmt.Sprintf("Getting this request %v", compareRequest))

	country, err := curve.ComparePerDayDeathsCountries(compareRequest.NameOne, compareRequest.NameTwo, structs.CountryCurve{}, structs.CountryCurve{})
	if err != nil {
		applogger.Log("ERROR", "compare", "PerformPerDayDeath", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(country)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "compare", "PerformPerDayDeath", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "compare", "PerformPerDayDeath",
		"Returning status: 200 with JSONbody "+string(jsonBody))
	return jsonBody, 200
}

//PerformCompareRecorey used in the /compare/recovery endpoint's handle to return
//	the Compare struct as a json response by calling
//	curve.CompareRecoveryCountries which get and return grobal statistics
//
//	Request used as the struct for the request
//		example:
//			{
//				"countryOne" : "Italy",
//				"countryTwo" : "Spain"
//			}
//
//	Data structure that returns for two countries the names
//  and an array that contains total recovery patients per day. It is sorted
//  and the first element is for the date 22/01/2020
//
//	In this JSON format
//  {
//    "countryOne": {
//        "country": "Spain",
//        "data": [
//            5,
//            10,
//		   	  17
//		   ]
//		},
//		"countryTwo": {
//      	"country": "Italy",
//       	"data": [
//            	197,
//            	233,
//				366
//			]
//		}
//	}
//
//	@param r *http.Request used to get http request's body
//
//	@return array of bytes of the json object
//	@return int http code status
func PerformCompareRecorey(r *http.Request) ([]byte, int) {
	var compareRequest Request

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		applogger.Log("ERROR", "compare", "Perform", errIoutilReadAll.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	unmarshallError := json.Unmarshal(b, &compareRequest)
	if unmarshallError != nil {
		applogger.Log("ERROR", "compare", "Perform", unmarshallError.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: unmarshallError.Error(), Code: 400})
		return statsErrJSONBody, 400
	}

	applogger.Log("INFO", "compare", "Perform",
		fmt.Sprintf("Getting this request %v", compareRequest))

	country, err := curve.CompareRecoveryCountries(compareRequest.NameOne, compareRequest.NameTwo, structs.CountryCurve{}, structs.CountryCurve{})
	if err != nil {
		applogger.Log("ERROR", "compare", "Perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(country)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "compare", "Perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "compare", "Perform",
		"Returning status: 200 with JSONbody "+string(jsonBody))
	return jsonBody, 200
}

//PerformCompareCases used in the /compare/cases endpoint's handle to return
//	the Compare struct as a json response by calling
//	curve.CompareCasesCountries which get and return grobal statistics
//
//	Request used as the struct for the request
//		example:
//			{
//				"countryOne" : "Italy",
//				"countryTwo" : "Spain"
//			}
//
//	Data structure that returns for two countries the names
//  and an array that contains total cases per day. It is sorted
//  and the first element is for the date 22/01/2020
//
//	In this JSON format
//  {
//    "countryOne": {
//        "country": "Spain",
//        "data": [
//            5,
//            10,
//		   	  17
//		   ]
//		},
//		"countryTwo": {
//      	"country": "Italy",
//       	"data": [
//            	197,
//            	233,
//				366
//			]
//		}
//	}
//
//	@param r *http.Request used to get http request's body
//
//	@return array of bytes of the json object
//	@return int http code status
func PerformCompareCases(r *http.Request) ([]byte, int) {
	var compareRequest Request

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		applogger.Log("ERROR", "compare", "Perform", errIoutilReadAll.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	unmarshallError := json.Unmarshal(b, &compareRequest)
	if unmarshallError != nil {
		applogger.Log("ERROR", "compare", "Perform", unmarshallError.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: unmarshallError.Error(), Code: 400})
		return statsErrJSONBody, 400
	}

	country, err := curve.CompareCasesCountries(compareRequest.NameOne, compareRequest.NameTwo, structs.CountryCurve{}, structs.CountryCurve{})
	if err != nil {
		applogger.Log("ERROR", "compare", "Perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(country)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "compare", "Perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}

//PerformCompareUniquePerDayCases used in the /compare/cases/unique endpoint's handle to return
//	the Compare struct as a json response by calling
//	curve.CompareCasesCountries which get and return grobal statistics
//
//	Request used as the struct for the request
//		example:
//			{
//				"countryOne" : "Italy",
//				"countryTwo" : "Spain"
//			}
//
//	Data structure that returns for two countries the names
//  and an array that contains unique cases per day. It is sorted
//  and the first element is for the date 22/01/2020
//
//	In this JSON format
//  {
//    "countryOne": {
//        "country": "Spain",
//        "data": [
//            5,
//            10,
//		   	  17
//		   ]
//		},
//		"countryTwo": {
//      	"country": "Italy",
//       	"data": [
//            	197,
//            	233,
//				366
//			]
//		}
//	}
//
//	@param r *http.Request used to get http request's body
//
//	@return array of bytes of the json object
//	@return int http code status
func PerformCompareUniquePerDayCases(r *http.Request) ([]byte, int) {
	var compareRequest Request

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		applogger.Log("ERROR", "compare", "Perform", errIoutilReadAll.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	unmarshallError := json.Unmarshal(b, &compareRequest)
	if unmarshallError != nil {
		applogger.Log("ERROR", "compare", "Perform", unmarshallError.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: unmarshallError.Error(), Code: 400})
		return statsErrJSONBody, 400
	}

	applogger.Log("INFO", "compare", "Perform",
		fmt.Sprintf("Getting this request %v", compareRequest))

	country, err := curve.ComparePerDayCasesCountries(compareRequest.NameOne, compareRequest.NameTwo, structs.CountryCurve{}, structs.CountryCurve{})
	if err != nil {
		applogger.Log("ERROR", "compare", "Perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(country)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "compare", "Perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "compare", "Perform",
		"Returning status: 200 with JSONbody "+string(jsonBody))
	return jsonBody, 200
}

func PerformAll(r *http.Request) ([]byte, int) {
	var compareRequest Request

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		applogger.Log("ERROR", "compare", "Perform", errIoutilReadAll.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	unmarshallError := json.Unmarshal(b, &compareRequest)
	if unmarshallError != nil {
		applogger.Log("ERROR", "compare", "Perform", unmarshallError.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: unmarshallError.Error(), Code: 400})
		return statsErrJSONBody, 400
	}

	applogger.Log("INFO", "compare", "Perform",
		fmt.Sprintf("Getting this request %v", compareRequest))

	getCountrySt, getCountryStErr := curve.GetCountry(compareRequest.NameOne)
	if getCountryStErr != nil {
		applogger.Log("ERROR", "compare", "PerformAll", getCountryStErr.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: getCountryStErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}
	getCountryStTwo, getCountryStTwoErr := curve.GetCountry(compareRequest.NameTwo)
	if getCountryStTwoErr != nil {
		applogger.Log("ERROR", "compare", "PerformAll", getCountryStTwoErr.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: getCountryStTwoErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	compareDeathsCountries, compareDeathsCountriesErr := curve.CompareDeathsCountries(compareRequest.NameOne, compareRequest.NameTwo, getCountrySt, getCountryStTwo)
	if compareDeathsCountriesErr != nil {
		applogger.Log("ERROR", "compare", "PerformAll", compareDeathsCountriesErr.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: compareDeathsCountriesErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	compareRecoveryCountries, compareRecoveryCountriesErr := curve.CompareRecoveryCountries(compareRequest.NameOne, compareRequest.NameTwo, getCountrySt, getCountryStTwo)
	if compareDeathsCountriesErr != nil {
		applogger.Log("ERROR", "compare", "PerformAll", compareRecoveryCountriesErr.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: compareRecoveryCountriesErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	compareCasesCountries, compareCasesCountriesErr := curve.CompareCasesCountries(compareRequest.NameOne, compareRequest.NameTwo, getCountrySt, getCountryStTwo)
	if compareDeathsCountriesErr != nil {
		applogger.Log("ERROR", "compare", "PerformAll", compareCasesCountriesErr.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: compareCasesCountriesErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	comparePerDayCasesCountries, comparePerDayCasesCountriesErr := curve.ComparePerDayCasesCountries(compareRequest.NameOne, compareRequest.NameTwo, getCountrySt, getCountryStTwo)
	if compareDeathsCountriesErr != nil {
		applogger.Log("ERROR", "compare", "PerformAll", comparePerDayCasesCountriesErr.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: comparePerDayCasesCountriesErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	comparePerDayDeathsCountries, comparePerDayDeathsCountriesErr := curve.ComparePerDayDeathsCountries(compareRequest.NameOne, compareRequest.NameTwo, getCountrySt, getCountryStTwo)
	if compareDeathsCountriesErr != nil {
		applogger.Log("ERROR", "compare", "PerformAll", comparePerDayDeathsCountriesErr.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: comparePerDayDeathsCountriesErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	compareDeathsFromFirstDeathCountries, compareDeathsFromFirstDeathCountriesErr := curve.CompareDeathsFromFirstDeathCountries(compareRequest.NameOne, compareRequest.NameTwo, getCountrySt, getCountryStTwo)
	if compareDeathsCountriesErr != nil {
		applogger.Log("ERROR", "compare", "PerformAll", compareDeathsFromFirstDeathCountriesErr.Error())
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

	countryTwoAllData.Country = compareRequest.NameOne
	countryTwoAllData.DataDeaths = compareDeathsCountries.CountryTwo.Data
	countryTwoAllData.DataDeathsFromFirst = compareDeathsFromFirstDeathCountries.CountryTwo.Data
	countryTwoAllData.DataDeathsPerDay = comparePerDayDeathsCountries.CountryTwo.Data
	countryTwoAllData.DataRecovered = compareRecoveryCountries.CountryTwo.Data
	countryTwoAllData.DataCases = compareCasesCountries.CountryTwo.Data
	countryTwoAllData.DataCasesFromFist = comparePerDayCasesCountries.CountryTwo.Data

	jsonBody, jsonBodyErr := json.Marshal(structs.CompareAll{countryOneAllData, countryTwoAllData})
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "compare", "Perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "compare", "Perform",
		"Returning status: 200 with JSONbody "+string(jsonBody))
	return jsonBody, 200
}
