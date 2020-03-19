// go run main.go -subreddit="golang" -limit="10"
// "https://www.reddit.com/r/golang/new/.json?limit=10" 

package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"strconv"

	"github.com/go-resty/resty/v2"
	"github.com/logrusorgru/aurora"

        // https://golang.org/doc/code.html
        // $go build first in models/
        // Refer to go.mod file also.
        Models "steadylearner.com/models"
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

	// https://godoc.org/github.com/go-resty/resty#Request.SetResult
        // https://github.com/go-resty/resty/blob/master/response.go#L59
	resp, err := resty.New().R().
		SetPathParams(map[string]string{"subreddit": subreddit}).
		SetQueryParam("limit", strconv.Itoa(limit)).
	        SetResult(&Models.Response{}).
		Get(target)
	if err != nil {
		log.Fatal(err)
		return
	}

        // Type conversion to Models.Response so verify ok value here?
	if response, ok := resp.Result().(*Models.Response); !ok {
		log.Fatal(errors.New("invalid response format"))
	} else {
		for index, post := range response.Data.Children {
			num := index + 1
			stdoutLink := fmt.Sprintf("%d. %s (%s)", num, post.Data.Title, au.Blue(post.Data.URL))
			fmt.Println(stdoutLink)
		}
	}
}
