package continentct

import (
	"encoding/json"

	applogger "github.com/junkd0g/covid/lib/applogger"
	continent "github.com/junkd0g/covid/lib/continent"
	structs "github.com/junkd0g/covid/lib/structs"
)

//Perform used in the /api/continent endpoint's handle to return
//	@return array of bytes of the json object
//	@return int http code status
func Perform() ([]byte, int) {

	continentData, err := continent.GetContinentData()
	if err != nil {
		applogger.Log("ERROR", "continentct", "Perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(continentData)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "continentct", "Perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "continentct", "Perform",
		"Returning status: 200 with JSONbody "+string(jsonBody))
	return jsonBody, 200
}
