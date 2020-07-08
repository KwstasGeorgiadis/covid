package mockredis

import (
	"github.com/gomodule/redigo/redis"
	mcontinent "github.com/junkd0g/covid/lib/model/continent"
	mcountry "github.com/junkd0g/covid/lib/model/country"
	mcsse "github.com/junkd0g/covid/lib/model/csse"
	mnews "github.com/junkd0g/covid/lib/model/news"
	mworld "github.com/junkd0g/covid/lib/model/world"
)

var (
	RedisOB redisOBInt
)

func init() {
	RedisOB = MockRedisST{}
}

type MockRedisST struct{}

type redisOBInt interface {
	NewPool() *redis.Pool
	SetCountriesData(countries mcountry.Countries) error
	GetCountriesData() (mcountry.Countries, error)
	SetCurveData(countries []mcountry.CountryCurve) error
	GetCurveData() ([]mcountry.CountryCurve, error)
	SetNewsData(newsType string, news mnews.ArticlesData) error
	GetNewsData(newsType string) (mnews.ArticlesData, bool, error)
	GetContinentData() (mcontinent.Response, bool, error)
	SetContinetData(ctn mcontinent.Response) error
	SetCSSEData(ctn []mcsse.ResponseCountry) error
	GetCSSEData() ([]mcsse.ResponseCountry, error)
	SetWorldData(ctn mworld.WorldTimeline) error
	GetWorldData() (mworld.WorldTimeline, bool, error)
}

var mockGetWorldData func() (mworld.WorldTimeline, bool, error)

func (u MockRedisST) GetWorldData() (mworld.WorldTimeline, bool, error) {
	return mockGetWorldData()
}

var mockSetWorldData func(ctn mworld.WorldTimeline) error

func (u MockRedisST) SetWorldData(ctn mworld.WorldTimeline) error {
	return mockSetWorldData(ctn)
}

var mockGetCSSEData func() ([]mcsse.ResponseCountry, error)

func (u MockRedisST) GetCSSEData() ([]mcsse.ResponseCountry, error) {
	return mockGetCSSEData()
}

var mockSetCSSEData func(ctn []mcsse.ResponseCountry) error

func (u MockRedisST) SetCSSEData(ctn []mcsse.ResponseCountry) error {
	return mockSetCSSEData(ctn)
}

var mockSetContinetData func(ctn mcontinent.Response) error

func (u MockRedisST) SetContinetData(ctn mcontinent.Response) error {
	return mockSetContinetData(ctn)
}

var mockGetContinentData func() (mcontinent.Response, bool, error)

func (u MockRedisST) GetContinentData() (mcontinent.Response, bool, error) {
	return mockGetContinentData()
}

var mockGetNewsData func(newsType string) (mnews.ArticlesData, bool, error)

func (u MockRedisST) GetNewsData(newsType string) (mnews.ArticlesData, bool, error) {
	return mockGetNewsData(newsType)
}

var mockSetNewsData func(newsType string, news mnews.ArticlesData) error

func (u MockRedisST) SetNewsData(newsType string, news mnews.ArticlesData) error {
	return mockSetNewsData(newsType, news)
}

var mockGetCurveData func() ([]mcountry.CountryCurve, error)

func (u MockRedisST) GetCurveData() ([]mcountry.CountryCurve, error) {
	return mockGetCurveData()
}

var mockSetCurveData func(countries []mcountry.CountryCurve) error

func (u MockRedisST) SetCurveData(countries []mcountry.CountryCurve) error {
	return mockSetCurveData(countries)
}

var mockGetCountriesData func() (mcountry.Countries, error)

func (u MockRedisST) GetCountriesData() (mcountry.Countries, error) {
	return mockGetCountriesData()
}

var mockSetCountriesData func(countries mcountry.Countries) error

func (u MockRedisST) SetCountriesData(countries mcountry.Countries) error {
	return mockSetCountriesData(countries)
}

var mockNewPool func() *redis.Pool

func (u MockRedisST) NewPool() *redis.Pool {
	return mockNewPool()
}
