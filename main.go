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
		checkDomain(sc.Text())
	}
	if sc.Err() != nil {
		log.Fatalf(sc.Err().Error())
	}
}

func checkDomain(domain string) {
	// var hasMX, hasSPF, hasDMARCH bool
	// var spfRecord, dmarcRecord string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Something went wrong\n %q", err)
	}
	fmt.Print(mxRecords)

	// similarly check for the spf record
	// similarly check for the dmarc record
}
