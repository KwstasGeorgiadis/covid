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
func (r RedisST) SetCountriesData(countries mcountry.Countries) error {
	pool := r.NewPool()
	conn := pool.Get()
	defer conn.Close()
	out, _ := json.Marshal(countries)

	_, err := conn.Do("SETEX", "total", 2500, string(out))
	if err != nil {
		return err
	}

	return nil
}

// GetCountriesData executes the redis GET command
func (r RedisST) GetCountriesData() (mcountry.Countries, error) {
	pool := r.NewPool()
	conn := pool.Get()
	defer conn.Close()
	s, err := redis.String(conn.Do("GET", "total"))
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
func (r RedisST) SetCurveData(countries []mcountry.CountryCurve) error {
	pool := r.NewPool()
	conn := pool.Get()
	defer conn.Close()
	vv, _ := json.Marshal(countries)
	_, err := conn.Do("SETEX", "curve", 2500, vv)
	if err != nil {
		return err
	}

	return nil
}

// GetCurveData executes the redis GET command
func (r RedisST) GetCurveData() ([]mcountry.CountryCurve, error) {
	pool := r.NewPool()
	conn := pool.Get()
	defer conn.Close()
	s, err := redis.String(conn.Do("GET", "curve"))
	if err != nil {
		return []mcountry.CountryCurve{}, nil
	}

	var data []mcountry.CountryCurve
	json.Unmarshal([]byte(s), &data)

	return data, nil
}

// SetNewsData executes the redis SET command
func (r RedisST) SetNewsData(newsType string, news mnews.ArticlesData) error {
	pool := r.NewPool()
	conn := pool.Get()
	defer conn.Close()
	vv, _ := json.Marshal(news)
	_, err := conn.Do("SETEX", newsType, 7200, vv)
	if err != nil {
		return err
	}

	return nil
}

// GetNewsData executes the redis GET command
func (r RedisST) GetNewsData(newsType string) (mnews.ArticlesData, bool, error) {
	pool := r.NewPool()
	conn := pool.Get()
	defer conn.Close()
	s, err := redis.String(conn.Do("GET", newsType))
	if err != nil {

		return mnews.ArticlesData{}, false, nil
	}

	var data mnews.ArticlesData
	json.Unmarshal([]byte(s), &data)

	return data, true, nil
}

// GetContinentData executes the redis GET command
func (r RedisST) GetContinentData() (mcontinent.Response, bool, error) {
	pool := r.NewPool()
	conn := pool.Get()
	defer conn.Close()
	s, err := redis.String(conn.Do("GET", "continent"))
	if err != nil {

		return mcontinent.Response{}, false, nil
	}

	var data mcontinent.Response
	json.Unmarshal([]byte(s), &data)

	return data, true, nil
}

// SetContinetData executes the redis SET command
func (r RedisST) SetContinetData(ctn mcontinent.Response) error {

	pool := r.NewPool()
	conn := pool.Get()
	defer conn.Close()
	out, _ := json.Marshal(ctn)
	_, err := conn.Do("SETEX", "continent", 2500, string(out))
	if err != nil {
		return err
	}

	return nil
}

// GetWorldData executes the redis GET command
func (r RedisST) GetWorldData() (mworld.WorldTimeline, bool, error) {
	pool := r.NewPool()
	conn := pool.Get()
	defer conn.Close()
	s, err := redis.String(conn.Do("GET", "world"))
	if err != nil {
		return mworld.WorldTimeline{}, false, nil
	}

	var data mworld.WorldTimeline
	json.Unmarshal([]byte(s), &data)

	return data, true, nil
}

// SetWorldData executes the redis SET command
func (r RedisST) SetWorldData(ctn mworld.WorldTimeline) error {
	pool := r.NewPool()
	conn := pool.Get()
	defer conn.Close()
	out, _ := json.Marshal(ctn)
	_, err := conn.Do("SETEX", "world", 2500, string(out))
	if err != nil {
		return err
	}

	return nil
}

// GetCSSEData executes the redis GET command
func (r RedisST) GetCSSEData() ([]mcsse.ResponseCountry, error) {
	pool := r.NewPool()
	conn := pool.Get()
	defer conn.Close()
	s, err := redis.String(conn.Do("GET", "csse"))
	if err != nil {
		return []mcsse.ResponseCountry{}, nil
	}

	var data []mcsse.ResponseCountry
	json.Unmarshal([]byte(s), &data)

	return data, nil
}

// SetCSSEData executes the redis SET command
func (r RedisST) SetCSSEData(ctn []mcsse.ResponseCountry) error {
	pool := r.NewPool()
	conn := pool.Get()
	defer conn.Close()
	out, _ := json.Marshal(ctn)

	_, err := conn.Do("SETEX", "csse", 2500, string(out))
	if err != nil {
		return err
	}

	return nil
}
