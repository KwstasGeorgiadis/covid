package stats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Country struct {
	Country            string  `json:"country"`
	Cases              int     `json:"cases"`
	TodayCases         int     `json:"todayCases"`
	Deaths             int     `json:"deaths"`
	TodayDeaths        int     `json:"todayDeaths"`
	Recovered          int     `json:"recovered"`
	Active             int     `json:"active"`
	Critical           int     `json:"critical"`
	CasesPerOneMillion float64 `json:"casesPerOneMillion"`
}

func requestData() []Country {
	url := "https://corona.lmao.ninja/countries?sort=country"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)

	//var obj Countries
	keys := make([]Country, 0)
	if err := json.Unmarshal(b, &keys); err != nil {
		panic(err)
	}

	return keys
}

func GetAllCountries() []Country {
	return requestData()
}

func GetCountry(name string) Country {
	allCountries := requestData()

	for _, v := range allCountries {
		if v.Country == name {
			return v
		}
	}

	return Country{}
}
