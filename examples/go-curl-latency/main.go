package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
	"io"
)

func main() {
	flag.Parse()

	if flag.NArg() != 2 {
		fmt.Fprintf(os.Stderr, "2 required args: URL and number requests\n")
		os.Exit(1)
	}

	url := flag.Arg(0)
	numRequestsString := flag.Arg(1)
	numRequests, _ := strconv.Atoi(numRequestsString)
	for count := 0; count < numRequests; count++ {
		start := time.Now()			
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Get(%s): %s", url, err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		elapsed := time.Since(start)
		fmt.Printf("Time for request: %s\n", elapsed)	
		n, err := io.Copy(os.Stdout, resp.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Copy: %s", err)
		}
		fmt.Fprintf(os.Stderr, "\ncopied %d bytes\n", n)
	}

}