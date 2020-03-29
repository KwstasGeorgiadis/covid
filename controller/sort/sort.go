package sortCon

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	stats "../../lib/stats"
)

type SortRequest struct {
	Type string `json:"type"`
}

func Perform(r *http.Request) ([]byte, int) {
	var sortRequest SortRequest

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		// return some 500 stuff here
		fmt.Println(errIoutilReadAll.Error())
	}

	json.Unmarshal(b, &sortRequest)

	sortType := sortRequest.Type
	var countries stats.Countries

	switch sortType {
	case "deaths":
		countries = stats.SortByDeaths()
	case "cases":
		countries = stats.SortByCases()
	case "todayCases":
		countries = stats.SortByTodayCases()
	case "todayDeaths":
		countries = stats.SortByTodayDeaths()
	case "recovered":
		countries = stats.SortByRecovered()
	case "active":
		countries = stats.SortByActive()
	case "critical":
		countries = stats.SortByCritical()
	case "casesPerOneMillion":
		countries = stats.SortByCasesPerOneMillion()
	default:
		countries = stats.GetAllCountries()
	}
	jsonBody, _ := json.Marshal(countries)

	return jsonBody, 200
}
