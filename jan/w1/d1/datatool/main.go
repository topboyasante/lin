package main

import (
	"encoding/json"
	"flag"
	"fmt"
)

func processArray(data []any) {
	for i, v := range data {
		switch v := v.(type) {
		case string:
			fmt.Println(i, "is a string")
		case float64:
			fmt.Println(i, "is an number")
		case bool:
			fmt.Println(i, "is a boolean")
		case map[string]any:
			processMap(v)
		case []any:
			processArray(v)
		}
	}
}

func processMap(data map[string]any) {

	for k, v := range data {
		switch v := v.(type) {
		case string:
			fmt.Println(k, "is a string")
		case float64:
			fmt.Println(k, "is an number")
		case bool:
			fmt.Println(k, "is a boolean")
		case map[string]any:
			processMap(v)
		case []any:
			processArray(v)
		}
	}
}

func ProcessDataTypes() {
	// Store the JSON data from the command line argument
	var jsonStr string

	// Use a map to hold the unmarshaled data
	var result map[string]any

	// create a cli flage called "data" to accept json string
	flag.StringVar(&jsonStr, "data", "", "JSON data as string")
	flag.Parse() // populate jsonStr with the values from command line

	// if no json string is provided, print a message and exit
	if jsonStr == "" {
		fmt.Println("No JSON data provided. Use -data flag to provide JSON string.")
		return
	}

	// convert the json string into a map
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)
		return
	}

	processMap(result)
}

func main() {
	ProcessDataTypes()
}
