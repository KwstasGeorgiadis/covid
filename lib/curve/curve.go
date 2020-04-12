package curve

//TODO add fucking caching you piece of shit and add expiration time

import (
	"sort"

	caching "../caching"
	pconf "../config"
	structs "../structs"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	serverConf = pconf.GetAppConfig("./config/covid.json")
)

// requestData does an HTTP GET request to the third party API that 
// contains covid-9 stats ' history (per day from 22/01/2020)
// It returns []structs.Country and any write error encountered.
func requestHistoryData() ([]structs.CountryCurve, error) {
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

// GetAllCountries returns an array of all countries per day
// Covid-19 stats (data starts from date 22/01/2020)
// Check if there are cached data if not does a HTTP
// request to the 3rd party API (check requestHistoryData())
// It returns []structs.CountryCurve and any write error encountered.
func GetAllCountries() ([]structs.CountryCurve, error) {

	pool := caching.NewPool()
	conn := pool.Get()
	defer conn.Close()

	cachedData, cacheGetError := caching.GetCurveData(conn)
	if cacheGetError != nil {
		return []structs.CountryCurve{}, cacheGetError
	}

	if len(cachedData) == 0 {
		data, err := requestHistoryData()
		if err != nil {
			return []structs.CountryCurve{}, err
		}
		caching.SetCurveData(conn, data)

		return data, err
	}

	return cachedData, nil
}

// GetCountry seach through an array of structs.CountryCurve and
// gets COVID-19 per day stats for that specific country
// It returns structs.CountryCurve
 and any write error encountered.
func GetCountry(name string) (structs.CountryCurve, error) {
	allCountries, errGetAllCountries := GetAllCountries()
	if errGetAllCountries != nil {
		return structs.CountryCurve{}, errGetAllCountries
	}

	for _, v := range allCountries {
		if v.Country == "UK" {
			if len(v.Province) == 0 {
				return v, nil
			}
			continue
		}
		if v.Country == name {
			return v, nil
		}
	}

	return structs.CountryCurve{}, nil
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
