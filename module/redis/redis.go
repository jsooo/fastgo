package redis

import (
	"fastgo/module/config"

	"github.com/garyburd/redigo/redis"
	"github.com/gosexy/to"
	"github.com/jsooo/log"
)

//var rs redis.Conn

func connect() (redis.Conn, error) {
	conf, err := config.Reader("database.conf")
	if err != nil {
		log.Debug("get redis config error: " + to.String(err))
	}
	host := conf.String("redis::host")
	port := conf.String("redis::port")
	password := conf.String("redis::password")
	rs, err := redis.Dial("tcp", host+":"+port, redis.DialPassword(password))
	if err != nil {
		log.Error("redis connect error: " + to.String(err))
	}

	return rs, err
}

func Push(key string, value string) (interface{}, error) {
	rs, err := connect()
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	return rs.Do("RPUSH", key, value)
}

func Pop(key string) (interface{}, error) {
	rs, err := connect()
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	return rs.Do("LPOP", key)
}

func Get(key string) (string, error) {
	rs, err := connect()
	if err != nil {
		return "", err
	}
	defer rs.Close()
	reply, err := rs.Do("GET", key)
	if reply == nil {
		return "", err
	} else {
		return string(reply.([]byte)), err
	}
}

func Set(key string, value string) (interface{}, error) {
	rs, err := connect()
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	return rs.Do("SET", key, value)
}

func Del(key string) (interface{}, error) {
	rs, err := connect()
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	return rs.Do("DEL", key)
}

func SetWithExistTime(key string, value string, existTime int) (interface{}, error) {
	if existTime == 0 {
		return Set(key, value)
	}

	rs, err := connect()
	if err != nil {
		return nil, err
	}
	defer rs.Close()
	return rs.Do("SET", key, value, "EX", existTime)
}
