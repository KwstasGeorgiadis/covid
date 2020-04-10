package statisticscon

// THIS oNE NEEDS ATTENTION
import (
	"encoding/json"

	stats "../../lib/stats"
	structs "../../lib/structs"

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
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	unmarshallError := json.Unmarshal(b, &countryRequest)
	if unmarshallError != nil {
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: unmarshallError.Error(), Code: 400})
		return statsErrJSONBody, 400
	}

	country, err := stats.PercentancePerCountry(countryRequest.Name)
	if err != nil {
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(country)
	if jsonBodyErr != nil {
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}
