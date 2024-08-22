package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"xkcdfetch/pkg/xkcdapi"
)

func main() {
	//slowPullComics()
	fastPullComics()
}

func slowPullComics() {
	//Set a flag to check that we skip the '404' joke comic
	f, err := os.OpenFile("./data/comics.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("main: could not open json file to write data. \n%v", err)
	}

	var i int = 1
	for {
		wrote, err := xkcdapi.GetComic(i, f)
		if err != nil {
			log.Printf("main: an error occurred when fetching the comic #%d: \n%v", i, err)
			continue
		} else if !wrote {
			//Check if we got a nil pointer even after the error was not nil. Then we must have run out of comics.
			//We must continue if i is 404 as we get a joke 404 response
			if i == 404 {
				i++
				continue
			}
			break
		}
		fmt.Printf("We have gotten #%d.\n", i)
		i++
	}
}

func fastPullComics() {
	f, err := os.OpenFile("./data/comics.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatalf("main: could not open json file to write data. \n%v", err)
	}
	ch := make(chan xkcdapi.ComicCh, 200)
	var cc xkcdapi.ComicCh
	i := 1
	prev := 0
	l := 200
	mu := sync.Mutex{}
	endFlag := false
requestloop:
	for {
		go xkcdapi.FastGetComic(i, f, ch, &mu)
		if i%l == 0 {
			for range l - prev {
				//Why is this not stopping... We should be waiting
				cc = <-ch
				if cc.Err == nil {
					continue
				} else if cc.Number == 404 {
					log.Println("Skipping 404 joke.")
				} else {
					log.Printf("Error on %d: %v\n", cc.Number, cc.Err)
					endFlag = true
				}
			}
			prev = l
			l += 200
		}
		i++
		if endFlag {
			break requestloop
		}
	}
}
