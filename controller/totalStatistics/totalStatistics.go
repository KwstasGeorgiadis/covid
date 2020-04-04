package totalStatisticsCon

import (
	"encoding/json"

	stats "../../lib/stats"
	structs "../../lib/structs"
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

	totalStats := stats.GetTotalStats()
	jsonBody, err := json.Marshal(totalStats)
	if err != nil {
		errorJSONBody, _ := json.Marshal(structs.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return errorJSONBody, 500
	}

	return jsonBody, 200
}
