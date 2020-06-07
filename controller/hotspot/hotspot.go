package hotspot

import (
	"encoding/json"

	analytics "github.com/junkd0g/covid/lib/analytics"
	applogger "github.com/junkd0g/covid/lib/applogger"
	structs "github.com/junkd0g/covid/lib/structs"
)

//Perform used in the /compare endpoint's handle to return
//	@return array of bytes of the json object
//	@return int http code status
func Perform() ([]byte, int) {

	worldData, err := analytics.MostCasesDeathsNearPast()
	if err != nil {
		applogger.Log("ERROR", "hotspot", "Perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(worldData)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "hotspot", "Perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "hotspot", "Perform",
		"Returning status: 200 with JSONbody "+string(jsonBody))
	return jsonBody, 200
}
