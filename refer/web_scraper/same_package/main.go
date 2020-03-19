// go run main.go -subreddit="golang" -limit="10"

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/logrusorgru/aurora"
)

var (
	subreddit string
	limit     int
	au        aurora.Aurora
)

func init() {
	flag.StringVar(&subreddit, "subreddit", "golang", "Subreddit")
	flag.IntVar(&limit, "limit", 10, "Limit")
	flag.Parse()

	au = aurora.NewAurora(true)
}

func main() {
	target := "https://www.reddit.com/r/{subreddit}/new/.json"

	// https://github.com/go-resty/resty#simple-get
	// 2. Should read the documenation here.
	resp, err := resty.New().R().
		SetPathParams(map[string]string{"subreddit": subreddit}).
		SetQueryParam("limit", strconv.Itoa(limit)).
	        SetResult(&Response{}).
		Get(target)
	if err != nil {
		log.Fatal(err)
		return
	}

	if response, ok := resp.Result().(*Response); !ok {
		log.Fatal(errors.New("invalid response format"))
	} else {
		for index, post := range response.Data.Children {
			num := index + 1
			stdoutLink := fmt.Sprintf("%d. %s (%s)", num, post.Data.Title, au.Blue(post.Data.URL))
			fmt.Println(stdoutLink)
		}
	}
}
