package main

import (
	"fmt"
	"net"
	"os"

)

func main (){

	fmt.Println("Logs appear here!!");
	
	l, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		fmt.Println("Failed to bind to port 8080")
		os.Exit(1)
	}

	_, err = l.Accept()
	if err != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}
}