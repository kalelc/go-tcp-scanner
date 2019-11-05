package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	for port := 1; port <= 10000; port++ {
		_, err := net.Dial("tcp", "localhost:"+strconv.Itoa(port))
		if err == nil {
			fmt.Println("Port " + strconv.Itoa(port) + " is open")
		} else {
			fmt.Println(err)
		}

	}
}
