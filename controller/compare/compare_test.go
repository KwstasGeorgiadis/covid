package comparectl

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CompareExpectedResponse struct {
	CountryOne struct {
		Country             string `json:"country"`
		DataDeaths          []int  `json:"dataDeaths"`
		DataDeathsFromFirst []int  `json:"dataDeathsFromFirst"`
		DataDeathsPerDay    []int  `json:"dataDeathsPerDay"`
		DataRecoverd        []int  `json:"dataRecoverd"`
		DataCases           []int  `json:"dataCases"`
		DataCasesFromFirst  []int  `json:"dataCasesFromFirst"`
	} `json:"countryOne"`
	CountryTwo struct {
		Country             string `json:"country"`
		DataDeaths          []int  `json:"dataDeaths"`
		DataDeathsFromFirst []int  `json:"dataDeathsFromFirst"`
		DataDeathsPerDay    []int  `json:"dataDeathsPerDay"`
		DataRecoverd        []int  `json:"dataRecoverd"`
		DataCases           []int  `json:"dataCases"`
		DataCasesFromFirst  []int  `json:"dataCasesFromFirst"`
	} `json:"countryTwo"`
}

func Test_APIAllCompare(t *testing.T) {
	var jsonStr = []byte(`{
		"countryOne" : "Spain",
		"countryTwo" : "Italy"
	}`)

	req, err := http.NewRequest("POST", "/api/compare/all", bytes.NewBuffer(jsonStr))
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

	var cer CompareExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &cer)

	if cer.CountryOne.Country != "Spain" {
		t.Errorf("CountryOne.Country  field seems to be broken has value %s but expected value is %s", cer.CountryOne.Country, "Spain")
	}

	if cer.CountryTwo.Country != "Italy" {
		t.Errorf("Country field seems to be broken has value %s but expected value is %s", cer.CountryTwo.Country, "Italy")
	}
}
