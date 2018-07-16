package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

func main() {
	if err := monitorEvents(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func monitorEvents() error {
	rdb := redis.NewClient(&redis.Options{Addr: "127.0.0.1:6379"})
	defer rdb.Close()

	ps := rdb.Subscribe("icinga:stats")
	defer ps.Close()

	for {
		msg, errRM := ps.ReceiveMessage()
		if errRM != nil {
			return errRM
		}

		switch msg.Channel {
		case "icinga:stats":
			fmt.Println("Received from icinga:stats")
		}
	}
}
