package cssectl

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	mcsse "github.com/junkd0g/covid/lib/model/csse"
)

func Test_APICsse(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/csse", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"country": "Russia",
	})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var cer mcsse.CSEECountryResponse
	json.Unmarshal([]byte(rr.Body.String()), &cer)

	if cer.Country != "Russia" {
		t.Errorf("Wrong country field %s", cer.Country)
	}

	if len(cer.Data) < 1 {
		t.Errorf("No data")
	}

}

func Test_APIDifferentNameCsse(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/csse", nil)
	if err != nil {
		t.Fatal(err)
	}
	req = mux.SetURLVars(req, map[string]string{
		"country": "USA",
	})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	var cer mcsse.CSEECountryResponse
	json.Unmarshal([]byte(rr.Body.String()), &cer)

	if cer.Country != "US" {
		t.Errorf("Wrong country field %s", cer.Country)
	}

	if len(cer.Data) < 1 {
		t.Errorf("No data")
	}

}
