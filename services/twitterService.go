package services

import (
	"dirty-tweets-deleter-web/models"
	"fmt"
	goTwitter "github.com/dghubble/go-twitter/twitter"
	"strconv"
	"strings"
)

type TwitterService struct {}

func (service TwitterService) CleanTweets(tweets []models.Tweet, swearWords [][]string, onlySwearwords bool) (tweetsToDelete []models.Tweet){
	for _, tweet := range tweets {
		if onlySwearwords {
			for _, swearWord := range swearWords[0] {
				if strings.Contains(strings.ToLower(tweet.Text), " "+swearWord+" ") {
					tweetsToDelete = append(tweetsToDelete, tweet)
					break
				}
			}
		} else {
			tweetsToDelete = append(tweetsToDelete, tweet)
		}
	}

	return
}

func (service TwitterService) DeleteTweets(client *goTwitter.Client, tweetsToDelete []models.Tweet) {
	for i, tweet := range tweetsToDelete {
		currentIndex := fmt.Sprintf("%.2f",float64(i)/float64(len(tweetsToDelete)) * 100)
		id, _ := strconv.ParseInt(tweet.Tweet_id, 0, 64)
		client.Statuses.Destroy(id, nil)
		fmt.Println("Deleting tweets... " + currentIndex + "%")
	}
}