package statisticscon

// THIS oNE NEEDS ATTENTION
import (
	"encoding/json"
	"fmt"

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

//Perform used in the /stats endpoint's handle to return
//	the structs.Countries struct as a json response by calling
//	stats.SortByDeaths() or tats.GetAllCountries() or stats.GetAllCountries()
//  or stats.SortByCasesPerOneMillion() or stats.SortByCritical() or
//  which get and return sorted by field data: array
//
//	CompareRequest used as the struct for the request
//		example:
//			{
//				"country" : "deaths"
//			}
//
//	In this JSON format
//
//
//
//	@param r *http.Request used to get http request's body
//
//	@return array of bytes of the json object
//	@return int http code status
func Perform(r *http.Request) ([]byte, int) {
	var countryRequest CountryRequest

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		applogger.Log("ERROR", "statisticscon", "Perform", errIoutilReadAll.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	unmarshallError := json.Unmarshal(b, &countryRequest)
	if unmarshallError != nil {
		applogger.Log("ERROR", "statisticscon", "Perform", unmarshallError.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: unmarshallError.Error(), Code: 400})
		return statsErrJSONBody, 400
	}

	applogger.Log("INFO", "statisticscon", "Perform",
		fmt.Sprintf("Getting this request %v", countryRequest))

	country, err := stats.PercentancePerCountry(countryRequest.Name)
	if err != nil {
		applogger.Log("ERROR", "statisticscon", "Perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(country)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "statisticscon", "Perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "statisticscon", "Perform",
		"Returning status: 200 with JSONbody "+string(jsonBody))
	return jsonBody, 200
}
