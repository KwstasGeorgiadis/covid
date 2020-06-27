package csse

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	applogger "github.com/junkd0g/covid/lib/applogger"
	caching "github.com/junkd0g/covid/lib/caching"
	pconf "github.com/junkd0g/covid/lib/config"
	mcsse "github.com/junkd0g/covid/lib/model/csse"
)

var (
	serverConf pconf.AppConf
	reqDataOB  requestAPI
	reqCacheOB requestCache
	redis      caching.RedisST
)

func init() {
	serverConf = pconf.GetAppConfig()
	reqDataOB = requestData{}
	reqCacheOB = requestCacheData{}
}

type requestData struct{}

type requestAPI interface {
	requestCSSEData() ([]mcsse.ResponseCountry, error)
}

type requestCacheData struct{}
type requestCache interface {
	getCacheData() ([]mcsse.ResponseCountry, error)
	setCacheData(ctn []mcsse.ResponseCountry) error
}

//requestCSSEData request csse data from external api
func (r requestData) requestCSSEData() ([]mcsse.ResponseCountry, error) {
	client := &http.Client{}
	requestURL := serverConf.API.CSSE

	req, reqErr := http.NewRequest("GET", requestURL, nil)
	if reqErr != nil {
		applogger.Log("ERROR", "csse", "requestCSSEData", reqErr.Error())
		return []mcsse.ResponseCountry{}, reqErr
	}

	res, resError := client.Do(req)
	if resError != nil {
		applogger.Log("ERROR", "csse", "requestCSSEData", resError.Error())
		return []mcsse.ResponseCountry{}, resError
	}
	defer res.Body.Close()

	b, errorReadAll := ioutil.ReadAll(res.Body)
	if errorReadAll != nil {
		applogger.Log("ERROR", "csse", "requestCSSEData", errorReadAll.Error())
		return []mcsse.ResponseCountry{}, errorReadAll
	}

	var responseData []mcsse.ResponseCountry
	if errUnmarshal := json.Unmarshal(b, &responseData); errUnmarshal != nil {
		applogger.Log("ERROR", "csse", "requestCSSEData", errUnmarshal.Error())
		return []mcsse.ResponseCountry{}, errUnmarshal
	}

	return responseData, nil
}

// getCacheData get data from redis for csse key
func (r requestCacheData) getCacheData() ([]mcsse.ResponseCountry, error) {
	pool := redis.NewPool()
	conn := pool.Get()
	defer conn.Close()
	cachedData, cacheGetError := redis.GetCSSEData(conn)

	return cachedData, cacheGetError
}

func (r requestCacheData) setCacheData(ctn []mcsse.ResponseCountry) error {
	pool := redis.NewPool()
	conn := pool.Get()
	defer conn.Close()
	err := redis.SetCSSEData(conn, ctn)

	return err
}

// GetCSSEData checks if continent data are on redis and return them
// else it request them using requestContinentData
func GetCSSEData() (mcsse.CSSEResponse, error) {
	data, dataErr := reqCacheOB.getCacheData()
	if dataErr != nil {
		applogger.Log("ERROR", "csse", "GetCSSEData", dataErr.Error())
		return mcsse.CSSEResponse{}, dataErr
	}

	if len(data) == 0 {
		applogger.Log("INFO", "csse", "GetCSSEData", "Request data instead of getting cached data")
		data, dataErr = reqDataOB.requestCSSEData()
		if dataErr != nil {
			applogger.Log("ERROR", "csse", "GetCSSEData", dataErr.Error())
			return mcsse.CSSEResponse{}, dataErr
		}
		reqCacheOB.setCacheData(data)
	}
	var countries []mcsse.CSEECountryResponse

	for _, k := range data {
		countries = insertProvince(countries, k)
	}

	return mcsse.CSSEResponse{Data: countries}, nil
}

// GetCSSECountryData returns csse data for a specific country
func GetCSSECountryData(country string) (mcsse.CSEECountryResponse, error) {
	country = getCountriesName(country)

	countriesData, err := GetCSSEData()
	if err != nil {
		applogger.Log("ERROR", "csse", "GetCSSECountryData", err.Error())
		return mcsse.CSEECountryResponse{}, err
	}

	for _, v := range countriesData.Data {
		if v.Country == country {
			return v, nil
		}
	}
	return mcsse.CSEECountryResponse{}, nil
}

// insertProvince normalise response data from having a one to
// one releastion between country and provision to one country
// having multiple provision (basically pu them into an array)
func insertProvince(arr []mcsse.CSEECountryResponse, element mcsse.ResponseCountry) []mcsse.CSEECountryResponse {
	for i := 0; i < len(arr); i++ {
		if arr[i].Country == element.Country {
			var cp mcsse.CSEEProvision
			cp.Cases = element.Stats.Confirmed
			cp.Deaths = element.Stats.Deaths
			cp.Recovered = element.Stats.Recovered
			cp.County = fmt.Sprintf("%v", element.County)
			cp.Province = element.Province
			arr[i].Data = append(arr[i].Data, cp)

			return arr
		}
	}

	var newC mcsse.CSEECountryResponse
	newC.Country = element.Country

	var cp mcsse.CSEEProvision
	cp.Cases = element.Stats.Confirmed
	cp.Deaths = element.Stats.Deaths
	cp.Recovered = element.Stats.Recovered
	cp.County = fmt.Sprintf("%v", element.County)
	cp.Province = element.Province

	var cpArr []mcsse.CSEEProvision
	cpArr = append(cpArr, cp)
	newC.Data = cpArr
	arr = append(arr, newC)
	return arr
}

//getCountriesName converts some country's name to the values we are using across the API
func getCountriesName(country string) string {

	countriesMap := make(map[string]string)
	countriesMap["Bosnia"] = "Bosnia and Herzegovina"
	countriesMap["Côte d'Ivoire"] = "Cote d'Ivoire"
	countriesMap["Holy See (Vatican City State)"] = "Holy See"
	countriesMap["S.Korea"] = "Korea, South"
	countriesMap["Lao People's Democratic Republic"] = "Laos"
	countriesMap["Macedonia"] = "North Macedonia"
	countriesMap["USA"] = "US"
	countriesMap["UAE"] = "United Arab Emirates"
	countriesMap["UK"] = "United Kingdom"

	if countriesMap[country] != "" {
		return countriesMap[country]
	}
	return country
}
