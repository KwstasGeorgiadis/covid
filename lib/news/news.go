package news

import (
	"encoding/xml"
	"io/ioutil"
	"net/http"

	caching "github.com/junkd0g/covid/lib/caching"
	pconf "github.com/junkd0g/covid/lib/config"

	applogger "github.com/junkd0g/covid/lib/applogger"
	mnews "github.com/junkd0g/covid/lib/model/news"
)

var (
	serverConf pconf.AppConf
)

func init() {
	serverConf = pconf.GetAppConfig()
}

// requestNewsData does an HTTP GET request to the third party API that
// contains covid-9 news article
// It returns structs.ArticlesData and any write error encountered.
func requestNewsData(url string) (mnews.ArticlesData, error) {

	client := &http.Client{}
	req, reqError := http.NewRequest("GET", url, nil)

	if reqError != nil {
		applogger.Log("ERROR", "news", "requestNewsData", reqError.Error())
		return mnews.ArticlesData{}, reqError
	}

	res, resError := client.Do(req)
	if resError != nil {
		applogger.Log("ERROR", "news", "requestNewsData", resError.Error())
		return mnews.ArticlesData{}, resError

	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		applogger.Log("ERROR", "news", "requestNewsData", err.Error())
		return mnews.ArticlesData{}, err

	}

	var reponseNews mnews.ReponseNews

	unmarshallError := xml.Unmarshal(body, &reponseNews)
	if unmarshallError != nil {
		applogger.Log("ERROR", "news", "requestNewsData", unmarshallError.Error())
	}

	keys := make([]mnews.Article, 0)

	for _, v := range reponseNews.Channel.Item {

		var article mnews.Article
		article.Title = v.Title
		article.Description = v.Description
		article.URL = v.Link
		article.PublishedAt = v.PubDate
		article.Source = v.Source.Text
		article.SourceURL = v.Source.URL
		keys = append(keys, article)
	}

	return mnews.ArticlesData{keys}, nil

}

// GetNews returns an array of articles for covid-19
// It returns structs.ArticlesData and any write error encountered.
func GetNews() (mnews.ArticlesData, error) {
	pool := caching.NewPool()
	conn := pool.Get()
	defer conn.Close()

	cachedData, exist, cacheGetError := caching.GetNewsData(conn, "general")
	if cacheGetError != nil {
		applogger.Log("ERROR", "curve", "GetNews", cacheGetError.Error())
		return mnews.ArticlesData{}, cacheGetError
	}

	if !exist {
		applogger.Log("INFO", "stats", "GetNews", "Request data instead of getting cached data")
		data, err := requestNewsData(serverConf.API.News)
		if err != nil {
			applogger.Log("ERROR", "curve", "GetNews", err.Error())
			return mnews.ArticlesData{}, err
		}
		caching.SetNewsData(conn, "general", data)
		return data, nil
	}

	return cachedData, nil
}

// GetVaccineNews returns an array of articles for covid-19
// It returns structs.ArticlesData and any write error encountered.
func GetVaccineNews() (mnews.ArticlesData, error) {

	pool := caching.NewPool()
	conn := pool.Get()
	defer conn.Close()

	cachedData, exist, cacheGetError := caching.GetNewsData(conn, "vaccine")
	if cacheGetError != nil {
		applogger.Log("ERROR", "curve", "GetVaccineNews", cacheGetError.Error())
		return mnews.ArticlesData{}, cacheGetError
	}

	if !exist {
		applogger.Log("INFO", "stats", "GetVaccineNews", "Request data instead of getting cached data")
		data, err := requestNewsData(serverConf.API.VaccineNews)
		if err != nil {
			applogger.Log("ERROR", "curve", "GetVaccineNews", err.Error())
			return mnews.ArticlesData{}, err
		}
		caching.SetNewsData(conn, "vaccine", data)
		return data, nil
	}

	return cachedData, nil
}

// GetTreatmentNews returns an array of articles for covid-19
// It returns structs.ArticlesData and any write error encountered.
func GetTreatmentNews() (mnews.ArticlesData, error) {

	pool := caching.NewPool()
	conn := pool.Get()
	defer conn.Close()

	cachedData, exist, cacheGetError := caching.GetNewsData(conn, "treatment")
	if cacheGetError != nil {
		applogger.Log("ERROR", "curve", "GetTreatmentNews", cacheGetError.Error())
		return mnews.ArticlesData{}, cacheGetError
	}

	if !exist {
		applogger.Log("INFO", "stats", "GetTreatmentNews", "Request data instead of getting cached data")
		data, err := requestNewsData(serverConf.API.TreatmentNews)
		if err != nil {
			applogger.Log("ERROR", "curve", "GetTreatmentNews", err.Error())
			return mnews.ArticlesData{}, err
		}
		caching.SetNewsData(conn, "treatment", data)
		return data, nil
	}

	return cachedData, nil
}
