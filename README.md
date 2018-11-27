# dirty-tweets-deleter-web
A small project to try out Go as part of a personal monthly challenge on my [blog](http://www.karam.io/2018/November-2018-Best-thing-since-bread/). This is a spin-off of the original [console application](https://github.com/karam94/dirty-tweets-deleter) and allows users to bulk delete all old tweets that contain swear words based on a pre-defined comma separated file of words.

**DISCLAIMER:** Running this application MAY delete all of your tweets. I am not to be held responsible for the result of your actions of using this application.

## Getting started
You need a [Twitter developers account](https://apps.twitter.com/) and a registered application. You can then specify your application's Twitter API consumer key, consumer secret, access token and access secret within a file called "config.json". You also need to specify a hash key and block key to encrypt your session cookies. These can also be placed within your "config.json". For more information, read [here](http://github.com/gorilla/securecookie). You may use "config.json.template" as a template to work off to create your config.

You can download the archive of your Twitter accounts tweets by going to Twitter > Settings > Account > Content > Request your archive. You can then upload the tweet archive CSV file e-mailed to you to the web application.

Place your own CSV of words in the root folder as "swearwords.csv". You can download a generic one from [here](http://www.bannedwordlist.com/) if you prefer.

## How to use
Providing you have [Go](https://golang.org/) installed on your machine, you can compile the code by running `go build` in the root directory which will give you an .exe file to run. E.g.`./main.go` or `./main.exe`. You can also run the unit tests by running `go test` in the root directory.

## dirty-tweets-deleter
As this application was created for learning purposes, I also created a similar console based alternative that also allows you to delete all tweets should you wish to do so, can be found [here](https://github.com/karam94/dirty-tweets-deleter).

## Credits
- [go-twitter by dghubble](https://github.com/dghubble/go-twitter)
- [securecookie by gorilla](https://github.com/gorilla/securecookie)
- [iris by kataras](https://github.com/kataras/iris)
- [goth by markbates](https://github.com/markbates/goth)
