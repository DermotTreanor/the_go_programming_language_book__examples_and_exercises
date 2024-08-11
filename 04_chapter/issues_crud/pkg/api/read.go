package api

import (
	"fmt"
	//"log"
	//"os"
)

//func useEditor()(bool, error){
	//fmt.Println("Do you want to use Vim to add issue data? [y/N]")
	//var v string
	//fmt.Scanln(&v)
	//if v[0] == 'y' || v[0] == 'Y'{
		//f, err := os.CreateTemp("", "github_issues_crud_*")
		//if err != nil{
			//log.Println()
		//}
	//}
//}

func ReadIssue() {
	fmt.Print("Please specify the owner, repo, and issue number you want: ")
	var rCall call
	fmt.Scan(&rCall.owner, &rCall.repo, &rCall.iNum)
	i := rCall.makeRequest()
	//fmt.Printf("The value for i is: %v\n", i)
	fmt.Println(i.Title)
}

func CreateIssue() {
	var cCall call
	fmt.Print("Please specify the owner and repo: ")
	fmt.Scan(&cCall.owner, &cCall.repo)

	cCall.makeRequest()
}
