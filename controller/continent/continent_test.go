package continentctl

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ContinentExpectedResponses []struct {
	Updated                int64    `json:"updated"`
	Cases                  int      `json:"cases"`
	TodayCases             int      `json:"todayCases"`
	Deaths                 int      `json:"deaths"`
	TodayDeaths            int      `json:"todayDeaths"`
	Recovered              int      `json:"recovered"`
	TodayRecovered         int      `json:"todayRecovered"`
	Active                 int      `json:"active"`
	Critical               int      `json:"critical"`
	CasesPerOneMillion     float64  `json:"casesPerOneMillion"`
	DeathsPerOneMillion    float64  `json:"deathsPerOneMillion"`
	Tests                  int      `json:"tests"`
	TestsPerOneMillion     float64  `json:"testsPerOneMillion"`
	Population             int      `json:"population"`
	Continent              string   `json:"continent"`
	ActivePerOneMillion    float64  `json:"activePerOneMillion"`
	RecoveredPerOneMillion float64  `json:"recoveredPerOneMillion"`
	CriticalPerOneMillion  int      `json:"criticalPerOneMillion"`
	Countries              []string `json:"countries"`
}

func Test_APIContinent(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/continents", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var cer ContinentExpectedResponses
	json.Unmarshal([]byte(rr.Body.String()), &cer)

	if len(cer) != 6 {
		t.Errorf("Missing continent in the response %v", cer)
	}

}
