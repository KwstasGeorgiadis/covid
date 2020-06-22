package analytics

import (
	applogger "github.com/junkd0g/covid/lib/applogger"
	curve "github.com/junkd0g/covid/lib/curve"
	mcountry "github.com/junkd0g/covid/lib/model/country"
	mhotspot "github.com/junkd0g/covid/lib/model/hotspot"
)

var (
	countryData getCountryData
)

func init() {
	countryData = countryOB{}
}

type countryOB struct{}

type getCountryData interface {
	getAllCountries() ([]mcountry.CountryCurve, error)
}

func (r countryOB) getAllCountries() ([]mcountry.CountryCurve, error) {
	countries, err := curve.GetAllCountries()
	return countries, err
}

// MostCasesDeathsNearPast returns 3 countries with
// most case and most deaths in n ammount of days
func MostCasesDeathsNearPast(days int) (mhotspot.Hotspot, error) {
	countries, err := countryData.getAllCountries()
	if err != nil {
		applogger.Log("ERROR", "analytics", "MostCasesDeathsLastWeek", err.Error())
		return mhotspot.Hotspot{}, err
	}
	var infoData mhotspot.Hotspot

	for _, v := range countries {
		countryData, countryDataError := curve.GetCountryData(v.Country, countries)
		if countryDataError != nil {
			applogger.Log("ERROR", "analytics", "MostCasesDeathsLastWeek", countryDataError.Error())
			return mhotspot.Hotspot{}, countryDataError
		}

		lastDaysCases := getLastData(countryData.CasesPerDay, days)

		if compare(infoData.MostCases.Data, lastDaysCases) == 2 {
			infoData.ThirdCases.Data = infoData.SecondCases.Data
			infoData.ThirdCases.Country = infoData.SecondCases.Country

			infoData.SecondCases.Data = infoData.MostCases.Data
			infoData.SecondCases.Country = infoData.MostCases.Country

			infoData.MostCases.Data = lastDaysCases
			infoData.MostCases.Country = v.Country
		} else if compare(infoData.SecondCases.Data, lastDaysCases) == 2 {
			infoData.ThirdCases.Data = infoData.SecondCases.Data
			infoData.ThirdCases.Country = infoData.SecondCases.Country

			infoData.SecondCases.Data = lastDaysCases
			infoData.SecondCases.Country = v.Country
		} else if compare(infoData.ThirdCases.Data, lastDaysCases) == 2 {
			infoData.ThirdCases.Data = lastDaysCases
			infoData.ThirdCases.Country = v.Country
		}

		lastDaysDeaths := getLastData(countryData.DeathsPerDay, days)

		if compare(infoData.MostDeaths.Data, lastDaysDeaths) == 2 {
			infoData.ThirdDeaths.Data = infoData.SecondDeaths.Data
			infoData.ThirdDeaths.Country = infoData.SecondDeaths.Country

			infoData.SecondDeaths.Data = infoData.MostDeaths.Data
			infoData.SecondDeaths.Country = infoData.MostDeaths.Country

			infoData.MostDeaths.Data = lastDaysDeaths
			infoData.MostDeaths.Country = v.Country
		} else if compare(infoData.SecondDeaths.Data, lastDaysDeaths) == 2 {
			infoData.ThirdDeaths.Data = infoData.SecondDeaths.Data
			infoData.ThirdDeaths.Country = infoData.SecondDeaths.Country

			infoData.SecondDeaths.Data = lastDaysDeaths
			infoData.SecondDeaths.Country = v.Country
		} else if compare(infoData.ThirdDeaths.Data, lastDaysDeaths) == 2 {
			infoData.ThirdDeaths.Data = lastDaysDeaths
			infoData.ThirdDeaths.Country = v.Country
		}

	}

	return infoData, nil
}

// return n ammount of last elements in an array
func getLastData(data []float64, days int) []float64 {
	lastDays := make([]float64, 0)
	for i := days; i >= 1; i-- {
		lastDays = append(lastDays, data[len(data)-i])
	}
	return lastDays
}

func compare(x []float64, y []float64) int {
	var xTotal = 0.0
	var yTotal = 0.0

	if len(x) == 0 {
		return 2
	}

	for i := 0; i < len(x); i++ {
		xTotal = xTotal + x[i]
		yTotal = yTotal + y[i]
	}

	if xTotal < yTotal {
		return 2
	}

	return 3
}
