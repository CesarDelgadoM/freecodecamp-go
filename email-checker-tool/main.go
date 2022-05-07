package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMX, hasSPF, sprRecord, hasDMARC, dmarcRecord\n")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	err := scanner.Err()
	if err != nil {
		fmt.Printf("error: could not read from input: %v\n", err)
		os.Exit(1)
	}
}

func checkDomain(domain string) {
	var hasMX bool
	var hasSPF bool
	var hasDMARC bool
	var spfRecord string
	var dmarcRecord string
	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	fmt.Println(*mxRecords[0])
	if len(mxRecords) > 0 {
		hasMX = true
	}
	fmt.Println(hasMX)
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}
	fmt.Printf("%s, hasMX:%t, hasSPF:%t, %s, hasDMARC:%t, %s\n",
		domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)
}
