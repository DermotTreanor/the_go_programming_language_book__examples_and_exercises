package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"time"
)

func handleConnection(c net.Conn){
	defer c.Close()
	fmt.Fprintln(c, "Hello", "there!")
	for{
		_, err := fmt.Fprint(c, time.Now().Format("15:04:05\r"))
		if err != nil{
			fmt.Println(err)
			return 
		}
		time.Sleep(1 * time.Second)
	}
}


var port *int = flag.Int("port", 8000, "selects the port number to bind the server to")

func main(){

	flag.Parse()
	address := fmt.Sprintf("localhost:%d", *port)
	fmt.Println(*port)
	//Create a listener
	listener, err := net.Listen("tcp", address)
	if err !=nil{
		fmt.Fprintf(os.Stderr, "Error while making listener:%T\n%v\n\nClosing server program...\n", err, err)
		return
	}
	
	for{
		//Block until we hear a connection
		connection, _ := listener.Accept()
		//Pass the net.Conn object to a function to handle
		go handleConnection(connection)
	}	
}

