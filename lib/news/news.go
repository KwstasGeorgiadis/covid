package news

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	pconf "github.com/junkd0g/covid/lib/config"

	applogger "github.com/junkd0g/covid/lib/applogger"
	structs "github.com/junkd0g/covid/lib/structs"
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
func requestNewsData() (structs.ArticlesData, error) {
	url := "https://newsapi.org/v2/everything?q=COVID19&sortBy=publishedAt&apiKey=e37f283ddec24c3fb9dcf26eb59601e9&pageSize=100&page=1%0A"

	client := &http.Client{}
	req, reqError := http.NewRequest("GET", url, nil)

	if reqError != nil {
		applogger.Log("ERROR", "news", "requestNewsData", reqError.Error())
		return structs.ArticlesData{}, reqError
	}

	res, resError := client.Do(req)
	if resError != nil {
		applogger.Log("ERROR", "news", "requestNewsData", resError.Error())
		return structs.ArticlesData{}, resError

	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		applogger.Log("ERROR", "news", "requestNewsData", err.Error())
		return structs.ArticlesData{}, err

	}

	var reponseNews structs.ReponseNews

	unmarshallError := json.Unmarshal(body, &reponseNews)
	if unmarshallError != nil {
		applogger.Log("ERROR", "news", "requestNewsData", unmarshallError.Error())
	}

	keys := make([]structs.Article, 0)

	for _, v := range reponseNews.Articles {
		var article structs.Article
		article.Title = v.Title
		article.Description = v.Description
		article.URL = v.URL
		article.URLToImage = v.URLToImage
		article.PublishedAt = v.PublishedAt
		article.Content = v.Content
		keys = append(keys, article)
	}

	return structs.ArticlesData{keys}, nil

}

// GetNews returns an array of articles for covid-19
// It returns structs.ArticlesData and any write error encountered.
func GetNews() (structs.ArticlesData, error) {
	dashboard, err := requestNewsData()
	return dashboard, err
}
