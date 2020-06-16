package worldct

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type ExpectedStructureResponse struct {
	Cases          []int `json:"cases"`
	Deaths         []int `json:"deaths"`
	Recovered      []int `json:"recovered"`
	CasesDaily     []int `json:"casesDaily"`
	DeathsDaily    []int `json:"deathsDaily"`
	RecoveredDaily []int `json:"recoveredDaily"`
}

func Test_APIWorld(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/world", nil)
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

	var esr *ExpectedStructureResponse
	json.Unmarshal([]byte(rr.Body.String()), &esr)
	if len(esr.Cases) <= 0 {
		t.Errorf("Cases array is empty")
	}

	if len(esr.Deaths) <= 0 {
		t.Errorf("Deaths array is empty")
	}

	if len(esr.Recovered) <= 0 {
		t.Errorf("Recovered  array is empty")
	}

	if len(esr.CasesDaily) <= 0 {
		t.Errorf("CasesDaily array is empty")
	}

	if len(esr.DeathsDaily) <= 0 {
		t.Errorf("DeathsDaily array is empty")
	}

	if len(esr.RecoveredDaily) <= 0 {
		t.Errorf("RecoveredDaily array is empty")
	}
}
