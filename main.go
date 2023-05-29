package main

import (
	"fmt"
	"net"
)

func main() {
	// Start the server
	startServer()
}

func startServer() {
	// Listen on TCP port 6849
	listener, err := net.Listen("tcp", ":6849")
	if err != nil {
		fmt.Println("Failed to start server:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on port 6849")

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}

		// Handle each connection in a separate goroutine
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	for {
		// Read data from the connection
		buffer := make([]byte, 1024)
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Failed to read data:", err)
			break
		}

		// Print the received data
    fmt.Println(buffer[:n])
		fmt.Println("Received:", string(buffer[:n]))
	}
}
