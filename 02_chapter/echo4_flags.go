package main

import (
	"flag"
	"fmt"
)

func main() {
	var n = flag.Bool("n", false, "prints something")
	
	//If we don't add this line then we will be stuck with just the defaults:
	flag.Parse()
	fmt.Println(*n)

	fmt.Println(flag.Arg(0))
	return
}
