package conn

import (
	"log"
	"time"

	"github.com/garyburd/redigo/redis"
)

// RedisConfig for configuration connecion
type RedisConfig struct {
	Address  string `json:"address"`
	Password string `json:"password"`
}

// Redis variable :: References variable Redis to store redis connection
var Redis *redis.Pool

// InitRedis for initial connection to redis
func InitRedis(cfg RedisConfig) {
	log.Println("Initializing Redis")
	Redis = &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 10 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", cfg.Address)
			if err != nil || len(cfg.Password) < 1 {
				return c, err
			}

			c.Do("AUTH", cfg.Password)
			return c, err
		},
	}
	conn := Redis.Get()
	_, err := conn.Do("PING")
	if err != nil {
		log.Fatalln("Redis connection failed")
		return
	}
	log.Println("Redis successfully connected")

	InitVisitor()
}

// InitVisitor for initialize key visitor
func InitVisitor() error {
	key := "visitor"
	client := Redis.Get()
	defer client.Close()
	el, err := client.Do("GET", key)
	if err != nil {
		return err
	}
	if el == nil {
		_, err = redis.String(client.Do("SET", key, 1))
		if err != nil {
			return err
		}
	}
	return nil
}
