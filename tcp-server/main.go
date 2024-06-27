package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

const (
	TYPE = "tcp"
	HOST = "localhost"
	PORT = "8080"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil {

			if err == io.EOF {
				fmt.Println("Client closed the connection")
				return
			}
			fmt.Println("Error reading:", err.Error())
			return
		}

		fmt.Println("\nReceived: ", string(buf[:n]))

		_, err = conn.Write([]byte("Message received\n"))

		if err != nil {
			fmt.Println("Error Writing:", err.Error())
			os.Exit(1)
			return
		}
	}
}

func main() {
	listener, err := net.Listen(TYPE, fmt.Sprintf("%s:%s", HOST, PORT))

	if err != nil {
		fmt.Println("Error creating listener:", err.Error())
		os.Exit(1)
		return
	}

	defer listener.Close()
	fmt.Printf("Server is listening on %v", fmt.Sprintf("%s:%s", HOST, PORT))

	for {
		connx, err := listener.Accept()

		if err != nil {
			fmt.Println("Error accepting connection:", err.Error())
			os.Exit(1)
			return
		}

		go handleConnection(connx)
	}

}
