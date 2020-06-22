package analytics

import (
	"testing"

	mcountry "github.com/junkd0g/covid/lib/model/country"
)

type countryDataAnalytics struct{}

var countryDataAnalyticsMock func() ([]mcountry.CountryCurve, error)

func (u countryDataAnalytics) getAllCountries() ([]mcountry.CountryCurve, error) {
	return countryDataAnalyticsMock()
}

func TestMostCasesDeathsNearPast(t *testing.T) {
	countryData = countryDataAnalytics{}

	countryDataAnalyticsMock = func() ([]mcountry.CountryCurve, error) {
		var articleArr []mcountry.CountryCurve
		xTimeline := mcountry.TimelineStruct{
			Cases: map[string]interface{}{
				"1/22/20": 15.0,
				"1/23/20": 2.0,
				"1/24/20": 3.0,
				"1/25/20": 4.0,
				"1/26/20": 5.0,
				"1/27/20": 400.0,
				"1/28/20": 500.0,
				"1/29/20": 601.0,
			},
			Deaths: map[string]interface{}{
				"1/22/20": 1.0,
				"1/23/20": 25.0,
				"1/24/20": 3.0,
				"1/25/20": 44.0,
				"1/26/20": 5.0,
				"1/27/20": 400.0,
				"1/28/20": 500.0,
				"1/29/20": 3.0,
			},
			Recovered: map[string]interface{}{
				"1/22/20": 13.0,
				"1/23/20": 23.0,
				"1/24/20": 3.0,
				"1/25/20": 42.0,
				"1/26/20": 5.0,
				"1/27/20": 400.0,
				"1/28/20": 500.0,
				"1/29/20": 601.0,
			},
		}
		x := mcountry.CountryCurve{
			Country:  "xCountry",
			Province: "",
			Timeline: xTimeline,
		}

		yTimeline := mcountry.TimelineStruct{
			Cases: map[string]interface{}{
				"1/22/20": 1.0,
				"1/23/20": 200.0,
				"1/24/20": 3.0,
				"1/25/20": 4.0,
				"1/26/20": 5.0,
				"1/27/20": 400.0,
				"1/28/20": 500.0,
				"1/29/20": 601.0,
			},
			Deaths: map[string]interface{}{
				"1/22/20": 1.0,
				"1/23/20": 2.0,
				"1/24/20": 30.0,
				"1/25/20": 4.0,
				"1/26/20": 5.0,
				"1/27/20": 400.0,
				"1/28/20": 500.0,
				"1/29/20": 601.0,
			},
			Recovered: map[string]interface{}{
				"1/22/20": 10.0,
				"1/23/20": 2.0,
				"1/24/20": 3.0,
				"1/25/20": 4.0,
				"1/26/20": 5.0,
				"1/27/20": 400.0,
				"1/28/20": 500.0,
				"1/29/20": 601.0,
			},
		}
		y := mcountry.CountryCurve{
			Country:  "yCountry",
			Province: "",
			Timeline: yTimeline,
		}

		zTimeline := mcountry.TimelineStruct{
			Cases: map[string]interface{}{
				"1/22/20": 1.0,
				"1/23/20": 243.0,
				"1/24/20": 3.0,
				"1/25/20": 453.0,
				"1/26/20": 5.0,
				"1/27/20": 400.0,
				"1/28/20": 500.0,
				"1/29/20": 601.0,
			},
			Deaths: map[string]interface{}{
				"1/22/20": 1.0,
				"1/23/20": 2.0,
				"1/24/20": 366.0,
				"1/25/20": 4.0,
				"1/26/20": 566.0,
				"1/27/20": 400.0,
				"1/28/20": 500.0,
				"1/29/20": 601.0,
			},
			Recovered: map[string]interface{}{
				"1/22/20": 15.0,
				"1/23/20": 2.0,
				"1/24/20": 3.0,
				"1/25/20": 45.0,
				"1/26/20": 5.0,
				"1/27/20": 400.0,
				"1/28/20": 500.0,
				"1/29/20": 601.0,
			},
		}
		z := mcountry.CountryCurve{
			Country:  "zCountry",
			Province: "",
			Timeline: zTimeline,
		}

		wTimeline := mcountry.TimelineStruct{
			Cases: map[string]interface{}{
				"1/22/20": 1.0,
				"1/23/20": 2.0,
				"1/24/20": 34.0,
				"1/25/20": 4.0,
				"1/26/20": 5.0,
				"1/27/20": 400.0,
				"1/28/20": 500.0,
				"1/29/20": 601.0,
			},
			Deaths: map[string]interface{}{
				"1/22/20": 1.0,
				"1/23/20": 20.0,
				"1/24/20": 3.0,
				"1/25/20": 4.0,
				"1/26/20": 5.0,
				"1/27/20": 400.0,
				"1/28/20": 5040.0,
				"1/29/20": 601.0,
			},
			Recovered: map[string]interface{}{
				"1/22/20": 1.0,
				"1/23/20": 2.0,
				"1/24/20": 3.0,
				"1/25/20": 4.0,
				"1/26/20": 5.0,
				"1/27/20": 440.0,
				"1/28/20": 500.0,
				"1/29/20": 601.0,
			},
		}
		w := mcountry.CountryCurve{
			Country:  "wCountry",
			Province: "",
			Timeline: wTimeline,
		}
		articleArr = append(articleArr, x)
		articleArr = append(articleArr, y)
		articleArr = append(articleArr, z)
		articleArr = append(articleArr, w)

		return articleArr, nil
	}

	daysAmmount := 4
	mcdnpOneDay, mcdnpOneDayError := MostCasesDeathsNearPast(daysAmmount)
	if mcdnpOneDayError != nil {
		t.Fatal(mcdnpOneDayError)
	}

	if mcdnpOneDay.MostCases.Country != "xCountry" {
		t.Fatalf("Country with most cases looks wrong")
	}

	if len(mcdnpOneDay.MostCases.Data) != daysAmmount {
		t.Fatalf("Wrong ammount of data in cases")
	}

	if mcdnpOneDay.MostDeaths.Country != "wCountry" {
		t.Fatalf("Country with most cases looks wrong")
	}

	if len(mcdnpOneDay.MostDeaths.Data) != daysAmmount {
		t.Fatalf("Wrong ammount of data in deaths")
	}

	if calculateTotalAmmount(mcdnpOneDay.MostDeaths.Data) < calculateTotalAmmount(mcdnpOneDay.SecondDeaths.Data) {
		t.Fatalf("Wrong sum of data in deaths from most to second")
	}

	if calculateTotalAmmount(mcdnpOneDay.SecondDeaths.Data) < calculateTotalAmmount(mcdnpOneDay.ThirdDeaths.Data) {
		t.Fatalf("Wrong sum of data in deaths from second to third")
	}

}

func calculateTotalAmmount(arr []float64) float64 {
	total := 0.0
	for _, v := range arr {
		total = total + v
	}
	return total
}
