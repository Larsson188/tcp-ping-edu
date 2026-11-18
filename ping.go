package main

import (
	"fmt"
	"net"
	"time"
)

func check(addr string) bool {
	c, err := net.DialTimeout("tcp", addr, 2*time.Second)
	if err != nil {
		return false
	}
	c.Close()
	return true
}

func main() {
	fmt.Println("example.com:80 ->", check("example.com:80"))
}
