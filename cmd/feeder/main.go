package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/romycode/go-feeder/cmd/feeder/bootstrap"
)

const addr = ":4000"
const maxConn = 5
const ttl = 60 * time.Second

func main() {
	file := createLogFile()
	defer file.Close()

	app, err := bootstrap.BuildApp(addr, maxConn, ttl, file)
	if err != nil {
		log.Fatalln(err)
	}

	if err := app.Start(); err != nil {
		log.Fatalln(err)
	}
}

func createLogFile() *os.File {
	filename := fmt.Sprintf("sku-%s.log", time.Now().Format("2006-02-01_15:03:04"))
	file, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("could not open output file: %s\n", err.Error())
	}

	return file
}
