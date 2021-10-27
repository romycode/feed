package main

import (
	"log"
	"time"

	"github.com/romycode/go-feeder/internal"
)

const addr = ":4000"
const maxConn = 5
const ttl = 60 * time.Second

func main() {
	if err := internal.NewApp(addr, maxConn, ttl).Start(); err != nil {
		log.Fatalln(err)
	}
}
