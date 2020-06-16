package hotspot

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/mux"
)

type HotspotExpectedResponse struct {
	MostCases struct {
		Country string `json:"country"`
		Data    []int  `json:"data"`
	} `json:"mostCases"`
	SecondCases struct {
		Country string `json:"country"`
		Data    []int  `json:"data"`
	} `json:"secondCases"`
	ThirdCases struct {
		Country string `json:"country"`
		Data    []int  `json:"data"`
	} `json:"thirdCases"`
	MostDeaths struct {
		Country string `json:"country"`
		Data    []int  `json:"data"`
	} `json:"mostDeaths"`
	SecondDeaths struct {
		Country string `json:"country"`
		Data    []int  `json:"data"`
	} `json:"secondDeaths"`
	ThirdDeaths struct {
		Country string `json:"country"`
		Data    []int  `json:"data"`
	} `json:"thirdDeaths"`
}

type ErrorExpectedResponse struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func Test_APIHotspotError(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/hotspot/ksdfsfd", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"days": "ksdfsfd",
	})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	var eer ErrorExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &eer)

	if eer.Code != 400 {
		t.Errorf("Wrong code value")
	}

	matched := strings.Contains(eer.Message, "ksdfsfd")
	if !matched {
		t.Errorf("Wrong message value")
	}

}

func Test_APIHotspotHandle(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/hotspot", nil)
	if err != nil {
		t.Fatal(err)
	}

	req = mux.SetURLVars(req, map[string]string{
		"days": "43",
	})
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Handle)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusBadRequest)
	}

	var her HotspotExpectedResponse
	json.Unmarshal([]byte(rr.Body.String()), &her)

	if len(her.MostCases.Data) != 43 {
		t.Errorf("Wrong ammount in most cases data")
	}

	if len(her.SecondCases.Data) != 43 {
		t.Errorf("Wrong ammount in second cases data ")
	}

	if len(her.ThirdCases.Data) != 43 {
		t.Errorf("Wrong ammount in third cases data ")
	}

	if len(her.MostDeaths.Data) != 43 {
		t.Errorf("Wrong ammount in most deaths data ")
	}

	if len(her.SecondDeaths.Data) != 43 {
		t.Errorf("Wrong ammount in second deaths data ")
	}

	if len(her.ThirdDeaths.Data) != 43 {
		t.Errorf("Wrong ammount in third deaths data ")
	}

	if &her.MostCases.Country == nil {
		t.Errorf("Something is wrong with the country field")
	}

	if calculateTotalAmmount(her.MostCases.Data) < calculateTotalAmmount(her.SecondCases.Data) {
		t.Errorf("MostCases' ammount is less than SecondCases'")
	}

	if calculateTotalAmmount(her.SecondCases.Data) < calculateTotalAmmount(her.ThirdCases.Data) {
		t.Errorf("SecondCases' ammount is less than ThirdCases'")
	}

	if calculateTotalAmmount(her.MostDeaths.Data) < calculateTotalAmmount(her.SecondDeaths.Data) {
		t.Errorf("MostDeaths' ammount is less than SecondDeaths'")
	}

	if calculateTotalAmmount(her.SecondDeaths.Data) < calculateTotalAmmount(her.ThirdDeaths.Data) {
		t.Errorf("SecondDeaths' ammount is less than ThirdDeaths'")
	}
}

func calculateTotalAmmount(arr []int) int {
	total := 0
	for _, v := range arr {
		total = total + v
	}
	return total
}
