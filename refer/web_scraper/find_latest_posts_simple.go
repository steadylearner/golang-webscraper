package main

import (
	"fmt"
	"log"

	"encoding/json"

	"github.com/go-resty/resty/v2"
	. "github.com/logrusorgru/aurora"
)

// Refer to this.
// https://github.com/steadylearner/Rust-Full-Stack/blob/master/bots/teloxide/src/community-bots/models/subreddit.rs

type Post struct {
	Title string
	Url   string
}

type Child struct {
	Data Post
}

type Children struct {
	Children []Child
}

type Response struct {
	Data Children
}

func main() {
	target := "https://www.reddit.com/r/golang/new/.json?limit=10"
	client := resty.New()

	// https://github.com/go-resty/resty#simple-get
	response, err := client.R().Get(target)
	if err != nil {
		log.Fatal(err)
	} else {
		// fmt.Printf("%s\n", response) // body

		bodyBytes := response.Body()
		response := Response{}

		err := json.Unmarshal(bodyBytes, &response)
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println(response.Data)

		for index, post := range response.Data.Children {
			num := index + 1
			stdoutLink := fmt.Sprintf("%d. %s(%s)", num, post.Data.Title, Blue(post.Data.Url))
			fmt.Println(stdoutLink)
		}
	}
}
