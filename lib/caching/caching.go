package caching

//TODO adding expiration time - need to be more dynamic
//TODO add expiration time on requestDat for redis

/*
	Caching the results of the external API request for covid-19 data
*/

import (
	"encoding/json"

	"github.com/gomodule/redigo/redis"
	pconf "github.com/junkd0g/covid/lib/config"
	mcontinent "github.com/junkd0g/covid/lib/model/continent"
	mcountry "github.com/junkd0g/covid/lib/model/country"
	mcsse "github.com/junkd0g/covid/lib/model/csse"
	mnews "github.com/junkd0g/covid/lib/model/news"
	mworld "github.com/junkd0g/covid/lib/model/world"
)

var (
	serverConf = pconf.GetAppConfig()
	RedisOB    redisOBInt
)

func init() {
	RedisOB = RedisST{}
}

type RedisST struct{}
type redisOBInt interface {
	NewPool() *redis.Pool
	SetCountriesData(c redis.Conn, countries mcountry.Countries) error
	GetCountriesData(c redis.Conn) (mcountry.Countries, error)
	SetCurveData(c redis.Conn, countries []mcountry.CountryCurve) error
	GetCurveData(c redis.Conn) ([]mcountry.CountryCurve, error)
	SetNewsData(c redis.Conn, newsType string, news mnews.ArticlesData) error
	GetNewsData(c redis.Conn, newsType string) (mnews.ArticlesData, bool, error)
	GetContinentData(c redis.Conn) (mcontinent.Response, bool, error)
	SetContinetData(c redis.Conn, ctn mcontinent.Response) error
	SetCSSEData(c redis.Conn, ctn []mcsse.ResponseCountry) error
	GetCSSEData(c redis.Conn) ([]mcsse.ResponseCountry, error)
	SetWorldData(c redis.Conn, ctn mworld.WorldTimeline) error
	GetWorldData(c redis.Conn) (mworld.WorldTimeline, bool, error)
}

//NewPool connects to redis
func (r RedisST) NewPool() *redis.Pool {
	return &redis.Pool{
		MaxIdle:   serverConf.Redis.MaxIdle,
		MaxActive: serverConf.Redis.MaxActive,

		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", serverConf.Redis.URL)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// SetCountriesData executes the redis SET command
// @param c redis.Conn redis connection
func (r RedisST) SetCountriesData(c redis.Conn, countries mcountry.Countries) error {
	out, _ := json.Marshal(countries)

	_, err := c.Do("SETEX", "total", 2500, string(out))
	if err != nil {
		return err
	}

	return nil
}

// GetCountriesData executes the redis GET command
func (r RedisST) GetCountriesData(c redis.Conn) (mcountry.Countries, error) {
	// Simple GET example with String helper

	s, err := redis.String(c.Do("GET", "total"))
	if err != nil {
		return mcountry.Countries{}, nil
	}

	bytStr := []byte(s)
	datsa := mcountry.Countries{}

	erra := json.Unmarshal(bytStr, &datsa)
	if erra != nil {
		return mcountry.Countries{}, nil
	}
	return datsa, nil
}

// SetCurveData executes the redis SET command
// @param c redis.Conn redis connection
func (r RedisST) SetCurveData(c redis.Conn, countries []mcountry.CountryCurve) error {
	vv, _ := json.Marshal(countries)
	_, err := c.Do("SETEX", "curve", 2500, vv)
	if err != nil {
		return err
	}

	return nil
}

// GetCurveData executes the redis GET command
func (r RedisST) GetCurveData(c redis.Conn) ([]mcountry.CountryCurve, error) {
	s, err := redis.String(c.Do("GET", "curve"))
	if err != nil {
		return []mcountry.CountryCurve{}, nil
	}

	var data []mcountry.CountryCurve
	json.Unmarshal([]byte(s), &data)

	return data, nil
}

// SetNewsData executes the redis SET command
// @param c redis.Conn redis connection
func (r RedisST) SetNewsData(c redis.Conn, newsType string, news mnews.ArticlesData) error {
	vv, _ := json.Marshal(news)
	_, err := c.Do("SETEX", newsType, 7200, vv)
	if err != nil {
		return err
	}

	return nil
}

// GetNewsData executes the redis GET command
func (r RedisST) GetNewsData(c redis.Conn, newsType string) (mnews.ArticlesData, bool, error) {
	s, err := redis.String(c.Do("GET", newsType))
	if err != nil {

		return mnews.ArticlesData{}, false, nil
	}

	var data mnews.ArticlesData
	json.Unmarshal([]byte(s), &data)

	return data, true, nil
}

// GetContinentData executes the redis GET command
func (r RedisST) GetContinentData(c redis.Conn) (mcontinent.Response, bool, error) {
	s, err := redis.String(c.Do("GET", "continent"))
	if err != nil {

		return mcontinent.Response{}, false, nil
	}

	var data mcontinent.Response
	json.Unmarshal([]byte(s), &data)

	return data, true, nil
}

// SetContinetData executes the redis SET command
// @param c redis.Conn redis connection
func (r RedisST) SetContinetData(c redis.Conn, ctn mcontinent.Response) error {
	out, _ := json.Marshal(ctn)

	_, err := c.Do("SETEX", "continent", 2500, string(out))
	if err != nil {
		return err
	}

	return nil
}

// GetWorldData executes the redis GET command
func (r RedisST) GetWorldData(c redis.Conn) (mworld.WorldTimeline, bool, error) {
	s, err := redis.String(c.Do("GET", "world"))
	if err != nil {
		return mworld.WorldTimeline{}, false, nil
	}

	var data mworld.WorldTimeline
	json.Unmarshal([]byte(s), &data)

	return data, true, nil
}

// SetWorldData executes the redis SET command
// @param c redis.Conn redis connection
func (r RedisST) SetWorldData(c redis.Conn, ctn mworld.WorldTimeline) error {
	out, _ := json.Marshal(ctn)

	_, err := c.Do("SETEX", "world", 2500, string(out))
	if err != nil {
		return err
	}

	return nil
}

// GetCSSEData executes the redis GET command
func (r RedisST) GetCSSEData(c redis.Conn) ([]mcsse.ResponseCountry, error) {
	s, err := redis.String(c.Do("GET", "csse"))
	if err != nil {
		return []mcsse.ResponseCountry{}, nil
	}

	var data []mcsse.ResponseCountry
	json.Unmarshal([]byte(s), &data)

	return data, nil
}

// SetCSSEData executes the redis SET command
// @param c redis.Conn redis connection
func (r RedisST) SetCSSEData(c redis.Conn, ctn []mcsse.ResponseCountry) error {
	out, _ := json.Marshal(ctn)

	_, err := c.Do("SETEX", "csse", 2500, string(out))
	if err != nil {
		return err
	}

	return nil
}
