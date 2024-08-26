package main

import (
	"log"
	"time"
)

func main() {
	log.Println("starting scraper server...")

	// time.Sleep(30 * time.Second)
	// log.Println("starting scraperMain...")
	// scraperMain()
	go func() {
		for {
			time.Sleep(5 * time.Second)
			//scraperMain()
		}
	}()

	select {}
}
