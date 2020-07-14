package mcontinent

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestUnmarshalContintent(t *testing.T) {
	jsonFile, err := os.Open("../../../test/files/continent.json")
	if err != nil {
		t.Fatal(err)
	}

	defer jsonFile.Close()
	continentData, continentDataErr := Continent.UnmarshalContintent(jsonFile)
	if continentDataErr != nil {
		t.Fatal(continentDataErr)
	}

	if len(continentData) != 6 {
		t.Fatalf("Wrong length of array %d", len(continentData))
	}

	byteValue, readFileError := ioutil.ReadAll(jsonFile)
	if readFileError != nil {
		t.Fatal(readFileError)

	}

	var sc Response
	json.Unmarshal(byteValue, &sc)

	if reflect.DeepEqual(sc, continentData) {
		t.Fatalf("Responses are not equal %v with %v", sc, continentData)
	}

}
