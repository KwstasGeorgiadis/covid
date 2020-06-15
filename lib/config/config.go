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

//APIConfig contains the data for exernal API http calls
type APIConfig struct {
	URL             string `json:"url"`
	URLHistory      string `json:"url_historical"`
	URLWorldHistory string `json:"url_world_historical"`
	News            string `json:"news"`
	VaccineNews     string `json:"vaccine_news"`
	TreatmentNews   string `json:"treatment_news"`
	Continent       string `json:"continent"`
}

//ServerConfig contains the data for the server like port
type ServerConfig struct {
	Port string `json:"port"`
	Log  string `json:"log"`
}

//RedisConfig contains the data for the redis server
type RedisConfig struct {
	MaxIdle   int    `json:"MaxIdle"`
	MaxActive int    `json:"MaxActive"`
	URL       string `json:"url"`
}

var (
	configPath = os.Getenv("env19")
)

func init() {
	if len(configPath) == 0 {
		//configPath = "./config/covid.development.json"
	}
}

//GetAppConfig reads a spefic file and return the json format of it
//@return ServerConfig struct json format of the config file
func GetAppConfig() AppConf {

	jsonFile, openfileError := os.Open(configPath)
	if openfileError != nil {
		fmt.Println("Cannot open server config file, filename: " + configPath)
		os.Exit(2)
	}

	byteValue, readFileError := ioutil.ReadAll(jsonFile)
	if readFileError != nil {
		fmt.Println("Cannot read server config file, filename: " + configPath)
		os.Exit(2)
	}

	var sc AppConf
	json.Unmarshal(byteValue, &sc)
	return sc
}
