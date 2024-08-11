package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"

	"gopl.io/ch4/github"
)

type call struct {
	owner string
	repo  string
	iNum  int
	issue github.Issue
	
}


func (c *call) makeRequest() github.Issue {
	var numPath string
	if c.iNum > 1 {
		numPath = fmt.Sprintf("/%d", c.iNum)
	}
	path := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", c.owner, c.repo) + numPath

	m := "GET"
	var br io.Reader
	if c.issue != (github.Issue{}) {
		m = "POST"
		t, _ := json.Marshal(c.issue.Title)
		data := fmt.Sprintf(`{"title":%s}`, t)
		fmt.Println(data)
		//if err != nil {
			//log.Fatalf("call.formRequest: unable to marshal call.issue: \n%v", err)
		//}
		br = strings.NewReader(data)

	}
	req, err := http.NewRequest(m, path, br)
	if err != nil {
		log.Fatalf("call.formRequest: Couldn't make request: \n%v", err)
	}
	pak, err := os.ReadFile("./pak")
	if err != nil{
		log.Printf("makeRequest: couldn't open file to get pak: \n%v", err)
	}
	fmt.Println(string(pak))
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", string(pak)))

	var ir github.Issue
	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		fmt.Println(err)
	}
	d, _ := io.ReadAll(resp.Body)
	fmt.Println(string(d))
	json.Unmarshal([]byte(d), &ir)
	return ir

}
