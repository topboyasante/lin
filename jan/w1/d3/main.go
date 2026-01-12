package main

import (
	"fmt"
	"log"
	"w1/d3/configparser"
)

func main() {
	config, err := configparser.ParseConfig("config.txt")
	if err != nil {
		log.Fatalf("error parsing file: %s", err)
	}

	fmt.Println("App Config", config)
}
