package countriesCon

import (
	"encoding/json"

	stats "../../lib/stats"
	structs "../../lib/structs"

)

func Perform() ([]byte, int) {

	countries, err := stats.GetAllCountries()
	if err != nil {
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, err := json.Marshal(countries)
	if err != nil {
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return errorJSONBody, 500
	}
	return jsonBody, 200
}
