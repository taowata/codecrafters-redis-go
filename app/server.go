package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:6379")
	if err != nil {
		fmt.Println("Failed to bind to port 6379")
		os.Exit(1)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go readClient(conn)
	}
}

func readClient(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err == io.EOF {
		return
	}
	if err != nil {
		fmt.Println("error reading from client: ", err.Error())
		//os.Exit(1)
		return
	}
	_, err = conn.Write([]byte("+PONG\r\n"))

}
