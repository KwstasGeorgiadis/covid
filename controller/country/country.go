package countryCon

import (
	"encoding/json"
	"fmt"

	stats "../../lib/stats"

	"io/ioutil"
	"net/http"
)

type CountryRequest struct {
	Name string `json:"country"`
}

func Perform(r *http.Request) ([]byte, int) {
	var countryRequest CountryRequest

	b, errIoutilReadAll := ioutil.ReadAll(r.Body)
	if errIoutilReadAll != nil {
		// return some 500 stuff here
		fmt.Println(errIoutilReadAll.Error())
	}

	json.Unmarshal(b, &countryRequest)

	country := stats.GetCountry(countryRequest.Name)
	jsonBody, _ := json.Marshal(country)

	return jsonBody, 200
}
