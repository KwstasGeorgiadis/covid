package mcsse

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

func TestUnmarshalCSSE(t *testing.T) {
	jsonFile, err := os.Open("../../../test/files/mcsse.json")
	if err != nil {
		t.Fatal(err)
	}

	defer jsonFile.Close()
	csseData, csseDataErr := CSSE.UnmarshalCSSE(jsonFile)
	if csseDataErr != nil {
		t.Fatal(csseDataErr)
	}

	if len(csseData) != 729 {
		t.Fatalf("Wrong length of array %d", len(csseData))
	}

	byteValue, readFileError := ioutil.ReadAll(jsonFile)
	if readFileError != nil {
		t.Fatal(readFileError)

	}

	var sc []ResponseCountry
	json.Unmarshal(byteValue, &sc)

	if reflect.DeepEqual(sc, csseData) {
		t.Fatalf("Responses are not equal %v with %v", sc, csseData)
	}

}
