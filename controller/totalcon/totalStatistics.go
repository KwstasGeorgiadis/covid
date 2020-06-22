package totalcon

import (
	"encoding/json"
	"net/http"
	"time"

	applogger "github.com/junkd0g/covid/lib/applogger"
	merror "github.com/junkd0g/covid/lib/model/error"
	stats "github.com/junkd0g/covid/lib/stats"
)

/*
	Get request to /api/total with no parameters

	Response:

	{
    	"todayPerCentOfTotalCases": 7,
    	"todayPerCentOfTotalDeaths": 6,
    	"totalCases": 1188489,
    	"totalDeaths": 64103,
    	"todayTotalCases": 71846,
    	"todayTotalDeaths": 4933
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
	applogger.LogHTTP("INFO", "totalcon", "Handle",
		"Endpoint /api/total called with response JSON body "+string(jsonBody), status, elapsed)
}

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
func perform() ([]byte, int) {

	totalStats, statsErr := stats.GetTotalStats()
	if statsErr != nil {
		applogger.Log("ERROR", "totalcon", "perform", statsErr.Error())
		statsErrJSONBody, _ := json.Marshal(merror.ErrorMessage{ErrorMessage: statsErr.Error(), Code: 500})
		return statsErrJSONBody, 500
	}

	jsonBody, err := json.Marshal(totalStats)
	if err != nil {
		applogger.Log("ERROR", "totalcon", "perform", err.Error())
		errorJSONBody, _ := json.Marshal(merror.ErrorMessage{ErrorMessage: err.Error(), Code: 500})
		return errorJSONBody, 500
	}
	return jsonBody, 200
}
