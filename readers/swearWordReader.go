package readers

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
)

type SwearWordReader struct {}

func (swearWordReader SwearWordReader) ReadSwearWordsCsv() (swearWords [][]string) {
	csvFile, _ := os.Open("swearWords.csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	swearWords, error := reader.ReadAll()

	if error != nil {
		log.Fatal(error)
	}

	return
}