package caching

//TODO adding expiration time - need to be more dynamic
//TODO add expiration time on requestDat for redis

/*
	Caching the results of the external API request for covid-19 data
*/

import (
	"encoding/json"

	pconf "github.com/junkd0g/covid/lib/config"
	structs "github.com/junkd0g/covid/lib/structs"

	"github.com/gomodule/redigo/redis"
)

var (
	serverConf = pconf.GetAppConfig()
)

//NewPool() connects to redis
func NewPool() *redis.Pool {
	return &redis.Pool{
		// Maximum number of idle connections in the pool.
		MaxIdle: 80,
		// max number of connections
		MaxActive: 12000,
		// Dial is an application supplied function for creating and
		// configuring a connection.
		Dial: func() (redis.Conn, error) {
			//c, err := redis.Dial("tcp", serverConf.Redis.Port)
			c, err := redis.Dial("tcp", "127.0.0.1:6379")
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}
}

// Set executes the redis SET command
// @param c redis.Conn redis connection
func Set(c redis.Conn, countries structs.Countries, key string) error {
	out, _ := json.Marshal(countries)

	_, err := c.Do("SETEX", key, 900, string(out))
	if err != nil {
		return err
	}

	return nil
}

// Get executes the redis GET command
func Get(c redis.Conn, key string) (structs.Countries, error) {
	// Simple GET example with String helper

	s, err := redis.String(c.Do("GET", key))
	if err != nil {
		return structs.Countries{}, nil
	}

	bytStr := []byte(s)
	datsa := structs.Countries{}

	erra := json.Unmarshal(bytStr, &datsa)
	if erra != nil {
		return structs.Countries{}, nil
	}
	return datsa, nil
}

// SetCurveData executes the redis SET command
// @param c redis.Conn redis connection
func SetCurveData(c redis.Conn, countries []structs.CountryCurve) error {
	vv, _ := json.Marshal(countries)
	_, err := c.Do("SETEX", "curve", 900, vv)
	if err != nil {
		return err
	}

	return nil
}

// GetCurveData executes the redis GET command
func GetCurveData(c redis.Conn) ([]structs.CountryCurve, error) {
	s, err := redis.String(c.Do("GET", "curve"))
	if err != nil {
		return []structs.CountryCurve{}, nil
	}

	var data []structs.CountryCurve
	json.Unmarshal([]byte(s), &data)

	return data, nil
}

// SetNewsData executes the redis SET command
// @param c redis.Conn redis connection
func SetNewsData(c redis.Conn, newsType string, news structs.ArticlesData) error {
	vv, _ := json.Marshal(news)
	_, err := c.Do("SETEX", newsType, 7200, vv)
	if err != nil {
		return err
	}

	return nil
}

// GetNewsData executes the redis GET command
func GetNewsData(c redis.Conn, newsType string) (structs.ArticlesData, bool, error) {
	s, err := redis.String(c.Do("GET", newsType))
	if err != nil {

		return structs.ArticlesData{}, false, nil
	}

	var data structs.ArticlesData
	json.Unmarshal([]byte(s), &data)

	return data, true, nil
}
