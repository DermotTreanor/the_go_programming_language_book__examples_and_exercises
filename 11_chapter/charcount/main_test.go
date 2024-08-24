package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"practest/test/testtypes"
	"regexp"
	"strconv"
	"testing"
)

func extractGot(rdr io.Reader) (int, bool) {
	p1 := `[1-4]\t\d+$`
	rgx1 := regexp.MustCompile(p1)

	p2 := `\d*$`
	rgx2 := regexp.MustCompile(p2)
	sc := bufio.NewScanner(rdr)
	tot := 0
	for sc.Scan() {
		tx := sc.Text()
		if rgx1.MatchString(tx) {
			s := rgx1.FindString(tx)
			i, err := strconv.Atoi(rgx2.FindString(s))
			if err != nil {
				log.Fatal("extractGot: can't convert count result to useful numeric")
			}
			tot += i
		}
	}
	return tot, true
}

func TestMain(t *testing.T) {
	d, err := os.ReadFile("./test/test-data.json")
	if err != nil {
		t.Errorf("TestMain: couldn't read json test data: %v\n", err)
	}
	var testData []testtypes.TestRun
	json.Unmarshal(d, &testData)
	for i := range testData {
		var err1, err2 error
		inFile, err1 = os.CreateTemp("", fmt.Sprintf("TestMain_%d_INPUT_*", testData[i].TestID))
		outFile, err2 = os.CreateTemp("", fmt.Sprintf("TestMain_%d_OUTPUT_*", testData[i].TestID))
		if (err1 != nil) || (err2 != nil) {
			t.Fatalf("TestMain: couldn't open temporary files in place of stdin and stdout:\n%v\n%v\n", err1, err2)
		}

		fmt.Fprint(inFile, testData[i].Phrase)
		inFile.Seek(0, 0)
		main()
		outFile.Seek(0, 0)
		got, found := extractGot(outFile)
		if !found {
			t.Fatal("TestRegPhrase can't get result from regex analysis of output")
		}
		if got != testData[i].Want {
			fs := "main(), got %d and want %d for input:\n%q\nSee %s for file used."
			t.Errorf(fs, got, testData[i].Want, testData[i].Phrase, outFile.Name())
			continue
		}
		os.Remove(inFile.Name())
		os.Remove(outFile.Name())
	}
}
