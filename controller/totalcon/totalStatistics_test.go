package totalcon

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type TotalExpectedResponse struct {
	TodayPerCentOfTotalCases  int `json:"todayPerCentOfTotalCases"`
	TodayPerCentOfTotalDeaths int `json:"todayPerCentOfTotalDeaths"`
	TotalCases                int `json:"totalCases"`
	TotalDeaths               int `json:"totalDeaths"`
	TodayTotalCases           int `json:"todayTotalCases"`
	TodayTotalDeaths          int `json:"todayTotalDeaths"`
}

func Test_APITotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/total", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(TotalHandle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var ter TotalExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &ter)

	if ter.TotalCases <= 0 {
		t.Errorf("Total cases seems to be broken")
	}

	if ter.TotalDeaths <= 0 {
		t.Errorf("Total deaths seems to be broken")
	}

	if &ter.TodayPerCentOfTotalCases == nil {
		t.Errorf("Today per cent Of total cases seems to be broken")
	}

	if &ter.TodayPerCentOfTotalDeaths == nil {
		t.Errorf("Today per cent of total cases seems to be broken")
	}

	if &ter.TodayTotalCases == nil {
		t.Errorf("Today per cent of total cases deaths seems to be broken")
	}

	if &ter.TodayTotalDeaths == nil {
		t.Errorf("Today per cent of total cases deaths seems to be broken")
	}
}
