package curve

import (
	"testing"

	mcountry "github.com/junkd0g/covid/lib/model/country"
)

type requestDataMock struct{}

var requestDataMockFunc func() ([]mcountry.CountryCurve, error)

func (u requestDataMock) requestHistoryData() ([]mcountry.CountryCurve, error) {
	return requestDataMockFunc()
}

type requestCacheDataMock struct{}

var requestCacheDataMockFunc func() ([]mcountry.CountryCurve, error)

func (u requestCacheDataMock) getCacheData() ([]mcountry.CountryCurve, error) {
	return requestCacheDataMockFunc()
}

var setCacheDataMockFunc func(ctn []mcountry.CountryCurve) error

func (u requestCacheDataMock) setCacheData(ctn []mcountry.CountryCurve) error {
	return setCacheDataMockFunc(ctn)
}

func franceMonkData() []mcountry.CountryCurve {
	france1 := mcountry.CountryCurve{
		Country:  "France",
		Province: "SomeRandom",
		Timeline: mcountry.TimelineStruct{Cases: 44, Deaths: 43, Recovered: 44},
	}

	france2 := mcountry.CountryCurve{
		Country:  "France",
		Province: "",
		Timeline: mcountry.TimelineStruct{Cases: 440, Deaths: 430, Recovered: 440},
	}

	france3 := mcountry.CountryCurve{
		Country:  "France",
		Province: "SomeRandom",
		Timeline: mcountry.TimelineStruct{Cases: 4, Deaths: 4, Recovered: 4},
	}

	franceArray := []mcountry.CountryCurve{france1, france2, france3}
	return franceArray
}

func ukMonkData() []mcountry.CountryCurve {
	UK1 := mcountry.CountryCurve{
		Country:  "UK",
		Province: "SomeRandom",
		Timeline: mcountry.TimelineStruct{Cases: 44, Deaths: 43, Recovered: 44},
	}

	UK2 := mcountry.CountryCurve{
		Country:  "UK",
		Province: "SomeRandom",
		Timeline: mcountry.TimelineStruct{Cases: 440, Deaths: 430, Recovered: 440},
	}

	UK3 := mcountry.CountryCurve{
		Country:  "UK",
		Province: "",
		Timeline: mcountry.TimelineStruct{Cases: 4, Deaths: 4, Recovered: 4},
	}

	UKArray := []mcountry.CountryCurve{UK1, UK2, UK3}
	return UKArray
}

func TestGetCountryBP(t *testing.T) {

	el, err := GetCountryBP("France", franceMonkData())
	if err != nil {
		t.Fatal(err)
	}

	if el.Country != "France" {
		t.Fatalf("Wrong country name %s", el.Country)
	}

	if el.Timeline.Cases != 440 {
		t.Fatalf("Wrong ammout of cases %d", el.Timeline.Cases)
	}

	el2, err2 := GetCountryBP("UK", ukMonkData())
	if err2 != nil {
		t.Fatal(err2)
	}

	if el2.Country != "UK" {
		t.Fatalf("Wrong country name %s", el2.Country)
	}

	if el2.Timeline.Cases != 4 {
		t.Fatalf("Wrong ammout of cases %d", el2.Timeline.Cases)
	}
}

func TestGetAllCountries(t *testing.T) {
	reqCacheOB = requestCacheDataMock{}
	reqDataOB = requestDataMock{}

	setCacheDataMockFunc = func(ctn []mcountry.CountryCurve) error {
		return nil
	}

	requestCacheDataMockFunc = func() ([]mcountry.CountryCurve, error) {
		return []mcountry.CountryCurve{}, nil
	}

	requestCacheDataMockFunc = func() ([]mcountry.CountryCurve, error) {
		return ukMonkData(), nil
	}

	withCacheData, err := GetAllCountries()
	if err != nil {
		t.Fatal(err)
	}

	if withCacheData[0].Country != "UK" {
		t.Fatalf("Wrong country name %s", withCacheData[0].Country)
	}

	requestCacheDataMockFunc = func() ([]mcountry.CountryCurve, error) {

		return franceMonkData(), nil
	}

	withNoCacheData, err := GetAllCountries()
	if err != nil {
		t.Fatal(err)
	}

	if withNoCacheData[0].Country != "France" {
		t.Fatalf("Wrong country name %s", withNoCacheData[0].Country)
	}
}

func TestCompareDeathsCountries(t *testing.T) {
	setCacheDataMockFunc = func(ctn []mcountry.CountryCurve) error {
		return nil
	}

	requestCacheDataMockFunc = func() ([]mcountry.CountryCurve, error) {
		return []mcountry.CountryCurve{}, nil
	}

	requestCacheDataMockFunc = func() ([]mcountry.CountryCurve, error) {
		return multipleCountriesMock(), nil
	}

	compareDeathsData, err := CompareDeathsCountries("Greece", "Italy")
	if err != nil {
		t.Fatal(err)
	}

	if compareDeathsData.CountryOne.Country != "Greece" {
		t.Fatalf("Wrong country name %s", compareDeathsData.CountryOne.Country)
	}

	if compareDeathsData.CountryTwo.Country != "Italy" {
		t.Fatalf("Wrong country name %s", compareDeathsData.CountryTwo.Country)
	}

	mockCountryOneData := []float64{0, 0, 0, 1, 1, 5, 23, 75, 86, 92, 111, 112, 343}
	if !equal(compareDeathsData.CountryOne.Data, mockCountryOneData) {
		t.Fatalf("Wrong data %v", compareDeathsData.CountryOne.Data)
	}

	mockCountryTwoData := []float64{0, 0, 0, 1, 1, 5, 23, 75, 86, 92, 1211, 1312, 3430}
	if !equal(compareDeathsData.CountryTwo.Data, mockCountryTwoData) {
		t.Fatalf("Wrong data %v", compareDeathsData.CountryTwo.Data)
	}

	compareDeathsData2, err2 := CompareDeathsCountries("Italy", "UK")
	if err2 != nil {
		t.Fatal(err2)
	}

	if compareDeathsData2.CountryOne.Country != "Italy" {
		t.Fatalf("Wrong country name %s", compareDeathsData2.CountryOne.Country)
	}

	if compareDeathsData2.CountryTwo.Country != "UK" {
		t.Fatalf("Wrong country name %s", compareDeathsData.CountryTwo.Country)
	}

	mock2CountryTwoData := []float64{0, 0, 0, 1, 1, 5, 23, 75, 86, 92, 1211, 1312, 3430}
	if !equal(compareDeathsData.CountryTwo.Data, mock2CountryTwoData) {
		t.Fatalf("Wrong data %v", compareDeathsData.CountryTwo.Data)
	}

}

func equal(a, b []float64) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func multipleCountriesMock() []mcountry.CountryCurve {
	UK := mcountry.CountryCurve{
		Country:  "UK",
		Province: "",
		Timeline: mcountry.TimelineStruct{
			Deaths: map[string]interface{}{
				"1/22/20": 0.0,
				"1/23/20": 0.0,
				"1/24/20": 0.0,
				"3/19/20": 2232.0,
				"2/26/20": 1.0,
				"2/27/20": 1.0,
				"3/10/20": 5.0,
				"3/15/20": 16.0,
				"3/16/20": 21.0,
				"3/17/20": 22.0,
				"3/18/20": 22.0,
				"3/24/20": 74.0,
				"3/25/20": 84.0,
				"3/26/20": 94.0,
				"3/27/20": 110.0,
				"3/28/20": 110.0,
			},
			Cases: map[string]interface{}{
				"1/22/20": 0.0,
				"1/23/20": 0.0,
				"1/24/20": 0.0,
				"3/19/20": 2232.0,
				"2/26/20": 1.0,
				"2/27/20": 1.0,
				"3/10/20": 5.0,
				"3/15/20": 16.0,
				"3/16/20": 21.0,
				"3/17/20": 22.0,
				"3/18/20": 22.0,
				"3/24/20": 74.0,
				"3/25/20": 84.0,
				"3/26/20": 94.0,
				"3/27/20": 110.0,
				"3/28/20": 110.0,
			},
			Recovered: map[string]interface{}{
				"1/22/20": 0.0,
				"1/23/20": 0.0,
				"1/24/20": 0.0,
				"3/19/20": 2232.0,
				"2/26/20": 1.0,
				"2/27/20": 1.0,
				"3/10/20": 5.0,
				"3/15/20": 16.0,
				"3/16/20": 21.0,
				"3/17/20": 22.0,
				"3/18/20": 22.0,
				"3/24/20": 74.0,
				"3/25/20": 84.0,
				"3/26/20": 94.0,
				"3/27/20": 110.0,
				"3/28/20": 110.0,
			}},
	}

	Greece := mcountry.CountryCurve{
		Country:  "Greece",
		Province: "",
		Timeline: mcountry.TimelineStruct{
			Deaths: map[string]interface{}{
				"1/22/20": 0.0,
				"1/23/20": 0.0,
				"1/24/20": 0.0,
				"3/19/20": 343.0,
				"2/26/20": 1.0,
				"2/27/20": 1.0,
				"3/10/20": 5.0,
				"3/15/20": 23.0,
				"3/24/20": 75.0,
				"3/25/20": 86.0,
				"3/26/20": 92.0,
				"3/27/20": 111.0,
				"3/28/20": 112.0,
			},
			Cases: map[string]interface{}{
				"1/22/20": 0.0,
				"1/23/20": 0.0,
				"1/24/20": 0.0,
				"3/19/20": 2232.0,
				"2/26/20": 1.0,
				"2/27/20": 1.0,
				"3/10/20": 5.0,
				"3/15/20": 16.0,
				"3/16/20": 21.0,
				"3/17/20": 22.0,
				"3/18/20": 22.0,
				"3/24/20": 74.0,
				"3/25/20": 84.0,
				"3/26/20": 94.0,
				"3/27/20": 110.0,
				"3/28/20": 113.0,
			},
			Recovered: map[string]interface{}{
				"1/22/20": 0.0,
				"1/23/20": 0.0,
				"1/24/20": 0.0,
				"3/19/20": 2232.0,
				"2/26/20": 1.0,
				"2/27/20": 1.0,
				"3/10/20": 5.0,
				"3/15/20": 16.0,
				"3/16/20": 21.0,
				"3/17/20": 22.0,
				"3/18/20": 22.0,
				"3/24/20": 74.0,
				"3/25/20": 84.0,
				"3/26/20": 94.0,
				"3/27/20": 110.0,
				"3/28/20": 113.0,
			}},
	}

	Italy := mcountry.CountryCurve{
		Country:  "Italy",
		Province: "",
		Timeline: mcountry.TimelineStruct{
			Deaths: map[string]interface{}{
				"1/22/20": 0.0,
				"1/23/20": 0.0,
				"1/24/20": 0.0,
				"3/19/20": 3430.0,
				"2/26/20": 1.0,
				"2/27/20": 1.0,
				"3/10/20": 5.0,
				"3/15/20": 23.0,
				"3/24/20": 75.0,
				"3/25/20": 86.0,
				"3/26/20": 92.0,
				"3/27/20": 1211.0,
				"3/28/20": 1312.0,
			},
			Cases: map[string]interface{}{
				"1/22/20": 0.0,
				"1/23/20": 0.0,
				"1/24/20": 0.0,
				"3/19/20": 2232.0,
				"2/26/20": 1.0,
				"2/27/20": 1.0,
				"3/10/20": 5.0,
				"3/15/20": 16.0,
				"3/16/20": 21.0,
				"3/17/20": 22.0,
				"3/18/20": 22.0,
				"3/24/20": 74.0,
				"3/25/20": 84.0,
				"3/26/20": 94.0,
				"3/27/20": 110.0,
				"3/28/20": 110.0,
			},
			Recovered: map[string]interface{}{
				"1/22/20": 0.0,
				"1/23/20": 0.0,
				"1/24/20": 0.0,
				"3/19/20": 2232.0,
				"2/26/20": 1.0,
				"2/27/20": 1.0,
				"3/10/20": 5.0,
				"3/15/20": 16.0,
				"3/16/20": 21.0,
				"3/17/20": 22.0,
				"3/18/20": 22.0,
				"3/24/20": 74.0,
				"3/25/20": 84.0,
				"3/26/20": 94.0,
				"3/27/20": 110.0,
				"3/28/20": 110.0,
			}},
	}

	multiArray := []mcountry.CountryCurve{UK, Greece, Italy}
	return multiArray
}

//CompareDeathsCountries
//CompareDeathsFromFirstDeathCountries
//ComparePerDayDeathsCountries
//CompareRecoveryCountries
//CompareCasesCountries
//ComparePerDayCasesCountries
//GetCountryData
