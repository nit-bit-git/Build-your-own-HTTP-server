package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {

	fmt.Println("Logs appear here!!")

	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("Failed to bind to port 8080")
		os.Exit(1)
	}
  for{
	conn, err := l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
	go handleConnection(conn)
  }
	
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	fmt.Println("New connection from:", conn.RemoteAddr())

	//read request

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil && err != io.EOF {
		fmt.Println("Error reading:", err)
		return
	}
	fmt.Printf("Received request:\n%s\n", string(buffer[:n]))

	//send response

	response := "HTTP/1.1 200 OK\r\nContent-Length: 12\r\n\r\nHelloWORLD!!"
	_, err = conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error sending response")
		return

	}

	fmt.Println("Response sent!")

}
