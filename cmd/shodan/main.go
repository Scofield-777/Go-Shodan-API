package main

import(
	"fmt"
	"os"
	"log"
	"shodan-layout/shodan"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatalln("Usage: shodan searchterm")
	}
	apiKey := os.Getenv("Shodan_API_Key")
	s := shodan.New(apiKey)
	fmt.Printf("%s\n",s)
	info, err := s.APIInfo()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("Query Credits: %d\nScan Credits:	%d\n\n", info.QueryCredits, info.ScanCredits)
	hostSearch, err := s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("IP Address        Port\n")
	for _, host := range hostSearch.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
