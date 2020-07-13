package continent

import (
	"net/http"

	applogger "github.com/junkd0g/covid/lib/applogger"
	caching "github.com/junkd0g/covid/lib/caching"
	pconf "github.com/junkd0g/covid/lib/config"
	mcontinent "github.com/junkd0g/covid/lib/model/continent"
)

var (
	serverConf      pconf.AppConf
	reqDataOB       requestAPI
	reqCacheOB      requestCache
	redis           caching.RedisST
	continentObject mcontinent.ContinentOB
)

func init() {
	serverConf = pconf.GetAppConfig()
	reqDataOB = requestData{}
	reqCacheOB = requestCacheData{}
}

type requestData struct{}
type requestAPI interface {
	requestContinentData() (mcontinent.Response, error)
}

type requestCacheData struct{}
type requestCache interface {
	getCacheData() (mcontinent.Response, error)
	setCacheData(ctn mcontinent.Response) error
}

func (r requestCacheData) setCacheData(ctn mcontinent.Response) error {
	err := redis.SetContinetData(ctn)
	return err
}

//requestContinentData does a GET http request to serverConf.API.Continent value ( https://corona.lmao.ninja​/v2/continents )
func (r requestData) requestContinentData() (mcontinent.Response, error) {
	client := &http.Client{}
	requestURL := serverConf.API.Continent

	req, reqErr := http.NewRequest("GET", requestURL, nil)
	if reqErr != nil {
		applogger.Log("ERROR", "continent", "requestContinentData", reqErr.Error())
		return mcontinent.Response{}, reqErr
	}

	res, resError := client.Do(req)
	if resError != nil {
		applogger.Log("ERROR", "continent", "requestContinentData", resError.Error())
		return mcontinent.Response{}, resError
	}
	defer res.Body.Close()

	responseData, errUnmarshal := continentObject.UnmarshalContintent(res.Body)
	if errUnmarshal != nil {
		applogger.Log("ERROR", "continent", "requestContinentData", errUnmarshal.Error())
		return mcontinent.Response{}, errUnmarshal
	}

	return responseData, nil
}

func (r requestCacheData) getCacheData() (mcontinent.Response, error) {
	cachedData, _, cacheGetError := redis.GetContinentData()
	return cachedData, cacheGetError
}

// GetContinentData checks if continent data are on redis and return them
// else it request them using requestContinentData
func GetContinentData() (mcontinent.Response, error) {
	cachedData, cacheGetError := reqCacheOB.getCacheData()
	if cacheGetError != nil {
		applogger.Log("ERROR", "continent", "GetContinentData", cacheGetError.Error())
		return mcontinent.Response{}, cacheGetError
	}

	if len(cachedData) == 0 {
		applogger.Log("INFO", "continent", "GetContinentData", "Request data instead of getting cached data")
		data, err := reqDataOB.requestContinentData()
		if err != nil {
			applogger.Log("ERROR", "continent", "GetContinentData", err.Error())
			return mcontinent.Response{}, err
		}
		reqCacheOB.setCacheData(data)
		return data, nil
	}

	return cachedData, nil
}
