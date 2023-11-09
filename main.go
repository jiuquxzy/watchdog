package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/CESSProject/watchdog/client"
)

func main() {

	if len(os.Args) < 2 {
		log.Fatal("No URL entered")
	}
	url := os.Args[1]
	client.InitWatchdogClient()

	var freq time.Duration
	if len(os.Args) >= 3 {
		d, err := strconv.Atoi(os.Args[2])
		if err != nil {
			log.Fatal("error parsing second parameter")
		}
		freq = time.Duration(d)
	}
	err := client.RunWatchdogClient(url, freq)
	log.Println("failed to launch program", err)
}