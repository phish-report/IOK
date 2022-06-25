package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	iok "phish.report/IOK"
)

type urlscanResult struct {
	Data struct {
		Cookies []struct {
			Name   string `json:"name"`
			Value  string `json:"value"`
			Path   string `json:"path"`
			Domain string `json:"domain"`
		} `json:"cookies"`
	} `json:"data"`
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	scanResp, err := http.Get("https://urlscan.io/api/v1/result/" + flag.Args()[0])
	if err != nil {
		return err
	}
	scanResult := urlscanResult{}
	json.NewDecoder(scanResp.Body).Decode(&scanResult)

	domResp, err := http.Get("https://urlscan.io/dom/" + flag.Args()[0])
	if err != nil {
		return err
	}
	dom, _ := io.ReadAll(domResp.Body)

	cookies := []http.Cookie{}
	for _, cookie := range scanResult.Data.Cookies {
		cookies = append(cookies, http.Cookie{
			Name:   cookie.Name,
			Value:  cookie.Value,
			Path:   cookie.Path,
			Domain: cookie.Domain,
		})
	}
	matches, err := iok.GetMatches(iok.Input{
		Dom:     string(dom),
		Cookies: cookies,
	})
	fmt.Println("Matching indicators:")
	fmt.Println()
	for _, match := range matches {
		fmt.Println("\t*", match.Title)
	}
	return nil
}
