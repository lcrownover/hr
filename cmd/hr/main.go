package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

var wg sync.WaitGroup

func getIPHostResult(h string, d string) {
	defer wg.Done()
	var res string
	addrs, err := net.LookupHost(h)
	if err != nil {
		res = "not_found"
	} else {
		res = addrs[0]
	}
	fmt.Printf("%s%s%s\n", h, d, res)

}

func main() {
	var dFlag = flag.String("d", ";", "output delimiter")
	flag.Parse()
	if flag.NArg() < 1 {
		fmt.Println("  usage: hr [-d=':'] host1 host2 host3")
		flag.PrintDefaults()
		os.Exit(0)
	}
	hostList := flag.Args()

	for _, h := range hostList {
		wg.Add(1)
		go getIPHostResult(h, *dFlag)
	}

	wg.Wait()

}
