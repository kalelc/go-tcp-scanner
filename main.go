package main

import (
	"flag"
	"fmt"
	"net"
	"sync"
)

var wg = &sync.WaitGroup{}

func main() {
	hostname := flag.String("hostname", "localhost", "hostname to scan ports")
	startPort := flag.Int("start-port", 1, "the start port to scan")
	endPort := flag.Int("end-port", 100, "the end port to scan")

	for i := *startPort; i <= *endPort; i++ {
		wg.Add(1)
		go Scanner(*hostname, i)
	}
	wg.Wait()
}

func Scanner(url string, port int) {
	_, err := net.Dial("tcp", fmt.Sprintf("%s:%d", url, port))
	if err == nil {
		fmt.Printf("%s port %d is open\n", url, port)
	} else {
		fmt.Println(err)
	}
	defer wg.Done()
}
