package pconf

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	enviroment = "testing"
	cp = conPath{fmt.Sprintf("testFiles/covid.%s.json", enviroment)}
}

//GetAppConfig reads a spefic file and return the json format of it
//@return ServerConfig struct json format of the config file
func TestGetAppConfig(t *testing.T) {
	assert := assert.New(t)

	a := AppConf{Server: ServerConfig{Port: ":9080",
		Log: "/var/log/covid/app.ndjson"}, API: APIConfig{URL: "https://corona.lmao.ninja/v2/countries",
		URLHistory: "https://corona.lmao.ninja/v2/historical/?lastdays=all"},
		Redis: RedisConfig{URI: "redis://172.18.0.1", Namespace: "resque:",
			Concurrency: 2, Queues: []string{"myqueue", "delimited", "queues"},
			Name: "", Port: ":6379"}}

	b := GetAppConfig()

	assert.Equal(a, b, "Config looks good")

}
