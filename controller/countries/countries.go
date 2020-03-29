package countriesCon

import (
	"encoding/json"

	stats "../../lib/stats"
)

func Perform() ([]byte, int) {

	countries := stats.GetAllCountries()
	jsonBody, _ := json.Marshal(countries)

	return jsonBody, 200
}
