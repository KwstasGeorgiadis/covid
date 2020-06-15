package pconf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetAppConfig(t *testing.T) {
	assert := assert.New(t)

	a := AppConf{
		Server: ServerConfig{
			Port: ":9080",
			Log:  "/var/log/covid/app.ndjson",
		},
		API: APIConfig{
			URL:             "https://corona.lmao.ninja/v2/countries",
			URLHistory:      "https://corona.lmao.ninja/v2/historical/?lastdays=all",
			URLWorldHistory: "https://corona.lmao.ninja/v2/historical/all/?lastdays=all",
			News:            "http://news.google.com/news?q=covid-19&hl=en-US&sort=date&gl=US&num=100&output=rss",
			VaccineNews:     "http://news.google.com/news?q=covid-19_vaccine&hl=en-US&sort=date&gl=US&num=100&output=rss",
			TreatmentNews:   "http://news.google.com/news?q=covid-19_treatment&hl=en-US&sort=date&gl=US&num=100&output=rss",
			Continent:       "https://corona.lmao.ninja/v2/continents",
		},
		Redis: RedisConfig{
			URL:       "127.0.0.1:6379",
			MaxActive: 1200,
			MaxIdle:   80,
		},
	}

	b := GetAppConfig()

	assert.Equal(a, b, "Config looks good")

}
