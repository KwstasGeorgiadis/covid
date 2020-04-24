package pconf

/*
	return stuct of a json config file
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	Example of a config file

	{
		"server" : {
			"port" : ":6660",
			"log" : "/var/log/covid/app.ndjson"
		},
		"API" : {
			"url" : "https://corona.lmao.ninja/countries?sort=country"
			"url_historical" : "https://corona.lmao.ninja/countires"
		},
		"redis" : {
			"namespace" 	: "resque:",
			"concurrency"	: 2,
			"uri" 			: "redis://localhost",
			"port"			: ":6379",
			"queues" 		: ["myqueue","delimited","queues"]
		}
	}
*/

//AppConf contains all main structs
type AppConf struct {
	Server ServerConfig `json:"server"`
	API    APIConfig    `json:"API"`
	Redis  RedisConfig  `json:"redis"`
}

//APIConfig contains the data for exernal API
type APIConfig struct {
	URL        string `json:"url"`
	URLHistory string `json:"url_historical"`
}

//ServerConfig contains the data for the server like port
type ServerConfig struct {
	Port string `json:"port"`
	Log  string `json:"log"`
}

//RedisConfig contains the data for the redis server
type RedisConfig struct {
	URI         string   `json:"URI"`
	Namespace   string   `json:"namespace"`
	Concurrency int      `json:"concurrency"`
	Queues      []string `json:"queues"`
	Name        string   `json:"name"`
	Port        string   `json:"port"`
}

var (
	enviroment = os.Getenv("env19")
)

type conPath struct {
	path string
}

func init() {
	if len(enviroment) == 0 {
		enviroment = "development"
	}
	fmt.Println(fmt.Sprintf("Running in %s mode", enviroment))
}

//GetAppConfig reads a spefic file and return the json format of it
//@return ServerConfig struct json format of the config file
func GetAppConfig() AppConf {
	cp := conPath{fmt.Sprintf("./config/covid.%s.json", enviroment)}
	configLocation := cp.path

	jsonFile, openfileError := os.Open(configLocation)
	if openfileError != nil {
		fmt.Println("Cannot open server config file, filename: " + configLocation)
		os.Exit(2)
	}

	byteValue, readFileError := ioutil.ReadAll(jsonFile)
	if readFileError != nil {
		fmt.Println("Cannot read server config file, filename: " + configLocation)
		os.Exit(2)
	}

	var sc AppConf
	json.Unmarshal(byteValue, &sc)
	return sc
}
