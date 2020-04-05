package allCountries

import (
	"encoding/json"

	stats "../../lib/stats"
	structs "../../lib/structs"
)

func Perform() ([]byte, int) {

	totalStats,err := stats.GetAllCountriesName()
	if err != nil {
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500 
	}

	jsonBody, jsonBodyErr := json.Marshal(totalStats)
	if jsonBodyErr != nil {
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}
