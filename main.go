package main

import (
	"flag"
	"fmt"
)

func main() {
	url := flag.String("d", "", "The domain for get the IP")

	flag.Parse()
	if *url == "" {
		fmt.Println("Error: d flag is required")
		return
	}
	result, err := getIPAddress(*url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(result)
}
