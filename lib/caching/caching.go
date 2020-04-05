package caching

//TODO adding expiration time - need to be more dynamic

/*
	Caching the results of the external API request for covid-19 data
*/

import (
	"encoding/json"

	structs "../structs"
	"github.com/gomodule/redigo/redis"
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
			c, err := redis.Dial("tcp", ":6379")
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

	_, err := c.Do("SET", key, countries)
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
		return structs.Countries{}, err
	}

	data := structs.Countries{}
	json.Unmarshal([]byte(s), &data)

	return data, nil
}
