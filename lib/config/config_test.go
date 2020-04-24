package pconf

import (
	"fmt"
	"testing"
)

//GetAppConfig reads a spefic file and return the json format of it
//@return ServerConfig struct json format of the config file
func TestGetAppConfig(t *testing.T) {
	cp := &conPath{
		path: fmt.Sprintf("../../config/covid.%s.json", "docker"),
	}
	fmt.Println(cp/)
	got := GetAppConfig()
	fmt.Println(got)
}
