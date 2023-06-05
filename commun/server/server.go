package server

import (
	"fmt"
	"net"
  "io"
  "log"

  "github.com/Gridmax/Hillock/utility/configload"
  "github.com/Gridmax/Hillock/commun/dcontroller"
//  "github.com/Gridmax/Hillock/utility/messages"
)

func StartMessage(listenport string, dbenable string){
  log.Println("- - - - - - - - - - - - - -")
  log.Println("Starting Hillock server")
  log.Println("- - - - - - - - - - - - - -")
  log.Println("Hillock is listening on", listenport)
  if dbenable != ""{
    log.Println("Data warehouse is set")
  }else{
    log.Println("Data warehouse is not set, data will not be stored")
  }
}
func Start(configFile string) {
	// Listen on TCP port 6849
  config, err := configload.LoadConfig(configFile)
  
  StartMessage(config.ServerPort, config.DatabaseType)
  listener, err := net.Listen("tcp", ":"+config.ServerPort)
	if err != nil {
		fmt.Println("Failed to start server:", err)
    return
	}
	defer listener.Close()

	//fmt.Println("Server listening on port 6849")
  log.Println("Hillock Server successfully started")
	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Failed to accept connection:", err)
			continue
		}

		// Handle each connection in a separate goroutine
	 	go handleConnection(conn, config.DatabaseType)
	}
}

func handleConnection(conn net.Conn, dflow string) {
	defer conn.Close()

	// Read data from the client
	buffer := make([]byte, 1024)
	disconnected := false // Flag to track disconnection

	for !disconnected {
		n, err := conn.Read(buffer)
		if err != nil {
			if err == io.EOF {
				// Client disconnected, set the flag and break the loop
				disconnected = true
				break
			}
			fmt.Printf("Failed to read data: %v\n", err)
			break
		}

		// Process the received data
		data := buffer[:n]
    addr := conn.RemoteAddr().(*net.TCPAddr)
    log.Println("Data received from", addr, ":", data )
		// fmt.Println("Received data:", string(data))
    if dflow != "" {
      dcontroller.DataFlow(dflow, string(data))
    }


    // For debugging on Message convertion
    //str, val, err := messages.ConvertToJSONAndKeyValue(string(buffer[:n])) 
    //fmt.Println(str, val, err)
	}

	if disconnected {
    addr := conn.RemoteAddr().(*net.TCPAddr)
    log.Println(addr,"is Disconnected")
		//fmt.Println("Client disconnected")
	}

	// Perform any cleanup or additional handling after the client has disconnected
	// ...
}

//func handleConnection(conn net.Conn) {
//	defer conn.Close()
//
//	for {
//		// Read data from the connection
//		buffer := make([]byte, 1024)
//    addr := conn.RemoteAddr().(*net.TCPAddr)
//    fmt.Println(addr.IP.String())
//		n, err := conn.Read(buffer)
//		if err != nil {
//			fmt.Println("Failed to read data:", err)
//			break
//		}

		// Print the received data
  //  fmt.Println(buffer[:n])
 //   fmt.Println(string(buffer[:n]))
   // str, val, err := messages.ConvertToJSONAndKeyValue(string(buffer[:n])) 
   // fmt.Println(str, val, err)
//	}
//}

