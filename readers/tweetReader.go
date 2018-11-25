package readers

import (
	"bufio"
	"dirty-tweets-deleter-web/models"
	"encoding/csv"
	"io"
	"log"
	"os"
)

type TweetReader struct {}

func (tweetReader TweetReader) ReadTweetsCsv(userId string) (tweets []models.Tweet) {
	csvFile, _ := os.Open("./uploads/"+userId+".csv")
	reader := csv.NewReader(bufio.NewReader(csvFile))

	reader.Read() // skip headers in csv file

	for {
		line, error := reader.Read()

		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		tweets = append(tweets, models.Tweet {
			Tweet_id: line[0],
			Timestamp: line[3],
			Text: line[5],
		})
	}

	return
}