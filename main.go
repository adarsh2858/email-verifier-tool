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
	r := bufio.NewReader(os.Stdin)

	fmt.Printf("Input your name:")
	name, err := r.ReadString('\n')
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf("User's name: %s\n", strings.TrimSpace(name))
	// waits here

	sc := bufio.NewScanner(os.Stdin)

	fmt.Println("domain,hasMX,hasSPF,spfRecord,hasDMARC,dmarcRecord")
	for sc.Scan() {
		CheckDomain(sc.Text())
	}
	if sc.Err() != nil {
		log.Fatalf(sc.Err().Error())
	}
}

func CheckDomain(domain string) (isValidDomain bool) {
	var hasMX, hasSPF, hasDMARC bool
	var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Something went wrong\n %q", err)
		return
	}
	if len(mxRecords) > 0 {
		hasMX = true
	}

	// similarly check for the spf record
	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Error:\n%q", err)
		return
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSPF = true
			spfRecord = record
			break
		}
	}

	// similarly check for the dmarc record
	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Error:\n%q", err)
		return
	}

	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDMARC = true
			dmarcRecord = record
			break
		}
	}

	fmt.Printf("%v,%v,%v,%v,%v,%v", domain, hasMX, hasSPF, spfRecord, hasDMARC, dmarcRecord)

	if hasMX || hasSPF || hasDMARC {
		isValidDomain = true
	}
	return
}
