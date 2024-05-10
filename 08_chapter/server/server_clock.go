package main

import(
	"fmt"
	"net"
	"time"
)

func handleConnection(c net.Conn){
	defer c.Close()
	fmt.Fprintln(c, "Hello", "there!")
	for{
		_, err := fmt.Fprintln(c, time.Now())
		if err != nil{
			fmt.Println(err)
			return 
		}
		time.Sleep(2 * time.Second)
	}
}

func main(){
	
	//Create a listener
	listener, _ := net.Listen("tcp", "localhost:8000")

	for{
		//Block until we hear a connection
		connection, _ := listener.Accept()
		//Pass the net.Conn object to a function to handle
		go handleConnection(connection)
	}
	
}

