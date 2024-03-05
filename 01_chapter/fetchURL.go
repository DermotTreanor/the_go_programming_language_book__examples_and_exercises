package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("https://styles.redditmedia.com/t5_2rc7j/styles/communityIcon_wy4riduoe9k11.png") //http://127.0.0.1:8000/support_ireland")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch error:%v\n", err)
		os.Exit(1)
	}
	// fmt.Print(resp)
	byteSlice, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fetch: %v\n", err)
	}
	fmt.Print(byteSlice)
	file, err := os.OpenFile("image", os.O_WRONLY|os.O_CREATE, 0666)
	file.Write(byteSlice)
}
