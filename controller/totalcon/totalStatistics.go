package totalcon

import (
	"encoding/json"

	applogger "github.com/junkd0g/covid/lib/applogger"
	stats "github.com/junkd0g/covid/lib/stats"
	structs "github.com/junkd0g/covid/lib/structs"
)

//Perform used in the /total endpoint's handle to return
//	the TotalStats struct as a json response by calling
//	stats.GetTotalStats which get and return grobal statistics
//
//	Total grobal covid-19 cases, total grobal deaths, today's grobal
//	cases, today's grobal deaths, today's percentance increase in todal grobal deaths
//	and  today's percentance increase in todal grobal cases.
//
//	In this JSON format
//  {
//		"todayPerCentOfTotalCases": 7,
//		"todayPerCentOfTotalDeaths": 6,
//		"totalCases": 1188489,
//		"totalDeaths": 64103,
//		"todayTotalCases": 71846,
//		"todayTotalDeaths": 4933
//	}
//
//	@return array of bytes of the json object
//	@return int http code status
func Perform() ([]byte, int) {

	totalStats, statsErr := stats.GetTotalStats()
	if statsErr != nil {
		applogger.Log("ERROR", "totalcon", "Perform", statsErr.Error())
		statsErrJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: statsErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, err := json.Marshal(totalStats)
	if err != nil {
		applogger.Log("ERROR", "totalcon", "Perform", err.Error())
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return errorJSONBody, 500
	}

	applogger.Log("INFO", "totalcon", "Perform",
		"Returning status: 200 with JSONbody "+string(jsonBody))
	return jsonBody, 200
}
