package main

import (
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const (
	homeTimelineAPI = "https://api.twitter.com/1.1/statuses/home_timeline.json"
	tweetAPI        = "https://api.twitter.com/1.1/statuses/update.json"
)

type request struct {
	client *twitter.Client
}

func newRequest(c *cred) *request {
	config := oauth1.NewConfig(c.consumerKey, c.consumerSecret)
	token := oauth1.NewToken(c.accessToken, c.accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	return &request{
		client: twitter.NewClient(httpClient),
	}
}

func (r *request) post(tweet string) error {
	_, _, err := r.client.Statuses.Update(tweet, nil)
	return err
}

func (r *request) list() ([]twitter.Tweet, error) {
	tweets, _, err := r.client.Timelines.HomeTimeline(&twitter.HomeTimelineParams{
		Count: 20,
	})
	return tweets, err
}
