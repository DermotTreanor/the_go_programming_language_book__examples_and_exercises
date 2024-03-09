package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	//"time"
)

func main() {
	//Here we send the request and return a pointer to the response struct.
	value := 120
	for _, url := range os.Args[1:] {
		resp, err := http.Get(url)
		if err != nil {
			//If there is an error then send it to the standard error file/device
			fmt.Fprintf(os.Stderr, "fetch error:%v\n", err)
			os.Exit(1)
		}

		//The body field is a stream that is read and outputs a slice of bytes. It's then closed.
		fmt.Print(fmt.Sprintf("\u001b[38;5;%dm", value))
		byteSlice, err := io.ReadAll(resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
		}
		//We print the slice of bytes in a string format.
		fmt.Printf("%s\n", byteSlice)
		fmt.Print("\u001b[0m")
		fmt.Println(resp.Status)

		value += 20
		//time.Sleep(5 * time.Second)
	}
}
