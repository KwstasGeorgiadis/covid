package allCountries

import (
	"encoding/json"

	stats "../../lib/stats"
)

func Perform() ([]byte, int) {

	totalStats := stats.GetAllCountriesName()
	jsonBody, _ := json.Marshal(totalStats)

	return jsonBody, 200
}
