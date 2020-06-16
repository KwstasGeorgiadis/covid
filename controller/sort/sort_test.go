package sortcon

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SortExpectedResponse struct {
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

func Test_APISortByDeaths(t *testing.T) {
	var jsonStr = []byte(`{"type" : "deaths"}`)

	req, err := http.NewRequest("POST", "/api/sort", bytes.NewBuffer(jsonStr))
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

	var ser SortExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &ser)

	if len(ser.Data) <= 0 {
		t.Errorf("Missing data in the response")
	}

	if ser.Data[0].Deaths < ser.Data[1].Deaths {
		t.Errorf("Deaths seems not to be sorted when type is deaths")
	}
}

func Test_APISortByCases(t *testing.T) {
	var jsonStr = []byte(`{"type" : "cases"}`)

	req, err := http.NewRequest("POST", "/api/sort", bytes.NewBuffer(jsonStr))
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

	var ser SortExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &ser)

	if len(ser.Data) <= 0 {
		t.Errorf("Missing data in the response")
	}

	if ser.Data[0].Cases < ser.Data[1].Cases {
		t.Errorf("Cases seems not to be sorted when type is cases")
	}
}

func Test_APISortByTodayCases(t *testing.T) {
	var jsonStr = []byte(`{"type" : "todayCases"}`)

	req, err := http.NewRequest("POST", "/api/sort", bytes.NewBuffer(jsonStr))
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

	var ser SortExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &ser)

	if len(ser.Data) <= 0 {
		t.Errorf("Missing data in the response")
	}

	if ser.Data[0].TodayCases < ser.Data[1].TodayCases {
		t.Errorf("TodayCases seems not to be sorted when type is cases")
	}
}

func Test_APISortByTodayDeaths(t *testing.T) {
	var jsonStr = []byte(`{"type" : "todayDeaths"}`)

	req, err := http.NewRequest("POST", "/api/sort", bytes.NewBuffer(jsonStr))
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

	var ser SortExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &ser)

	if len(ser.Data) <= 0 {
		t.Errorf("Missing data in the response")
	}

	if ser.Data[0].TodayDeaths < ser.Data[1].TodayDeaths {
		t.Errorf("TodayDeaths seems not to be sorted when type is todayDeaths")
	}
}

func Test_APISortByTodayRecovered(t *testing.T) {
	var jsonStr = []byte(`{"type" : "recovered"}`)

	req, err := http.NewRequest("POST", "/api/sort", bytes.NewBuffer(jsonStr))
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

	var ser SortExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &ser)

	if len(ser.Data) <= 0 {
		t.Errorf("Missing data in the response")
	}

	if ser.Data[0].Recovered < ser.Data[1].Recovered {
		t.Errorf("Recovered seems not to be sorted when type is recovered")
	}
}

func Test_APISortByTodayActive(t *testing.T) {
	var jsonStr = []byte(`{"type" : "active"}`)

	req, err := http.NewRequest("POST", "/api/sort", bytes.NewBuffer(jsonStr))
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

	var ser SortExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &ser)

	if len(ser.Data) <= 0 {
		t.Errorf("Missing data in the response")
	}

	if ser.Data[0].Active < ser.Data[1].Active {
		t.Errorf("Active seems not to be sorted when type is active")
	}
}

func Test_APISortByTodayCritical(t *testing.T) {
	var jsonStr = []byte(`{"type" : "critical"}`)

	req, err := http.NewRequest("POST", "/api/sort", bytes.NewBuffer(jsonStr))
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

	var ser SortExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &ser)

	if len(ser.Data) <= 0 {
		t.Errorf("Missing data in the response")
	}

	if ser.Data[0].Critical < ser.Data[1].Critical {
		t.Errorf("Critical seems not to be sorted when type is critical")
	}
}

func Test_APISortByTodayCasesPerOneMillion(t *testing.T) {
	var jsonStr = []byte(`{"type" : "casesPerOneMillion"}`)

	req, err := http.NewRequest("POST", "/api/sort", bytes.NewBuffer(jsonStr))
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

	var ser SortExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &ser)

	if len(ser.Data) <= 0 {
		t.Errorf("Missing data in the response")
	}

	if ser.Data[0].CasesPerOneMillion < ser.Data[1].CasesPerOneMillion {
		t.Errorf("Critical seems not to be sorted when type is casesPerOneMillion")
	}
}
