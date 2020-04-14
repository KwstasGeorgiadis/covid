package applogger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	pconf "../config"

	"github.com/gofrs/uuid"
)

type logNDJOSN struct {
	PID        string    `json:"pid"`
	Level      string    `json:"level"`
	LogPackage string    `json:"package"`
	LogFunc    string    `json:"func"`
	Message    string    `json:"message"`
	DOB        time.Time `json:"time"`
}

type logNDJOSNHTTP struct {
	PID        string    `json:"pid"`
	Level      string    `json:"level"`
	LogPackage string    `json:"package"`
	LogFunc    string    `json:"func"`
	Message    string    `json:"message"`
	DOB        time.Time `json:"time"`
	Code       int       `json:"code"`
	Duration   float64   `json:"duration"`
}

var (
	generalLogger *log.Logger
	errorLogger   *log.Logger
	serverConf    = pconf.GetAppConfig("./config/covid.json")
)

func init() {
	generalLog, err := os.OpenFile(serverConf.Server.Log, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	generalLogger = log.New(generalLog, "", 0)
	errorLogger = log.New(generalLog, "", 0)
}

func Log(level string, logPackage string, logFunc string, message string) {

	s1 := time.Now()
	u := uuid.Must(uuid.NewV4())

	x := logNDJOSN{PID: u.String(), Level: level, LogPackage: logPackage, LogFunc: logFunc, Message: message, DOB: s1}
	res2B, _ := json.Marshal(x)
	generalLogger.Println(string(res2B))
}

func LogHTTP(level string, logPackage string, logFunc string, message string, code int, duration float64) {

	s1 := time.Now()
	u := uuid.Must(uuid.NewV4())

	x := logNDJOSNHTTP{PID: u.String(), Level: level, LogPackage: logPackage, LogFunc: logFunc, Message: message, DOB: s1, Code: code, Duration: duration}
	res2B, _ := json.Marshal(x)
	generalLogger.Println(string(res2B))
}
