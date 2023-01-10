package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"sync"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}

	wg := &sync.WaitGroup{}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}
		wg.Add(1)
		go readClient(conn, wg)
	}
}

func readClient(conn net.Conn, wg *sync.WaitGroup) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("error reading from client: ", err.Error())
			break
		}
		_, err = conn.Write([]byte("+PONG\r\n"))
	}
	wg.Done()
}
