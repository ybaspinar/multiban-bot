package main

import "github.com/nicklaw5/helix"

func main() {
	client, err := helix.NewClient(&helix.Options{
		ClientID: "your-client-id",
	})
	if err != nil {
		// handle error
	}
}
