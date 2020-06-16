package worldct

import (
	"encoding/json"
	"net/http"
	"time"

	applogger "github.com/junkd0g/covid/lib/applogger"
	cworld "github.com/junkd0g/covid/lib/cworld"
	structs "github.com/junkd0g/covid/lib/structs"
)

/*
	Get request to /api/world with no parameters

	Response:

{
    "cases": [
        555,
        654,
        941,
        1434,

    ],
    "deaths": [
        17,
        18,
        26,
        42,
        56,
        82,
    ],
    "recovered": [
        28,
        30,
        36,
        39,
        52,
    ],
    "casesDaily": [
        99,
        287,
        493,
        684,
        809,
    ],
    "deathsDaily": [
        1,
        8,
        16,
        14,
        26,
        49,
    ],
    "recoveredDaily": [
        2,
        6,
        3,
        13,
        9,
        46,
        19,
        17,
    ]
}
}
*/
func Handle(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	jsonBody, status := perform()
	w.WriteHeader(status)
	w.Write(jsonBody)
	elapsed := time.Since(start).Seconds()
	applogger.LogHTTP("INFO", "worldct", "Handle",
		"Endpoint /api/world called with response JSON body "+string(jsonBody), status, elapsed)
}

//Perform used in the /compare endpoint's handle to return
//	@return array of bytes of the json object
//	@return int http code status
func perform() ([]byte, int) {

	worldData, err := cworld.GetaWorldHistory()
	if err != nil {
		applogger.Log("ERROR", "worldct", "perform", err.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, jsonBodyErr := json.Marshal(worldData)
	if jsonBodyErr != nil {
		applogger.Log("ERROR", "worldct", "perform", jsonBodyErr.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: jsonBodyErr.Error(), Code: 500})
		return errorJSONBody, 500
	}
	return jsonBody, 200
}
