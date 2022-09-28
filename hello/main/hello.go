package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/example/stringutil"
)

func main() {
	fmt.Println(stringutil.ToUpper("Hello"))
	ticker := time.NewTicker(2 * time.Second)
	// ch := make(chan bool)
	go func(ticker *time.Ticker) {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				log.Println("Ticker running...")
				// case stop := <-ch:
				// 	if stop {
				// 		log.Println("Ticker Reset")
				// 	}
			}
		}
	}(ticker)
	time.Sleep(6 * time.Second)
	log.Println("Ticker Reset")
	ticker.Reset(1 * time.Second)
	// ch <- true
	time.Sleep(4 * time.Second)
	log.Println("Main end")
	// close(ch)
}
