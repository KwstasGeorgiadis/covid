package caching

import (
	"encoding/json"
	"fmt"

	structs "../structs"
	"github.com/gomodule/redigo/redis"
)

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

// set executes the redis SET command
func Set(c redis.Conn, countries structs.Countries) error {
	fmt.Println(countries)

	_, err := c.Do("SET", "total", countries)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

// get executes the redis GET command
func Get(c redis.Conn) (structs.Countries, error) {
	// Simple GET example with String helper

	s, err := redis.String(c.Do("GET", "total"))
	if err != nil {
		return structs.Countries{}, err
	}

	data := structs.Countries{}
	json.Unmarshal([]byte(s), &data)

	return data, nil
}
