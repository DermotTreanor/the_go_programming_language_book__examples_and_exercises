package main

import (
	"encoding/json"
	"log"
	"os"
	"practest/test/testtypes"
)

func main() {
	testData := []testtypes.TestRun{
		{
			Phrase: "I am here to kick ass and chew bubblegum\nAnd €I'm all out of bubblegum",
			TestID: 1,
			Want:   70,
		},
		{
			Phrase: " ",
			TestID: 2,
			Want:   1,
		},
		{
			Phrase: "abcdefghijklmnopqrstuvwxyz",
			TestID: 3,
			Want:   26,
		},
		{
			Phrase: "\t\n\t\n\t",
			TestID: 4,
			Want:   5,
		},
		{
			Phrase: "The quick brown fox jumped over the lazy dog.",
			TestID: 5,
			Want:   45,
		},
		{
			Phrase: "!\"£$%^&*(')",
			TestID: 6,
			Want:   11,
		},
		{
			Phrase: "Hello, 世界",
			TestID: 7,
			Want:   9,
		},
		{
			Phrase: "Fīat iūstitia ruat cælum",
			TestID: 8,
			Want:   24,
		},
		{
			Phrase: "",
			TestID: 9,
			Want:   0,
		},
	}

	b, err := json.MarshalIndent(testData, "", "\t")
	if err != nil {
		log.Printf("main: error while marshalling testData: %v\n", err)
	}

	f, err := os.OpenFile("./test-data.json", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	defer f.Close()
	if err != nil {
		log.Printf("main: error while opening file to write json to: %v\n", err)
	}

	_, err = f.Write(b)
	if err != nil {
		log.Printf("main: error writing to file: %v\n", err)
	}
}
