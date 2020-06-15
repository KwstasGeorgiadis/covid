package allcountries

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type AllCountriesExpectedResponse struct {
	Countries []string `json:"countries"`
}

func Test_APIAllCountries(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/countries/all", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(AllCountriesHandle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var acer AllCountriesExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &acer)
	if len(acer.Countries) <= 0 {
		t.Errorf("Countries seems to be broken")
	}
}
