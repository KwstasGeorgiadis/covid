package applogger

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

type logNDJOSN struct {
	PID        string    `json:"pid"`
	Level      string    `json:"level"`
	LogPackage string    `json:"package"`
	LogFunc    string    `json:"func"`
	Message    string    `json:"message"`
	DOB        time.Time `json:"name"`
}

var (
	generalLogger *log.Logger
	errorLogger   *log.Logger
)

func init() {
	absPath, err := filepath.Abs("/var/log/covid")
	if err != nil {
		fmt.Println("Error reading given path:", err)
	}

	generalLog, err := os.OpenFile(absPath+"/app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	generalLogger = log.New(generalLog, "", 0)
	errorLogger = log.New(generalLog, "", 0)
}

func Log(level string, logPackage string, logFunc string, message string) {

	s1, _ := time.Parse(time.RFC3339, "2018-12-12")
	x := logNDJOSN{PID: "sad", Level: level, LogPackage: logPackage, LogFunc: logFunc, Message: message, DOB: s1}
	res2B, _ := json.Marshal(x)
	generalLogger.Println(string(res2B))
}
