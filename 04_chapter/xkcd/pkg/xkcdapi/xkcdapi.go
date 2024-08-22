package xkcdapi

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

type Comic struct {
	Month      string
	Num        int
	Link       string
	Year       string
	SafeTitle  string `json:"safe_title"`
	Transcript string
	Alt        string
	Img        string
	Title      string
	Day        string
}

var formatURL = "https://xkcd.com/%d/info.0.json"

func GetComic(num int, file *os.File) (bool, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf(formatURL, num), nil)
	if err != nil {
		return false, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	} else if resp.StatusCode == 404 {
		return false, nil
	}
	d, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}
	file.Write(d)
	return true, nil
}

type ComicCh struct {
	Err    error
	Wrote  bool
	Number int
}

func FastGetComic(num int, file *os.File, ch chan ComicCh, mutex *sync.Mutex) {
	fmt.Printf("Getting number %d\n", num)
	cc := ComicCh{
		Number: num,
	}
	req, err := http.NewRequest("GET", fmt.Sprintf(formatURL, num), nil)
	if err != nil {
		cc.Err = err
		ch <- cc
		return
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		panic("AHHHHH")
	}
	if resp.StatusCode != 200 {
		if cc.Number == 404 {
			cc.Err = nil
		} else {
			cc.Err = fmt.Errorf("Didn't get the right status code for comic #%d. Got number: %d", cc.Number, resp.StatusCode)
		}
		ch <- cc
		return
	}

	d, err := io.ReadAll(resp.Body)
	if err != nil {
		cc.Err = err
		ch <- cc
		return
	}
	d = append(d, byte('\n'))
	mutex.Lock()
	_, err = file.Write(d)
	if err != nil {
		cc.Err = err
		ch <- cc
		return
	}
	mutex.Unlock()
	//We send with a nil error and a true wrote
	cc.Wrote, cc.Err = true, nil
	ch <- cc
	return
}
