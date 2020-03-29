package pconf

/*
	return stuct of a json config
*/

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

/*
	{
		"server" : {
			"port" : ":6660"
		},
		"redis" : {
			"namespace" 	: "resque:",
			"concurrency"	: 2,
			"uri" 			: "redis://localhost:6379/",
			"queues" 		: ["myqueue","delimited","queues"]
		}
	}
*/

type AppConf struct {
	Server ServerConfig  	`json:"server"`
	Redis  RedisConfig   	`json:"redis"`
}

type ServerConfig struct {
	Port string `json:"port"`
}

type RedisConfig struct {
	URI         string   `json:"URI"`
	Namespace   string   `json:"namespace"`
	Concurrency int      `json:"concurrency"`
	Queues      []string `json:"queues"`
	Name		string   `json:"name"`
}

/*
	reads a spefic file and return the json format of it

	@param configLocation string configs location
	@return ServerConfig struct json format of the config file
*/
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
