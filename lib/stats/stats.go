package stats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	pconf "../config"
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

type Countries struct {
	Data	[]Country  `json:"data"`
}

var (
	serverConf = pconf.GetAppConfig("./config/covid.json")
)

func requestData() []Country {

	client := &http.Client{}
	req, err := http.NewRequest("GET", serverConf.API.URL, nil)

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

func GetAllCountries() Countries {
	s := Countries{Data:  requestData()}
	return s
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
