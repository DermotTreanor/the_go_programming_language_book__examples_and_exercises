package main

import (
	"fmt"
	"issues/pkg/api"

	"log"
)

func main() {

	fmt.Println("Select an operation to perform on an issue: [C]reate, [R]ead, [U]pdate, or [D]elete")
	var in string
	fmt.Scanf("%s\n", &in)
	fmt.Println(in)
	switch in[0] {
	case byte('c'), byte('C'):
		fmt.Println("You want to create")
		api.CreateIssue()
	case byte('r'), byte('R'):
		fmt.Println("You want to read")
		api.ReadIssue()
	case byte('u'), byte('U'):
		fmt.Println("You want to update")
		api.UpdateIssue()
	case byte('d'), byte('D'):
		fmt.Println("You want to delete")
		api.DeleteIssue()
	default:
		log.Fatal("Invalid option selected.")
	}

}
