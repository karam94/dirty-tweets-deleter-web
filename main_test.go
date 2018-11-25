package main

import (
	"dirty-tweets-deleter-web/models"
	"dirty-tweets-deleter-web/services"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_OnlyTweetsWithSwearwordsAreRemoved(t *testing.T){
	var tweets = []models.Tweet {
		{
			Tweet_id:  "1",
			Timestamp: "2018-11-15 08:56:54 +0000",
			Text:      "this is a karam tweet",
		},
		{
			Tweet_id: "2",
			Timestamp: "2018-11-15 08:56:54 +0000",
			Text: "this is a good tweet",
		},
		{
			Tweet_id: "3",
			Timestamp: "2018-11-15 08:56:54 +0000",
			Text: "this is a nice tweet",
		},
	}

	var swearWords [][]string
	swearWord := []string{"karam"}
	swearWords = append(swearWords, swearWord)

	tweetsToDelete := services.TwitterService{}.CleanTweets(tweets, swearWords, true)
	assert.Equal(t, 1, len(tweetsToDelete), "Tweet containing swear word should be removed.")
	assert.Equal(t, tweetsToDelete[0].Tweet_id, "1")
}

func Test_NoSwearwordsSoNoTweetsAreRemoved(t *testing.T){
	var tweets = []models.Tweet {
		{
			Tweet_id: "1",
			Timestamp: "2018-11-15 08:56:54 +0000",
			Text: "this is a lovely tweet",
		},
		{
			Tweet_id: "2",
			Timestamp: "2018-11-16 08:56:54 +0000",
			Text: "this is a good tweet",
		},
		{
			Tweet_id: "3",
			Timestamp: "2018-11-16 08:56:54 +0000",
			Text: "this is a nice tweet",
		},
	}

	var swearWords [][]string
	swearWord := []string{"karam"}
	swearWords = append(swearWords, swearWord)

	tweetsToDelete := services.TwitterService{}.CleanTweets(tweets, swearWords, true)
	assert.Equal(t, 0, len(tweetsToDelete), "One tweet will be removed.")
}