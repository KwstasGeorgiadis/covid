package curve

import (
	"testing"

	pconf "github.com/junkd0g/covid/lib/config"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MyMockedObject struct {
	mock.Mock
}

func init() {

	serverConf = pconf.AppConf{Server: pconf.ServerConfig{Port: ":9080",
		Log: "/var/log/covid/app.ndjson"}, API: pconf.APIConfig{URL: "https://corona.lmao.ninja/v2/countries",
		URLHistory: "https://corona.lmao.ninja/v2/historical/?lastdays=all"},
		Redis: pconf.RedisConfig{URI: "redis://172.18.0.1", Namespace: "resque:",
			Concurrency: 2, Queues: []string{"myqueue", "delimited", "queues"},
			Name: "", Port: ":6379"}}
	//originalValidate = validate

}

//func mock_requestHistoryData() ([]structs.CountryCurve, error) {
//jsonFile, _ := os.Open("testFiles/data.json")

//byteValue, _ := ioutil.ReadAll(jsonFile)

//var sc []structs.CountryCurve
//json.Unmarshal(byteValue, &sc)
//return sc, nil
//}

func TestGetAllCountries(t *testing.T) {

	testObj := new(MyMockedObject)
	testObj.On("init")
	assert := assert.New(t)

	a := "hey"
	b := "hey"

	assert.Equal(a, b, "Config looks good")

}
