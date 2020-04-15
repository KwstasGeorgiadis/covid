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

	country, err := curve.CompareDeathsCountries(compareRequest.NameOne, compareRequest.NameTwo)
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

	country, err := curve.CompareDeathsFromFirstDeathCountries(compareRequest.NameOne, compareRequest.NameTwo)
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

	country, err := curve.ComparePerDayDeathsCountries(compareRequest.NameOne, compareRequest.NameTwo)
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

//PerformPercentangePerDayDeath used in the /compare/firstdeath endpoint's handle to return
//	the Compare struct as a json response by calling
//	curve.ComparePerCentDeathsCountries which get
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
//  and an array that contains incremental percentage per day. It is sorted
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
func PerformPercentangePerDayDeath(r *http.Request) ([]byte, int) {
	var compareRequest Request

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		applogger.Log("ERROR", "compare", "PerformPercentangePerDayDeath", errIoutilReadAll.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	unmarshallError := json.Unmarshal(b, &compareRequest)
	if unmarshallError != nil {
		applogger.Log("ERROR", "compare", "PerformPercentangePerDayDeath", unmarshallError.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: unmarshallError.Error(), Code: 400})
		return statsErrJSONBody, 400
	}

	applogger.Log("INFO", "compare", "PerformPercentangePerDayDeath",
		fmt.Sprintf("Getting this request %v", compareRequest))

	country, err := curve.ComparePerCentDeathsCountries(compareRequest.NameOne, compareRequest.NameTwo)
	if err != nil {
		applogger.Log("ERROR", "compare", "PerformPercentangePerDayDeath", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(country)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "compare", "PerformPercentangePerDayDeath", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "compare", "PerformPercentangePerDayDeath",
		"Returning status: 200 with JSONbody "+string(jsonBody))
	return jsonBody, 200
}
