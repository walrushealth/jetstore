package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var apiSecret          = flag.String("API_SECRET", "", "Secret used for signing jwt tokens (required)")
var dropTable          = flag.Bool  ("d", false, "drop users table if it exists, default is false")
var dsn                = flag.String("dsn", "", "primary database connection string (required)")
var serverAddr         = flag.String("serverAddr", ":8080", "server address to ListenAndServe (required)")

func main() {
	flag.Parse()
	hasErr := false
	var errMsg []string
	if *apiSecret == "" {
		hasErr = true
		errMsg = append(errMsg, "API_SECRET must be provided.")
	}
	if *dsn == "" {
		hasErr = true
		errMsg = append(errMsg, "dsn for primary database node (-dsn) must be provided.")
	}
	if *serverAddr == "" {
		hasErr = true
		errMsg = append(errMsg, "Server address (-serverAddr) must be provided.")
	}
	if hasErr {
		flag.Usage()
		for _, msg := range errMsg {
			fmt.Println("**",msg)
		}
		os.Exit((1))
	}

	fmt.Println("apiserver argument:")
	fmt.Println("-------------------")
	fmt.Println("Got argument: apiSecret",*apiSecret)
	fmt.Println("Got argument: dropTable",*dropTable)
	fmt.Println("Got argument: dsn",*dsn)
	fmt.Println("Got argument: serverAddr",*serverAddr)

	log.Fatal(listenAndServe())
}