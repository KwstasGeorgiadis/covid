package countryCon

import (
	"encoding/json"

	stats "../../lib/stats"
	structs "../../lib/structs"

	"io/ioutil"
	"net/http"
)

type CountryRequest struct {
	Name string `json:"country"`
}

func Perform(r *http.Request) ([]byte, int) {
	var countryRequest CountryRequest

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: errIoutilReadAll.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	json.Unmarshal(b, &countryRequest)

	country,err := stats.GetCountry(countryRequest.Name)
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
