package countriescon

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type CountriesExpectedResponse struct {
	Data []struct {
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
	} `json:"data"`
}

func Test_APITotal(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/countries", nil)
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

	var cer CountriesExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &cer)

	if len(cer.Data) <= 0 {
		t.Errorf("Data seems to be broken")
	}

	if cer.Data[0].Country == "" {
		t.Errorf("Country field looks broken")
	}

	if &cer.Data[0].Cases == nil {
		t.Errorf("Cases field looks broken")
	}
}
