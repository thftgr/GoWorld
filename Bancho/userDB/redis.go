package userDB

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/pterm/pterm"
	"os"
)

var (
	KeySql  string
	KeyMail string
	Host    string
)

var Redis [16]*redis.Client

// "github.com/go-redis/redis/v8"
func RedisConnect() {
	for i := 0; i < 16; i++ {
		func(iz int) {
			Redis[iz] = redis.NewClient(&redis.Options{
				Addr:     os.Getenv("X_REDIS_ADDRESS"),
				Username: os.Getenv("X_REDIS_USERNAME"),
				Password: os.Getenv("X_REDIS_PASSWORD"),
				DB:       iz,
			})
		}(i)
	}
	ReloadConfig()
}

func ReloadConfig() {
	if err := Redis[0].Ping(context.TODO()).Err(); err != nil {
		panic(err)
	}
	if err := Redis[15].HGet(context.TODO(), "BANCHO:CONFIG:RELEASE", "RDB").Scan(&KeySql); err != nil {
		panic(err)
	}
	if err := Redis[15].HGet(context.TODO(), "BANCHO:CONFIG:RELEASE", "MAIL:VERIFY:REGISTER").Scan(&KeyMail); err != nil {
		panic(err)
	}
	if err := Redis[15].HGet(context.TODO(), "BANCHO:CONFIG:RELEASE", "HOST:DOMAIN").Scan(&Host); err != nil {
		panic(err)
	}

	pterm.Info.Println("CONNECTED REDIS.")
}
