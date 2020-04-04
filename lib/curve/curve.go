package curve

import (
	"sort"

	pconf "../config"
	structs "../structs"
	//caching "../caching"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	serverConf = pconf.GetAppConfig("./config/covid.json")
)

func requestData() []structs.CountryCurve {
	//caching.ok
	client := &http.Client{}
	req, err := http.NewRequest("GET", serverConf.API.URL_history, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)

	//var obj Countries
	keys := make([]structs.CountryCurve, 0)
	if err := json.Unmarshal(b, &keys); err != nil {
		panic(err)
	}

	return keys
}

func GetAllCountries() []structs.CountryCurve {
	s := requestData()
	return s
}

func GetCountry(name string) structs.CountryCurve {
	allCountries := GetAllCountries()

	for _, v := range allCountries {
		if v.Country == name {
			return v
		}
	}

	return structs.CountryCurve{}
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
}

func CompareDeathsCountries(nameOne string, nameTwo string) structs.Compare {

	country := GetCountry(nameOne)
	countryTwo := GetCountry(nameTwo)

	var countrySortedDeath []float64
	var countryTwoSortedDeath []float64

	for _, v := range country.Timeline.Deaths.(map[string]interface{}) {
		countrySortedDeath = append(countrySortedDeath, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Deaths.(map[string]interface{}) {
		countryTwoSortedDeath = append(countryTwoSortedDeath, v.(float64))
	}
	sort.Float64s(countrySortedDeath)
	sort.Float64s(countryTwoSortedDeath)

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countrySortedDeath
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoSortedDeath

	return structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}
}

func CompareDeathsFromFirstDeathCountries(nameOne string, nameTwo string) structs.Compare {

	country := GetCountry(nameOne)
	countryTwo := GetCountry(nameTwo)

	var countrySortedDeath []float64
	var countryTwoSortedDeath []float64

	for _, v := range country.Timeline.Deaths.(map[string]interface{}) {
		if (v.(float64) == 0){
			continue
		}
		countrySortedDeath = append(countrySortedDeath, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Deaths.(map[string]interface{}) {
		if (v.(float64) == 0){
			continue
		}
		countryTwoSortedDeath = append(countryTwoSortedDeath, v.(float64))
	}
	sort.Float64s(countrySortedDeath)
	sort.Float64s(countryTwoSortedDeath)

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countrySortedDeath
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoSortedDeath

	return structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}
}

func ComparePerCentDeathsCountries(nameOne string, nameTwo string) structs.Compare {

	country := GetCountry(nameOne)
	countryTwo := GetCountry(nameTwo)

	var countrySortedDeath []float64
	var countryTwoSortedDeath []float64

	for _, v := range country.Timeline.Deaths.(map[string]interface{}) {
		countrySortedDeath = append(countrySortedDeath, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Deaths.(map[string]interface{}) {
		countryTwoSortedDeath = append(countryTwoSortedDeath, v.(float64))
	}
	sort.Float64s(countrySortedDeath)
	sort.Float64s(countryTwoSortedDeath)

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	for _, v := range country.Timeline.Deaths.(map[string]interface{}) {
		countrySortedDeath = append(countrySortedDeath, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Deaths.(map[string]interface{}) {
		countryTwoSortedDeath = append(countryTwoSortedDeath, v.(float64))
	}

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countrySortedDeath
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoSortedDeath

	return structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}
}

func ComparePerDayDeathsCountries(nameOne string, nameTwo string) structs.Compare {

	country := GetCountry(nameOne)
	countryTwo := GetCountry(nameTwo)

	var countrySortedDeath []float64
	var countryTwoSortedDeath []float64

	for _, v := range country.Timeline.Deaths.(map[string]interface{}) {
		if (v.(float64) == 0){
			continue
		}
		countrySortedDeath = append(countrySortedDeath, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Deaths.(map[string]interface{}) {
		if (v.(float64) == 0){
			continue
		}
		countryTwoSortedDeath = append(countryTwoSortedDeath, v.(float64))
	}
	sort.Float64s(countrySortedDeath)
	sort.Float64s(countryTwoSortedDeath)

	var tempCountryOneSortedDeath []float64
	for i := 0; i < len(countrySortedDeath); i++ {
		tempCountryOneSortedDeath = append(tempCountryOneSortedDeath, countrySortedDeath [i])
		if i == 0{
			continue
		}

		countrySortedDeath [i] = countrySortedDeath [i] - tempCountryOneSortedDeath [i-1] 
	}

	var tempCountryTwoSortedDeath []float64
	for i := 0; i < len(countryTwoSortedDeath); i++ {
		tempCountryTwoSortedDeath = append(tempCountryTwoSortedDeath, countryTwoSortedDeath [i])

		if i == 0{
			continue
		}
		countryTwoSortedDeath [i] = countryTwoSortedDeath [i] - tempCountryTwoSortedDeath [i-1] 
	}


	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countrySortedDeath
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoSortedDeath

	return structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}
}
