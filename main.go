package main

import (
	"fmt"
	"net"
)

func main() {
	//Listen for Incoming Connections: Use the net.Listen() function to specify the network and
	//address to listen on. In this case, we’re listening on a specific IP address and port number.
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {

		}
	}(listener)

	fmt.Println("Listening on :8080")

	for {
		//Accept Incoming Connections: Once you’re listening, you can accept incoming client
		//connections using Listener.Accept(). This function blocks until a client connects.
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		//Handle Client Connections: After accepting a client connection, you’ll typically
		//spawn a new goroutine to handle that connection. This allows your server to handle
		//multiple clients concurrently.
		go handleClient(conn)
		//Read and Write Data: In the goroutine handling the client, you can use the connobject
		//to read data from and write data to the client.
	}
}

func handleClient(conn net.Conn) {
	//Close the Connection: When the communication is done, you should close the connection
	//using conn.Close().
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)
}
