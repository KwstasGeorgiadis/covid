package hotspot

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
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

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HotspotHandle)
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

	//matched, _ := regexp.MatchString(eer.Message, "ksdfsfd")
	//t.Error(eer.Message)
	//if !matched {
	//	t.Errorf("Wrong message value")
	//}

}
