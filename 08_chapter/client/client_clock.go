package main

import(
	"os"
	"net"
	"io"
)

func main(){
	//Connect to a server
	connection, _ := net.Dial("tcp", "localhost:8000")

	io.Copy(os.Stdout, connection)
}