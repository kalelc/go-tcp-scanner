package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/fatih/color"
)

var wg = &sync.WaitGroup{}
var timeout = time.Millisecond * 200

func main() {
	hostname := flag.String("hostname", "localhost", "hostname to scan ports")
	startPort := flag.Int("start-port", 1, "the start port to scan")
	endPort := flag.Int("end-port", 100, "the end port to scan")

	flag.Parse()

	for i := *startPort; i <= *endPort; i++ {
		wg.Add(1)
		go Scanner(*hostname, i)
	}
	wg.Wait()
}

func Scanner(url string, port int) {
	_, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", url, port), timeout)
	if err == nil {
		color.Green("%s port %d is open\n", url, port)
	} else {
		color.Red("%s", err.Error())
	}
	defer wg.Done()
}
