package csse

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	mcsse "github.com/junkd0g/covid/lib/model/csse"
)

func TestInsertProvince(t *testing.T) {
	jsonFile, _ := os.Open("t1.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var sc []mcsse.ResponseCountry
	json.Unmarshal(byteValue, &sc)
	var countries []mcsse.CSEECountryResponse

	for _, k := range sc {
		countries = insertProvince(countries, k)
	}

	for _, v := range countries {
		if v.Country == "Italy" {
			if len(v.Data) != 1 {
				t.Fatal("Ammount of data for Italy ")
			}

			if v.Data[0].Province != "Abruzzo" {
				t.Fatal("Italy's province data was wrong")
			}
		}

		if v.Country == "Russia" {
			if len(v.Data) != 3 {
				t.Fatal("Ammount of data for Russia ")
			}
		}
	}
}

type requestDataMock struct{}

var requestDataMockFunc func() ([]mcsse.ResponseCountry, error)

func (u requestDataMock) requestCSSEData() ([]mcsse.ResponseCountry, error) {
	return requestDataMockFunc()
}

type requestCacheDataMock struct{}

var requestCacheDataMockFunc func() ([]mcsse.ResponseCountry, error)

func (u requestCacheDataMock) getCacheData() ([]mcsse.ResponseCountry, error) {
	return requestCacheDataMockFunc()
}

func TestCSSEData(t *testing.T) {
	reqCacheOB = requestCacheDataMock{}
	reqDataOB = requestDataMock{}

	jsonFile, _ := os.Open("t2re.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var sc []mcsse.ResponseCountry
	json.Unmarshal(byteValue, &sc)

	requestDataMockFunc = func() ([]mcsse.ResponseCountry, error) {
		return sc, nil
	}

	requestCacheDataMockFunc = func() ([]mcsse.ResponseCountry, error) {
		return []mcsse.ResponseCountry{}, nil
	}

	withNoCashedData, err := GetCSSEData()
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range withNoCashedData.Data {
		if v.Country == "Italy" {
			if len(v.Data) != 1 {
				t.Fatal("Ammount of data for Italy ")
			}

			if v.Data[0].Province != "Abruzzo" {
				t.Fatal("Italy's province data was wrong")
			}
		}

		if v.Country == "Russia" {
			if len(v.Data) != 3 {
				t.Fatal("Ammount of data for Russia ")
			}
		}
	}

	jsonFileCa, _ := os.Open("t3ca.json")
	byteValueCa, _ := ioutil.ReadAll(jsonFileCa)

	var scCA []mcsse.ResponseCountry
	json.Unmarshal(byteValueCa, &scCA)
	withCashedData, errWithCacheData := GetCSSEData()
	if errWithCacheData != nil {
		t.Fatal(errWithCacheData)
	}

	for _, v := range withCashedData.Data {
		if v.Country == "US" {
			if len(v.Data) != 1 {
				t.Fatal("Ammount of data for US ")
			}

			if v.Data[0].Province != "Abruzzo" {
				t.Fatal("US's province data was wrong")
			}
		}

		if v.Country == "Greece" {
			if len(v.Data) != 3 {
				t.Fatal("Ammount of data for Russia ")
			}
		}
	}
}

func TestCSSECountryData(t *testing.T) {
	reqCacheOB = requestCacheDataMock{}
	reqDataOB = requestDataMock{}

	jsonFile, _ := os.Open("t2re.json")
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var sc []mcsse.ResponseCountry
	json.Unmarshal(byteValue, &sc)

	requestDataMockFunc = func() ([]mcsse.ResponseCountry, error) {
		return sc, nil
	}

	requestCacheDataMockFunc = func() ([]mcsse.ResponseCountry, error) {
		return []mcsse.ResponseCountry{}, nil
	}

	withNoCashedData, err := GetCSSECountryData("Russia")
	if err != nil {
		t.Fatal(err)
	}

	if withNoCashedData.Country != "Russia" {
		t.Fatal("Wrong country data ")
	}
	if len(withNoCashedData.Data) != 3 {
		t.Fatal("Ammount of data for Russia ")
	}
}

func TestGetCountriesName(t *testing.T) {
	v1 := getCountriesName("USA")
	expecetedV1 := "US"
	if v1 != expecetedV1 {
		t.Fatalf("Wrong value for %s which should be coverted to %s", v1, expecetedV1)
	}
	v2 := getCountriesName("Greece")
	if v2 != "Greece" {
		t.Fatalf("Wrong value for %s which should not be coverted and stay %s", v2, v2)
	}
}
