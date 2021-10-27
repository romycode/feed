package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/romycode/go-feeder/internal"
)

func main() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go sendData(&wg)
	}

	wg.Wait()
}

func sendData(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		conn, err := getTCPConn()

		_, err = fmt.Fprint(conn, internal.NewSKU().Value())
		if err != nil {
			log.Fatalln("could not write to socket:", err)
		}

		_ = conn.Close()
	}

	wg.Done()
}

func getTCPConn() (*net.TCPConn, error) {
	addr, err := net.ResolveTCPAddr("tcp", ":4000")
	if err != nil {
		log.Fatalln(err)
	}

	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Fatalln(err)
	}

	return conn, err
}
