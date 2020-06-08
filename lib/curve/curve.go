package curve

import (
	"sort"

	applogger "github.com/junkd0g/covid/lib/applogger"
	caching "github.com/junkd0g/covid/lib/caching"
	pconf "github.com/junkd0g/covid/lib/config"
	structs "github.com/junkd0g/covid/lib/structs"

	"encoding/json"
	"io/ioutil"
	"net/http"
)

var (
	serverConf pconf.AppConf
)

func init() {
	serverConf = pconf.GetAppConfig()
}

// requestData does an HTTP GET request to the third party API that
// contains covid-9 stats ' history (per day from 22/01/2020)
// It returns []structs.Country and any write error encountered.
func requestHistoryData() ([]structs.CountryCurve, error) {
	client := &http.Client{}
	requestURL := serverConf.API.URLHistory

	req, reqErr := http.NewRequest("GET", requestURL, nil)
	if reqErr != nil {
		applogger.Log("ERROR", "curve", "requestHistoryData", reqErr.Error())
		return []structs.CountryCurve{}, reqErr
	}

	res, resError := client.Do(req)
	if resError != nil {
		applogger.Log("ERROR", "curve", "requestHistoryData", resError.Error())
		return []structs.CountryCurve{}, resError
	}
	defer res.Body.Close()

	b, errorReadAll := ioutil.ReadAll(res.Body)
	if errorReadAll != nil {
		applogger.Log("ERROR", "curve", "requestHistoryData", errorReadAll.Error())
		return []structs.CountryCurve{}, errorReadAll
	}

	keys := make([]structs.CountryCurve, 0)
	if errUnmarshal := json.Unmarshal(b, &keys); errUnmarshal != nil {
		applogger.Log("ERROR", "curve", "requestHistoryData", errUnmarshal.Error())
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
		applogger.Log("ERROR", "curve", "GetAllCountries", cacheGetError.Error())
		return []structs.CountryCurve{}, cacheGetError
	}

	if len(cachedData) == 0 {
		applogger.Log("INFO", "stats", "GetAllCountries", "Request data instead of getting cached data")
		data, err := requestHistoryData()
		if err != nil {
			applogger.Log("ERROR", "curve", "GetAllCountries", err.Error())
			return []structs.CountryCurve{}, err
		}
		caching.SetCurveData(conn, data)
		return data, nil
	}

	return cachedData, nil
}

// GetCountry seach through an array of structs.CountryCurve and
// gets COVID-19 per day stats for that specific country
// It returns structs.CountryCurve and any write error encountered.
func GetCountry(name string) (structs.CountryCurve, error) {
	allCountries, errGetAllCountries := GetAllCountries()
	if errGetAllCountries != nil {
		applogger.Log("ERROR", "curve", "GetCountry", errGetAllCountries.Error())
		return structs.CountryCurve{}, errGetAllCountries
	}

	for _, v := range allCountries {
		if name == "UK" && v.Country == "UK" {
			if len(v.Province) == 0 {
				return v, nil
			}
			continue
		}
		if name == "France" && v.Country == "France" {
			if v.Province == "" {
				return v, nil
			}
			continue
		}

		if v.Country == name {
			return v, nil
		}
	}

	applogger.Log("WARN", "curve", "GetCountry", "Returning empty country")
	return structs.CountryCurve{}, nil
}

// GetCountryBP
func GetCountryBP(name string, allCountries []structs.CountryCurve) (structs.CountryCurve, error) {

	for _, v := range allCountries {
		if name == "UK" && v.Country == "UK" {
			if len(v.Province) == 0 {
				return v, nil
			}
			continue
		}
		if name == "France" && v.Country == "France" {
			if v.Province == "" {
				return v, nil
			}
			continue
		}

		if v.Country == name {
			return v, nil
		}
	}

	applogger.Log("WARN", "curve", "GetCountry", "Returning empty country")
	return structs.CountryCurve{}, nil
}

// CompareDeathsCountries returns two integer arrays (one per country passed
// in parameter) which contain total number of deaths from  22/01/2020
// It returns structs.Compare and any write error encountered.
func CompareDeathsCountries(nameOne string, nameTwo string, getCountryST structs.CountryCurve, getCountryStTwo structs.CountryCurve) (structs.Compare, error) {
	var country structs.CountryCurve
	var errGetCountryOne error

	if (getCountryST == structs.CountryCurve{}) {
		country, errGetCountryOne = GetCountry(nameOne)
		if errGetCountryOne != nil {
			applogger.Log("ERROR", "curve", "CompareDeathsCountries", errGetCountryOne.Error())
			return structs.Compare{}, errGetCountryOne
		}
	} else {
		country = getCountryST
	}

	var countryTwo structs.CountryCurve
	var errGetCountryTwo error

	if (getCountryStTwo == structs.CountryCurve{}) {
		countryTwo, errGetCountryTwo = GetCountry(nameTwo)
		if errGetCountryTwo != nil {
			applogger.Log("ERROR", "curve", "CompareDeathsCountries", errGetCountryTwo.Error())
			return structs.Compare{}, errGetCountryTwo
		}
	} else {
		countryTwo = getCountryStTwo
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

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}

	return compareStructs, nil
}

// CompareDeathsFromFirstDeathCountries returns two integer arrays (one per country passed
// in parameter) which contain total number of deaths from  the first confirm death.
// It returns structs.Compare and any write error encountered.
func CompareDeathsFromFirstDeathCountries(nameOne string, nameTwo string, getCountryST structs.CountryCurve, getCountryStTwo structs.CountryCurve) (structs.Compare, error) {

	var country structs.CountryCurve
	var errGetCountryOne error

	if (getCountryST == structs.CountryCurve{}) {
		country, errGetCountryOne = GetCountry(nameOne)
		if errGetCountryOne != nil {
			applogger.Log("ERROR", "curve", "CompareDeathsFromFirstDeathCountries", errGetCountryOne.Error())
			return structs.Compare{}, errGetCountryOne
		}
	} else {
		country = getCountryST
	}

	var countryTwo structs.CountryCurve
	var errGetCountryTwo error

	if (getCountryStTwo == structs.CountryCurve{}) {
		countryTwo, errGetCountryTwo = GetCountry(nameTwo)
		if errGetCountryTwo != nil {
			applogger.Log("ERROR", "curve", "CompareDeathsFromFirstDeathCountries", errGetCountryTwo.Error())
			return structs.Compare{}, errGetCountryTwo
		}
	} else {
		countryTwo = getCountryStTwo
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

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}

	return compareStructs, nil
}

// ComparePerDayDeathsCountries returns two integer arrays (one per country passed
// in parameter) which contain unique per day number of deaths from first confrim death
// It returns structs.Compare and any write error encountered.
func ComparePerDayDeathsCountries(nameOne string, nameTwo string, getCountryST structs.CountryCurve, getCountryStTwo structs.CountryCurve) (structs.Compare, error) {

	var country structs.CountryCurve
	var errGetCountryOne error

	if (getCountryST == structs.CountryCurve{}) {
		country, errGetCountryOne = GetCountry(nameOne)
		if errGetCountryOne != nil {
			applogger.Log("ERROR", "curve", "ComparePerDayDeathsCountries", errGetCountryOne.Error())
			return structs.Compare{}, errGetCountryOne
		}
	} else {
		country = getCountryST
	}

	var countryTwo structs.CountryCurve
	var errGetCountryTwo error

	if (getCountryStTwo == structs.CountryCurve{}) {
		countryTwo, errGetCountryTwo = GetCountry(nameTwo)
		if errGetCountryTwo != nil {
			applogger.Log("ERROR", "curve", "ComparePerDayDeathsCountries", errGetCountryTwo.Error())
			return structs.Compare{}, errGetCountryTwo
		}
	} else {
		countryTwo = getCountryStTwo
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

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}
	return compareStructs, nil
}

// CompareRecoveryCountries returns two integer arrays (one per country passed
// in parameter) which contain total number of recovery patients from  22/01/2020
// It returns structs.Compare and any write error encountered.
func CompareRecoveryCountries(nameOne string, nameTwo string, getCountryST structs.CountryCurve, getCountryStTwo structs.CountryCurve) (structs.Compare, error) {
	var country structs.CountryCurve
	var errGetCountryOne error

	if (getCountryST == structs.CountryCurve{}) {
		country, errGetCountryOne = GetCountry(nameOne)
		if errGetCountryOne != nil {
			applogger.Log("ERROR", "curve", "CompareRecoveryCountries", errGetCountryOne.Error())
			return structs.Compare{}, errGetCountryOne
		}
	} else {
		country = getCountryST
	}

	var countryTwo structs.CountryCurve
	var errGetCountryTwo error

	if (getCountryStTwo == structs.CountryCurve{}) {
		countryTwo, errGetCountryTwo = GetCountry(nameTwo)
		if errGetCountryTwo != nil {
			applogger.Log("ERROR", "curve", "CompareRecoveryCountries", errGetCountryTwo.Error())
			return structs.Compare{}, errGetCountryTwo
		}
	} else {
		countryTwo = getCountryStTwo
	}

	var countrySortedRecovery []float64
	var countryTwoSortedRecovery []float64

	for _, v := range country.Timeline.Recovered.(map[string]interface{}) {
		countrySortedRecovery = append(countrySortedRecovery, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Recovered.(map[string]interface{}) {
		countryTwoSortedRecovery = append(countryTwoSortedRecovery, v.(float64))
	}
	sort.Float64s(countrySortedRecovery)
	sort.Float64s(countryTwoSortedRecovery)

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countrySortedRecovery
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoSortedRecovery

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}

	return compareStructs, nil
}

// CompareCasesCountries returns two integer arrays (one per country passed
// in parameter) which contain total number of cases from  22/01/2020
// It returns structs.Compare and any write error encountered.
func CompareCasesCountries(nameOne string, nameTwo string, getCountryST structs.CountryCurve, getCountryStTwo structs.CountryCurve) (structs.Compare, error) {
	var country structs.CountryCurve
	var errGetCountryOne error

	if (getCountryST == structs.CountryCurve{}) {
		country, errGetCountryOne = GetCountry(nameOne)
		if errGetCountryOne != nil {
			applogger.Log("ERROR", "curve", "CompareCasesCountries", errGetCountryOne.Error())
			return structs.Compare{}, errGetCountryOne
		}
	} else {
		country = getCountryST
	}

	var countryTwo structs.CountryCurve
	var errGetCountryTwo error

	if (getCountryStTwo == structs.CountryCurve{}) {
		countryTwo, errGetCountryTwo = GetCountry(nameTwo)
		if errGetCountryTwo != nil {
			applogger.Log("ERROR", "curve", "CompareCasesCountries", errGetCountryTwo.Error())
			return structs.Compare{}, errGetCountryTwo
		}
	} else {
		countryTwo = getCountryStTwo
	}

	var countrySortedCases []float64
	var countryTwoSortedCases []float64

	for _, v := range country.Timeline.Cases.(map[string]interface{}) {
		countrySortedCases = append(countrySortedCases, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Cases.(map[string]interface{}) {
		countryTwoSortedCases = append(countryTwoSortedCases, v.(float64))
	}
	sort.Float64s(countrySortedCases)
	sort.Float64s(countryTwoSortedCases)

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countrySortedCases
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoSortedCases

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}

	return compareStructs, nil
}

// ComparePerDayCasesCountries returns two integer arrays (one per country passed
// in parameter) which contain unique per day number of case from first confrim case
// It returns structs.Compare and any write error encountered.
func ComparePerDayCasesCountries(nameOne string, nameTwo string, getCountryST structs.CountryCurve, getCountryStTwo structs.CountryCurve) (structs.Compare, error) {

	var country structs.CountryCurve
	var errGetCountryOne error

	if (getCountryST == structs.CountryCurve{}) {
		country, errGetCountryOne = GetCountry(nameOne)
		if errGetCountryOne != nil {
			applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", errGetCountryOne.Error())
			return structs.Compare{}, errGetCountryOne
		}
	} else {
		country = getCountryST
	}

	var countryTwo structs.CountryCurve
	var errGetCountryTwo error

	if (getCountryStTwo == structs.CountryCurve{}) {
		countryTwo, errGetCountryTwo = GetCountry(nameTwo)
		if errGetCountryTwo != nil {
			applogger.Log("ERROR", "curve", "ComparePerDayCasesCountries", errGetCountryTwo.Error())
			return structs.Compare{}, errGetCountryTwo
		}
	} else {
		countryTwo = getCountryStTwo
	}

	var countrySortedCases []float64
	var countryTwoSortedCases []float64

	for _, v := range country.Timeline.Cases.(map[string]interface{}) {
		if v.(float64) == 0 {
			continue
		}
		countrySortedCases = append(countrySortedCases, v.(float64))
	}
	for _, v := range countryTwo.Timeline.Cases.(map[string]interface{}) {
		if v.(float64) == 0 {
			continue
		}
		countryTwoSortedCases = append(countryTwoSortedCases, v.(float64))
	}
	sort.Float64s(countrySortedCases)
	sort.Float64s(countryTwoSortedCases)

	var tempCountryOneSortedCases []float64
	for i := 0; i < len(countrySortedCases); i++ {
		tempCountryOneSortedCases = append(tempCountryOneSortedCases, countrySortedCases[i])
		if i == 0 {
			continue
		}

		countrySortedCases[i] = countrySortedCases[i] - tempCountryOneSortedCases[i-1]
	}

	var tempCountryTwoSortedCases []float64
	for i := 0; i < len(countryTwoSortedCases); i++ {
		tempCountryTwoSortedCases = append(tempCountryTwoSortedCases, countryTwoSortedCases[i])

		if i == 0 {
			continue
		}
		countryTwoSortedCases[i] = countryTwoSortedCases[i] - tempCountryTwoSortedCases[i-1]
	}

	var countryOneStruct structs.CompareData
	var countryTwoStruct structs.CompareData

	countryOneStruct.Country = nameOne
	countryOneStruct.Data = countrySortedCases
	countryTwoStruct.Country = nameTwo
	countryTwoStruct.Data = countryTwoSortedCases

	compareStructs := structs.Compare{CountryOne: countryOneStruct, CountryTwo: countryTwoStruct}
	return compareStructs, nil
}

func GetCountryData(countryName string, countries []structs.CountryCurve) (structs.MainCurveData, error) {
	country, err := GetCountryBP(countryName, countries)
	if err != nil {
		return structs.MainCurveData{}, err
	}

	deaths := make([]float64, 0)
	cases := make([]float64, 0)
	recovered := make([]float64, 0)

	for _, vv := range country.Timeline.Deaths.(map[string]interface{}) {
		deaths = append(deaths, vv.(float64))
	}

	for _, vv := range country.Timeline.Cases.(map[string]interface{}) {
		cases = append(cases, vv.(float64))
	}

	for _, v := range country.Timeline.Recovered.(map[string]interface{}) {
		recovered = append(recovered, v.(float64))
	}

	sort.Float64s(deaths)
	sort.Float64s(cases)
	sort.Float64s(recovered)

	var tempCountryOneSortedDeath []float64
	deathsPerDay := make([]float64, 0)

	for i := 0; i < len(deaths); i++ {
		tempCountryOneSortedDeath = append(tempCountryOneSortedDeath, deaths[i])
		if i == 0 {
			continue
		}
		deathsPerDay = append(deathsPerDay, (deaths[i] - tempCountryOneSortedDeath[i-1]))
	}

	var tempCountryOneSortedCases []float64
	casesPerDay := make([]float64, 0)
	for i := 0; i < len(cases); i++ {
		tempCountryOneSortedCases = append(tempCountryOneSortedCases, cases[i])
		if i == 0 {
			continue
		}
		casesPerDay = append(casesPerDay, (cases[i] - tempCountryOneSortedCases[i-1]))
	}

	var tempCountryOneSortedRecovered []float64
	recoveredPerDay := make([]float64, 0)
	for i := 0; i < len(deaths); i++ {
		tempCountryOneSortedRecovered = append(tempCountryOneSortedRecovered, recovered[i])
		if i == 0 {
			continue
		}

		recoveredPerDay = append(recoveredPerDay, (recovered[i] - tempCountryOneSortedRecovered[i-1]))
	}

	return structs.MainCurveData{deaths, deathsPerDay, cases, casesPerDay, recovered, recoveredPerDay}, nil

}
