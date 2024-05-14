package main

import (
	"fmt"
	"log"
	"os"
	"shodan/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Usage: shodan [searchterm]")
	}

	apiKey := os.Getenv("SHODAN_API_KEY")
	client := shodan.New(apiKey)
	info, err := client.ApiStatus()

	if err != nil {
		log.Panicln(err)
	}

	fmt.Printf("Query Credits: %d\nScan Credits: %d\n\n", info.QueryCredits, info.ScanCredits)

	hostSearch, err := client.HostSearch(os.Args[1])

	if err != nil {
		log.Panicln(err)
	}

	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
