package sortCon

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	stats "../../lib/stats"
	structs "../../lib/structs"
)

type SortRequest struct {
	Type string `json:"type"`
}

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
