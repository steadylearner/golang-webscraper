package main

import (
    "fmt"
    "log"
     "github.com/go-resty/resty/v2"
)

func main() {
    target := "https://www.reddit.com/r/rust/new/.json?limit=10"
    client := resty.New()

    // https://github.com/go-resty/resty#simple-get
    response, err := client.R().Get(target)
    if err != nil {
        log.Fatal(err)
    } else {
        fmt.Printf("%s\n", response) // body
    }
}

