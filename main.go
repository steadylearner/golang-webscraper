package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"

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

var (
	lang  string
	limit int
)

func init() {
	flag.StringVar(&lang, "lang", "golang", "Programming Language")
	flag.IntVar(&limit, "limit", 10, "Limit")
	flag.Parse()
}

func main() {
	target := "https://www.reddit.com/r/{lang}/new/.json"

	// https://github.com/go-resty/resty#simple-get
	resp, err := resty.New().R().
		SetPathParams(map[string]string{"lang": lang}).
		SetQueryParam("limit", strconv.Itoa(limit)).
		SetResult(&Response{}).
		Get(target)
	if err != nil {
		log.Fatal(err)
		return
	}

	if response, ok := resp.Result().(*Response); !ok {
		log.Fatal(errors.New("invalid response format"))
		return
	} else {
		for index, post := range response.Data.Children {
			num := index + 1
			stdoutLink := fmt.Sprintf("%d. %s (%s)", num, post.Data.Title, Blue(post.Data.Url))
			fmt.Println(stdoutLink)
		}
	}
}
