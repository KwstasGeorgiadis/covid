package cworld

//TODO add fucking caching you piece of shit and add expiration time

import (
	"sort"

	applogger "github.com/junkd0g/covid/lib/applogger"
	"github.com/junkd0g/covid/lib/caching"
	pconf "github.com/junkd0g/covid/lib/config"
	mcountry "github.com/junkd0g/covid/lib/model/country"
	mworld "github.com/junkd0g/covid/lib/model/world"

	"encoding/json"
	"io/ioutil"
	"net/http"
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
	requestHistoryData() (mworld.WorldTimeline, error)
}

type requestCacheData struct{}
type requestCache interface {
	getCacheData() (mworld.WorldTimeline, bool, error)
	setCacheData(ctn mworld.WorldTimeline) error
}

func (r requestCacheData) setCacheData(ctn mworld.WorldTimeline) error {
	err := redis.SetWorldData(ctn)
	return err
}

func (r requestCacheData) getCacheData() (mworld.WorldTimeline, bool, error) {
	cachedData, exist, cacheGetError := redis.GetWorldData()
	return cachedData, exist, cacheGetError
}

// requestData does an HTTP GET request to the third party API that
// contains covid-9 stats ' history (per day from 22/01/2020)
// It returns []mcountry.Country and any write error encountered.
func (r requestData) requestHistoryData() (mworld.WorldTimeline, error) {
	client := &http.Client{}
	requestURL := serverConf.API.URLWorldHistory

	req, reqErr := http.NewRequest("GET", requestURL, nil)
	if reqErr != nil {
		applogger.Log("ERROR", "cworld", "requestHistoryData", reqErr.Error())
		return mworld.WorldTimeline{}, reqErr
	}

	res, resError := client.Do(req)
	if resError != nil {
		applogger.Log("ERROR", "cworld", "requestHistoryData", resError.Error())
		return mworld.WorldTimeline{}, resError
	}
	defer res.Body.Close()

	b, errorReadAll := ioutil.ReadAll(res.Body)
	if errorReadAll != nil {
		applogger.Log("ERROR", "cworld", "requestHistoryData", errorReadAll.Error())
		return mworld.WorldTimeline{}, errorReadAll
	}

	var timeline mcountry.TimelineStruct
	if errUnmarshal := json.Unmarshal(b, &timeline); errUnmarshal != nil {
		applogger.Log("ERROR", "cworld", "requestHistoryData", errUnmarshal.Error())
		return mworld.WorldTimeline{}, errUnmarshal
	}

	deaths := make([]float64, 0)
	cases := make([]float64, 0)
	recovered := make([]float64, 0)

	for _, v := range timeline.Deaths.(map[string]interface{}) {
		deaths = append(deaths, v.(float64))
	}

	for _, v := range timeline.Cases.(map[string]interface{}) {
		cases = append(cases, v.(float64))
	}

	for _, v := range timeline.Recovered.(map[string]interface{}) {
		recovered = append(recovered, v.(float64))
	}

	sort.Float64s(deaths)
	sort.Float64s(cases)
	sort.Float64s(recovered)

	var worldTimeline mworld.WorldTimeline

	worldTimeline.Deaths = deaths
	worldTimeline.Cases = cases
	worldTimeline.Recovered = recovered

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

	worldTimeline.DeathsDaily = deathsPerDay
	worldTimeline.CasesDaily = casesPerDay
	worldTimeline.RecoveredDaily = recoveredPerDay

	return worldTimeline, nil
}

//GetaWorldHistory returns world history for covid-19
func GetaWorldHistory() (mworld.WorldTimeline, error) {
	cachedData, exist, cacheGetError := reqCacheOB.getCacheData()
	if cacheGetError != nil {
		applogger.Log("ERROR", "cworld", "GetaWorldHistory", cacheGetError.Error())
		return mworld.WorldTimeline{}, cacheGetError
	}

	if !exist {
		applogger.Log("INFO", "cworld", "GetaWorldHistory", "Request data instead of getting cached data")
		data, err := reqDataOB.requestHistoryData()
		if err != nil {
			applogger.Log("ERROR", "cworld", "GetaWorldHistory", err.Error())
			return mworld.WorldTimeline{}, err
		}
		reqCacheOB.setCacheData(data)
		return data, nil
	}
	return cachedData, nil
}
