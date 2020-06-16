package countrycon

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CountryExpectedResponse struct {
	Country            string `json:"country"`
	Cases              int    `json:"cases"`
	TodayCases         int    `json:"todayCases"`
	Deaths             int    `json:"deaths"`
	TodayDeaths        int    `json:"todayDeaths"`
	Recovered          int    `json:"recovered"`
	Active             int    `json:"active"`
	Critical           int    `json:"critical"`
	CasesPerOneMillion int    `json:"casesPerOneMillion"`
	Tests              int    `json:"tests"`
	TestsPerOneMillion int    `json:"testsPerOneMillion"`
}

func Test_APIAllCountries(t *testing.T) {
	var jsonStr = []byte(`{"country" : "Greece"}`)

	req, err := http.NewRequest("POST", "/api/country", bytes.NewBuffer(jsonStr))
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Country)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var cer CountryExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &cer)

	if cer.Country != "Greece" {
		t.Errorf("Country field seems to be broken has value %s but expected value is %s", cer.Country, "Greece")
	}

	if cer.Cases <= 0 {
		t.Errorf("Cases field seems to be broken")
	}

	if cer.Deaths <= 0 {
		t.Errorf("Cases field seems to be broken")
	}

	if cer.Tests <= 0 {
		t.Errorf("Tests field seems to be broken")
	}
}
