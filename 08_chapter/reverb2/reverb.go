// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 224.

// Reverb2 is a TCP server that simulates an echo.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

// !+
func handleConn(c net.Conn) {
	var magic_counter sync.WaitGroup
	ch := make(chan struct{})
	input := bufio.NewScanner(c)

	for input.Scan() {
		magic_counter.Add(1)

		go func(c net.Conn, shout string, delay time.Duration) {
			defer magic_counter.Done()
			fmt.Fprintln(c, "\t", strings.ToUpper(shout))
			time.Sleep(delay)
			fmt.Fprintln(c, "\t", shout)
			time.Sleep(delay)
			fmt.Fprintln(c, "\t", strings.ToLower(shout))
			ch <- struct{}{}
		}(c, input.Text(), 1*time.Second)
	}
	go func() {
		magic_counter.Wait()
		close(ch)
	}()
	for range ch {
	}
	c.Close()

}

//!-

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
