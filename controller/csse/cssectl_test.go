package cssectl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mcsse "github.com/junkd0g/covid/lib/model/csse"
)

func Test_APICsse(t *testing.T) {
	req, err := http.NewRequest("GET", "/api/csse/Russia", nil)
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

	var cer mcsse.CSEECountryResponse
	json.Unmarshal([]byte(rr.Body.String()), &cer)
	fmt.Println(rr.Body.String())
	if cer.Country != "Russia" {
		t.Errorf("Wrong country field %s", cer.Country)
	}

	if len(cer.Data) < 1 {
		t.Errorf("No data")
	}

}
