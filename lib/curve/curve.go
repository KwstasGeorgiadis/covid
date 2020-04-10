package curve

//TODO add fucking caching you piece of shit and add expiration time

import (
	"sort"

	caching "../caching"
	pconf "../config"
	structs "../structs"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var (
	serverConf = pconf.GetAppConfig("./config/covid.json")
)

func requestData() ([]structs.CountryCurve, error) {
	client := &http.Client{}
	req, reqErr := http.NewRequest("GET", serverConf.API.URLHistory, nil)
	if reqErr != nil {
		return []structs.CountryCurve{}, reqErr
	}

	res, resError := client.Do(req)
	if resError != nil {
		return []structs.CountryCurve{}, resError
	}
	defer res.Body.Close()

	b, errorReadAll := ioutil.ReadAll(res.Body)
	if errorReadAll != nil {
		return []structs.CountryCurve{}, errorReadAll
	}

	keys := make([]structs.CountryCurve, 0)
	if errUnmarshal := json.Unmarshal(b, &keys); errUnmarshal != nil {
		return []structs.CountryCurve{}, errUnmarshal
	}

	return keys, nil
}

func GetAllCountries() ([]structs.CountryCurve, error) {

	pool := caching.NewPool()
	conn := pool.Get()
	defer conn.Close()

	cachedData, cacheGetError := caching.GetCurveData(conn)
	if cacheGetError != nil {
		return []structs.CountryCurve{}, cacheGetError
	}

	if len(cachedData) == 0 {
		data, err := requestData()
		if err != nil {
			return []structs.CountryCurve{}, err
		}
		caching.SetCurveData(conn, data)

		return data, err
	}

	return cachedData, nil
}

func GetCountry(name string) (structs.CountryCurve, error) {
	allCountries, errGetAllCountries := GetAllCountries()
	if errGetAllCountries != nil {
		return structs.CountryCurve{}, errGetAllCountries
	}

	if name == "UK" {
		for _, v := range allCountries {
			if v.Country == name && len(v.Province) == 0 {
				return v, nil
			}
		}
	}

	for _, v := range allCountries {
		if v.Country == name {
			return v, nil
		}
	}

	return structs.CountryCurve{}, nil
}

func GetDataByDate(date string) (map[string]interface{}, error) {
	dic := make(map[string]interface{})
	allCountries, errGetAllCountries := GetAllCountries()
	if errGetAllCountries != nil {
		return nil, errGetAllCountries
	}

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
	return dic, nil

}

//something weird here look into that
func DeathsPercentByDay(name string) {

	country, errGetCountry := GetCountry(name)
	if errGetCountry != nil {
		fmt.Println(errGetCountry)
	}

	var xs []float64

	for _, v := range country.Timeline.Deaths.(map[string]interface{}) {
		xs = append(xs, v.(float64))
	}
	sort.Float64s(xs) //sort keys alphabetically

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

func CompareDeathsCountries(nameOne string, nameTwo string) (structs.Compare, error) {

	country, errGetCountryOne := GetCountry(nameOne)
	if errGetCountryOne != nil {
		return structs.Compare{}, errGetCountryOne
	}

	countryTwo, errGetCountryTwo := GetCountry(nameTwo)
	if errGetCountryTwo != nil {
		return structs.Compare{}, errGetCountryTwo
	}

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

	return structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}, nil
}

func CompareDeathsFromFirstDeathCountries(nameOne string, nameTwo string) (structs.Compare, error) {

	country, errGetCountryOne := GetCountry(nameOne)
	if errGetCountryOne != nil {
		return structs.Compare{}, errGetCountryOne
	}

	countryTwo, errGetCountryTwo := GetCountry(nameTwo)
	if errGetCountryTwo != nil {
		return structs.Compare{}, errGetCountryTwo
	}

	var countrySortedDeath []float64
	var countryTwoSortedDeath []float64

	for _, v := range country.Timeline.Deaths.(map[string]interface{}) {
		if v.(float64) == 0 {
			continue
		}
		countrySortedDeath = append(countrySortedDeath, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Deaths.(map[string]interface{}) {
		if v.(float64) == 0 {
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

	return structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}, nil
}

func ComparePerCentDeathsCountries(nameOne string, nameTwo string) (structs.Compare, error) {

	country, errGetCountryOne := GetCountry(nameOne)
	if errGetCountryOne != nil {
		return structs.Compare{}, errGetCountryOne
	}

	countryTwo, errGetCountryTwo := GetCountry(nameTwo)
	if errGetCountryTwo != nil {
		return structs.Compare{}, errGetCountryTwo
	}

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

	return structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}, nil
}

func ComparePerDayDeathsCountries(nameOne string, nameTwo string) (structs.Compare, error) {

	country, errGetCountryOne := GetCountry(nameOne)
	if errGetCountryOne != nil {
		return structs.Compare{}, errGetCountryOne
	}

	countryTwo, errGetCountryTwo := GetCountry(nameTwo)
	if errGetCountryTwo != nil {
		return structs.Compare{}, errGetCountryTwo
	}

	var countrySortedDeath []float64
	var countryTwoSortedDeath []float64

	for _, v := range country.Timeline.Deaths.(map[string]interface{}) {
		if v.(float64) == 0 {
			continue
		}
		countrySortedDeath = append(countrySortedDeath, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Deaths.(map[string]interface{}) {
		if v.(float64) == 0 {
			continue
		}
		countryTwoSortedDeath = append(countryTwoSortedDeath, v.(float64))
	}
	sort.Float64s(countrySortedDeath)
	sort.Float64s(countryTwoSortedDeath)

	var tempCountryOneSortedDeath []float64
	for i := 0; i < len(countrySortedDeath); i++ {
		tempCountryOneSortedDeath = append(tempCountryOneSortedDeath, countrySortedDeath[i])
		if i == 0 {
			continue
		}

		countrySortedDeath[i] = countrySortedDeath[i] - tempCountryOneSortedDeath[i-1]
	}

	var tempCountryTwoSortedDeath []float64
	for i := 0; i < len(countryTwoSortedDeath); i++ {
		tempCountryTwoSortedDeath = append(tempCountryTwoSortedDeath, countryTwoSortedDeath[i])

		if i == 0 {
			continue
		}
		countryTwoSortedDeath[i] = countryTwoSortedDeath[i] - tempCountryTwoSortedDeath[i-1]
	}

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countrySortedDeath
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoSortedDeath

	return structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}, nil
}
