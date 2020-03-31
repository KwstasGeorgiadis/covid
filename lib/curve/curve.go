package curve

import (
	"sort"

	pconf "../config"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Country struct {
	Country  string         `json:"country"`
	Timeline TimelineStruct `json:"timeline"`
}

type TimelineStruct struct {
	Cases     interface{} `json:"cases"`
	Deaths    interface{} `json:"deaths"`
	Recovered interface{} `json:"recovered"`
}

var (
	serverConf = pconf.GetAppConfig("./config/covid.json")
)

func requestData() []Country {

	client := &http.Client{}
	req, err := http.NewRequest("GET", serverConf.API.URL_history, nil)

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
	s := requestData()
	return s
}

func GetCountry(name string) Country {
	allCountries := GetAllCountries()

	for _, v := range allCountries {
		if v.Country == name {
			return v
		}
	}

	return Country{}
}

func GetDataByDate(date string) map[string]interface{} {
	dic := make(map[string]interface{})
	allCountries := GetAllCountries()

	var cases float64
	var deaths float64
	var recovered float64

	for _, v := range allCountries {

		for kk, vv := range v.Timeline.Cases.(map[string]interface{}) {
			if kk == date {
				cases = vv.(float64)
			}
		}

		for kk, vv := range v.Timeline.Deaths.(map[string]interface{}) {
			if kk == date {
				deaths = vv.(float64)
			}
		}

		for kk, vv := range v.Timeline.Recovered.(map[string]interface{}) {
			if kk == date {
				recovered = vv.(float64)
			}
		}
		dic[v.Country] = map[string]interface{}{
			"cases":     cases,
			"deaths":    deaths,
			"recovered": recovered,
		}
	}
	return dic

}

func DeathsPercentByDay(name string) {

	country := GetCountry(name)

	var xs []float64

	for _, v := range country.Timeline.Deaths.(map[string]interface{}) {
		xs = append(xs, v.(float64))
	}
	sort.Float64s(xs) //sort keys alphabetically
	fmt.Println(xs)

	for _, v := range xs {
		if v == 0 {
			continue
		}

	}

	for i := 0; i < len(xs); i++ {
		if xs[i] == 0 {
			continue
		}
		if i < (len(xs)-1) && xs[i-1] > 0 {
			fmt.Println(100 * ((float64(xs[i]) - float64(xs[i-1])) / float64(xs[i-1])))
		}
	}
	//if
	//}

}
