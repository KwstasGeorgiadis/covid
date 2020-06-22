package cworld

import (
	"testing"

	mworld "github.com/junkd0g/covid/lib/model/world"
)

type requestDataMock struct{}

var requestDataMockFunc func() (mworld.WorldTimeline, error)

func (u requestDataMock) requestHistoryData() (mworld.WorldTimeline, error) {
	return requestDataMockFunc()
}

type requestCacheDataMock struct{}

var requestCacheDataMockFunc func() (mworld.WorldTimeline, bool, error)

func (u requestCacheDataMock) getCacheData() (mworld.WorldTimeline, bool, error) {
	return requestCacheDataMockFunc()
}

func TestRegisterUser(t *testing.T) {
	reqCacheOB = requestCacheDataMock{}
	reqDataOB = requestDataMock{}

	requestDataMockFunc = func() (mworld.WorldTimeline, error) {
		return mworld.WorldTimeline{
			Cases: []float64{
				1,
				10,
				100,
				1000,
				10000,
				30000,
				34054,
				45432,
			},
			Deaths: []float64{
				1,
				10,
				100,
				1000,
				10000,
				30000,
				34054,
				45432,
			},
			Recovered: []float64{
				1,
				10,
				100,
				1000,
				10000,
				30000,
				34054,
				45432,
			},
			CasesDaily: []float64{
				1,
				9,
				99,
				999,
				9999,
				19999,
				4054,
				5432,
			},
			DeathsDaily: []float64{
				1,
				9,
				99,
				999,
				9999,
				19999,
				4054,
				5432,
			},
			RecoveredDaily: []float64{
				1,
				9,
				99,
				999,
				9999,
				19999,
				4054,
				5432,
			},
		}, nil
	}

	requestCacheDataMockFunc = func() (mworld.WorldTimeline, bool, error) {
		return mworld.WorldTimeline{}, false, nil
	}

	withNoCashedData, err := GetaWorldHistory()
	if err != nil {
		t.Fatal(err)
	}

	if len(withNoCashedData.Cases.([]float64)) == 0 {
		t.Fatal("Getting cached data instead of requested data")
	}

	requestCacheDataMockFunc = func() (mworld.WorldTimeline, bool, error) {
		return mworld.WorldTimeline{
			Cases: []float64{
				1,
				10,
			},
			Deaths: []float64{
				1,
				10,
			},
			Recovered: []float64{
				1,
				10,
			},
			CasesDaily: []float64{
				1,
				9,
			},
			DeathsDaily: []float64{
				1,
				9,
			},
			RecoveredDaily: []float64{
				1,
				9,
			},
		}, true, nil
	}

	withEmptryButTrueCashedData, err := GetaWorldHistory()
	if err != nil {
		t.Fatal(err)
	}
	if len(withEmptryButTrueCashedData.Cases.([]float64)) != 2 {
		t.Fatal("Not getting cached data")
	}
}
