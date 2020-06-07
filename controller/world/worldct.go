package worldct

import (
	"encoding/json"

	applogger "github.com/junkd0g/covid/lib/applogger"
	cworld "github.com/junkd0g/covid/lib/cworld"
	structs "github.com/junkd0g/covid/lib/structs"
)

//Perform used in the /compare endpoint's handle to return
//	@return array of bytes of the json object
//	@return int http code status
func Perform() ([]byte, int) {

	worldData, err := cworld.GetaWorldHistory()
	if err != nil {
		applogger.Log("ERROR", "worldct", "Perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(worldData)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "worldct", "Perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "worldct", "Perform",
		"Returning status: 200 with JSONbody "+string(jsonBody))
	return jsonBody, 200
}
