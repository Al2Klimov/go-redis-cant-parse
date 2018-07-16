package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"io/ioutil"
	"os"
	"time"
)

func main() {
	if err := mkEvents(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func mkEvents() error {
	payloadRaw, errRA := ioutil.ReadFile("/PUBLISH.json")
	if errRA != nil {
		return errRA
	}

	payload := string(payloadRaw)
	rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})

	defer rdb.Close()

	for {
		time.Sleep(10 * time.Second)

		if _, errPublish := rdb.Publish("icinga:stats", payload).Result(); errPublish != nil {
			return errPublish
		}

		fmt.Println("Published to icinga:stats")
	}
}
