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
			"port" : ":6660"
		},
		"API" : {
			"url" : "https://corona.lmao.ninja/countries?sort=country"
			"url_historical" : "https://corona.lmao.ninja/countires"
		},
		"redis" : {
			"namespace" 	: "resque:",
			"concurrency"	: 2,
			"uri" 			: "redis://localhost:6379/",
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
}

//RedisConfig contains the data for the redis server
type RedisConfig struct {
	URI         string   `json:"URI"`
	Namespace   string   `json:"namespace"`
	Concurrency int      `json:"concurrency"`
	Queues      []string `json:"queues"`
	Name        string   `json:"name"`
}

//GetAppConfig reads a spefic file and return the json format of it
//@param configLocation string configs location
//@return ServerConfig struct json format of the config file
func GetAppConfig(configLocation string) AppConf {
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
