package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:8000/support_ireland")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch error:%v\n", err)
		os.Exit(1)
	}
	// fmt.Print(resp)
	byteSlice, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
	}
	fmt.Printf("%s\n", byteSlice)

}
