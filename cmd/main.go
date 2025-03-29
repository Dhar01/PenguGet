package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
)

var (
	errNoURL      = errors.New("no URL found")
	errInvalidURL = errors.New("invalid URL")
)

func main() {
	urlFlag := flag.String("url", "", "The URL to download from")
	flag.Parse()

	if ok := validateURL(urlFlag); !ok {
		log.Fatalf("Error parsing the URL")
		os.Exit(1)
	}

	fmt.Printf("Preparing to download: %s\n", *urlFlag)
}

func validateURL(urlFlag *string) bool {
	if urlFlag == nil {
		log.Printf("URL should be provided: %v", errNoURL)
		return false
	}

	parsedURL, err := url.Parse(*urlFlag)
	if err != nil || parsedURL.Scheme == "" || parsedURL.Host == "" {
		log.Printf("Error: %v", errInvalidURL)
		return false
	}

	return true
}
