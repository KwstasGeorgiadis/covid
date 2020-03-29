package totalStatisticsCon

import (
	"encoding/json"

	stats "../../lib/stats"
)

func Perform() ([]byte, int) {

	totalStats := stats.GetTotalStats()
	jsonBody, _ := json.Marshal(totalStats)

	return jsonBody, 200
}
