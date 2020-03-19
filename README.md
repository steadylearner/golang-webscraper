<!-- https://docs.google.com/document/d/1Zb9GCWPKeEJ4Dyn2TkT-O3wJ8AFc-IMxZzTugNCjr-8/edit#heading=h.y3m7qt48hh5l -->

# Before you test this

[Refer to this to find how entire project works.](https://github.com/steadylearner/Rust-Full-Stack/tree/master/nodemon_with_other_languages)

You can use the first command if you want to live edit the project.

```console
1. $(yarn && cd models && go build && cd .. &&) ./go.sh
2. $go run main.go
```

## The current state of this project

Make a seprate project for gRPC.

https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go
clone this https://gitlab.com/pantomath-io/demo-grpc
https://grpc.io/docs/quickstart/go/
https://stackoverflow.com/questions/43682366/how-is-grpc-different-from-rest

1. gRPC, login, desktop?

https://blog.usejournal.com/authentication-in-golang-c0677bcce1a8
https://github.com/googleapis/googleapis/blob/master/google/rpc/error_details.proto
https://github.com/linux08/auth/blob/master/main.go

## Telebot

[You can make something similar to this.](https://github.com/teloxide/community-bots/tree/master/subreddit_reader)

![rust_example](https://raw.githubusercontent.com/teloxide/community-bots/master/subreddit_reader/rust_example.png)

## Infinite subreddit CLI

```console
Which subreddit you want to scrape?(golang)

How many you want?(50)
10
1. Rest Api with orm, need advice on project structure (https://www.reddit.com/r/golang/comments/fj1q8r/rest_api_with_orm_need_advice_on_project_structure/)
2. [net/url] invalid url escape %2f error after creating query using url (https://www.reddit.com/r/golang/comments/fj00h7/neturl_invalid_url_escape_2f_error_after_creating/)
3. Dynamic service discovery in Golang micro-service. (https://www.reddit.com/r/golang/comments/fizssa/dynamic_service_discovery_in_golang_microservice/)
4. Re-learning slice (https://kilabit.info/journal/2020/re-learning_slice/)
5. hostctl: manage your /etc/hosts like a pro (https://github.com/guumaster/hostctl)
6. Understand unsafe in GoLang (https://www.pixelstech.net/article/1584241521-Understand-unsafe-in-GoLang)
7. Using modules locally without publishing to VCS (https://www.reddit.com/r/golang/comments/fix89j/using_modules_locally_without_publishing_to_vcs/)
8. What are idiomatic patterns for getting inheritance-style dynamic dispatch? (https://www.reddit.com/r/golang/comments/fix630/what_are_idiomatic_patterns_for_getting/)
9. If Go channels are FIFO, why this prints the last one first always? (https://www.reddit.com/r/golang/comments/fix1k1/if_go_channels_are_fifo_why_this_prints_the_last/)
10. blog.golang.org: mistake in code (https://www.reddit.com/r/golang/comments/fiu8ba/bloggolangorg_mistake_in_code/)

End?([n]/y])

```
