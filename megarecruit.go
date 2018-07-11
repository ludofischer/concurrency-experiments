package main

import "log"
import (
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/clientcredentials"
)

func analyzeUser(user *twitter.User) {
	log.Print(user.Location)
	log.Print(user.ScreenName)
}

func getUsers(client *twitter.Client, cursor int64) (*twitter.Followers, error) {
	followers, _, err := client.Followers.List(&twitter.FollowerListParams{ScreenName: "csswizardry", Cursor: cursor})

	return followers, err
}

func main() {
	cursor := int64(-1)
	config := &clientcredentials.Config{}
	httpClient := config.Client(oauth2.NoContext)
	client := twitter.NewClient(httpClient)
	for {
		followers, err := getUsers(client, cursor)
		if err != nil {
			log.Fatalf("Error communicating with Twitter %s", err)
		}
		for _, user := range followers.Users {
			analyzeUser(&user)
		}
		cursor = followers.NextCursor
		if cursor == 0 {
			break
		}
	}
}
