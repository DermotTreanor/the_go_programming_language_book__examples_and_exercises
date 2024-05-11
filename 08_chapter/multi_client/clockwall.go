package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"errors"
)

func make_connection(add string){
	//Connect to a server
	connection, _ := net.Dial("tcp", add)
	//Read from the connection and write to StdOut
	n, _ := io.Copy(os.Stdout, connection)
	fmt.Printf("\nThere were %d bytes written.\n\n", n)
}


func main() {
	//location_regex := regexp.MustCompile(`^\w+=`)
	//server_regex := regexp.MustCompile(`localhost:\w+$`)


	for _, v := range os.Args[1:]{
		input := strings.Split(v, "=")
		if len(input) != 2{
			err := errors.New("Input is problematic.")
			fmt.Fprintf(os.Stderr, "Encountered the following error: %v\n", err)
			os.Exit(20)
		}
		location := input[0]
		server_addr := input[1]
		fmt.Printf("Trying location, %s, at the following address: \u001b[31m%s\u001b[0m\n", strings.ToUpper(location), server_addr)
		make_connection(server_addr)
	}

}
