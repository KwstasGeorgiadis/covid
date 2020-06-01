package cworld

//TODO add fucking caching you piece of shit and add expiration time

import (
	"sort"

	applogger "github.com/junkd0g/covid/lib/applogger"
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
func requestHistoryData() (structs.WorldTimeline, error) {
	client := &http.Client{}
	requestURL := serverConf.API.URLWorldHistory

	req, reqErr := http.NewRequest("GET", requestURL, nil)
	if reqErr != nil {
		applogger.Log("ERROR", "cworld", "requestHistoryData", reqErr.Error())
		return structs.WorldTimeline{}, reqErr
	}

	res, resError := client.Do(req)
	if resError != nil {
		applogger.Log("ERROR", "cworld", "requestHistoryData", resError.Error())
		return structs.WorldTimeline{}, resError
	}
	defer res.Body.Close()

	b, errorReadAll := ioutil.ReadAll(res.Body)
	if errorReadAll != nil {
		applogger.Log("ERROR", "cworld", "requestHistoryData", errorReadAll.Error())
		return structs.WorldTimeline{}, errorReadAll
	}

	var timeline structs.TimelineStruct
	if errUnmarshal := json.Unmarshal(b, &timeline); errUnmarshal != nil {
		applogger.Log("ERROR", "cworld", "requestHistoryData", errUnmarshal.Error())
		return structs.WorldTimeline{}, errUnmarshal
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

	var worldTimeline structs.WorldTimeline

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

func GetaWorldHistory() (structs.WorldTimeline, error) {
	c, v := requestHistoryData()
	return c, v
}
