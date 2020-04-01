package compare

import (
	"encoding/json"
	"fmt"

	curve "../../lib/curve"

	"io/ioutil"
	"net/http"
)

type CompareRequest struct {
	NameOne string `json:"countryOne"`
	NameTwo string `json:"countryTwo"`
}

func Perform(r *http.Request) ([]byte, int) {
	var compareRequest CompareRequest
	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		// return some 500 stuff here
		fmt.Println(errIoutilReadAll.Error())
	}

	json.Unmarshal(b, &compareRequest)

	country := curve.CompareDeathsCountries(compareRequest.NameOne, compareRequest.NameTwo)
	jsonBody, _ := json.Marshal(country)

	return jsonBody, 200
}
