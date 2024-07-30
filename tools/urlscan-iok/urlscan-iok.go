package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	iok "phish.report/IOK/v2"
)

var uuid = flag.String("uuid", "", "[required] the urlscan.io result ID to scan for IOKs")

func main() {
	flag.Parse()
	if *uuid == "" {
		flag.Usage()
		os.Exit(1)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM)
	defer cancel()
	input, err := iok.InputFromURLScan(ctx, *uuid, http.DefaultClient)
	if err != nil {
		log.Fatal(err)
	}

	matches, err := iok.GetMatches(input)
	if err != nil {
		log.Fatal(err)
	}

	if len(matches) == 0 {
		fmt.Println("No matches!")
		return
	}

	fmt.Println("Matching indicators:")
	for _, match := range matches {
		fmt.Println("\t*", match.Title, "https://phish.report/IOK/indicators/"+match.ID)
	}
}
