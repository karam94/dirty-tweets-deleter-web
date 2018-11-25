package main

import (
	"dirty-tweets-deleter-web/helpers"
	"dirty-tweets-deleter-web/readers"
	"dirty-tweets-deleter-web/services"
	"errors"
	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"golang.org/x/oauth2"
	"io"
	"os"
	"path/filepath"

	"github.com/markbates/goth"
	"github.com/markbates/goth/providers/twitter"

	goTwitter "github.com/dghubble/go-twitter/twitter"
)

var sessionsManager *sessions.Sessions
var cfg = helpers.ConfigHelper{}.LoadConfiguration()

func init() {
	cookieName := "dirty-tweets-deleter-cookie"
	// AES only supports key sizes of 16, 24 or 32 bytes.
	// You either need to provide exactly that amount or you derive the key from what you type in.
	hashKey := []byte(cfg.CookieHashKey)
	blockKey := []byte(cfg.CookieBlockKey)
	secureCookie := securecookie.New(hashKey, blockKey)

	sessionsManager = sessions.New(sessions.Config{
		Cookie: cookieName,
		Encode: secureCookie.Encode,
		Decode: secureCookie.Decode,
	})
}

var CheckUserLoggedIn = func(ctx iris.Context) (goth.User, error) {
	provider, err := goth.GetProvider("twitter")
	if err != nil {
		return goth.User{}, err
	}
	session := sessionsManager.Start(ctx)
	value := session.GetString("twitter")
	if value == "" {
		return goth.User{}, errors.New("session value for " + "twitter" + " not found")
	}

	sess, err := provider.UnmarshalSession(value)
	if err != nil {
		return goth.User{}, err
	}

	user, err := provider.FetchUser(sess)
	if err == nil {
		// user can be found with existing session data
		return user, err
	}

	// get new token and retry fetch
	_, err = sess.Authorize(provider, ctx.Request().URL.Query())
	if err != nil {
		return goth.User{}, err
	}

	session.Set("twitter", sess.Marshal())
	return provider.FetchUser(sess)
}

func main() {
	goth.UseProviders(
		twitter.New(cfg.ConsumerKey, cfg.ConsumerSecret, cfg.CallbackUrl),
	)

	app := iris.New()
	app.RegisterView(iris.HTML("./templates", ".html"))
	app.Layout("layouts/layout.html")

	app.Get("/login", func(ctx iris.Context) {
		if user, err := CheckUserLoggedIn(ctx); err == nil {
			ctx.ViewData("", user)
			if err := ctx.View("user.html"); err != nil {
				ctx.Writef("%v", err)
			}
		} else {
			services.TwitterAuthService{}.BeginAuthHandler(ctx, sessionsManager)
		}
	})

	app.Get("/login/callback", func(ctx iris.Context) {
		if user, err := CheckUserLoggedIn(ctx); err == nil {
			ctx.ViewData("", user)

			if err := ctx.View("user.html"); err != nil {
				ctx.Writef("%v", err)
			}
		} else {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.Writef("%v", err)
			return
		}
	})

	app.Get("/logout", func(ctx iris.Context) {
		session := sessionsManager.Start(ctx)
		session.Delete("twitter")
		ctx.Redirect("/", iris.StatusTemporaryRedirect)
	})

	app.Get("/", func(ctx iris.Context) {
		if user, err := CheckUserLoggedIn(ctx); err == nil {
			ctx.ViewData("", user)
			if err := ctx.View("user.html"); err != nil {
				ctx.Writef("%v", err)
			}
		} else {
			ctx.ViewData("", "twitter")
			if err := ctx.View("index.html", nil); err != nil {
				ctx.Writef("%v", err)
			}
		}
	})

	app.Post("/deleteDirtyTweets", iris.LimitRequestBodySize(10<<20),
		func(ctx iris.Context) {
			if user, err := CheckUserLoggedIn(ctx); err == nil {
				ctx.ViewData("", user)

				config := &oauth2.Config{}
				token := &oauth2.Token{AccessToken: user.AccessToken}
				httpClient := config.Client(oauth2.NoContext, token)
				client := goTwitter.NewClient(httpClient)

				file, info, err := ctx.FormFile("uploadfile")

				if filepath.Ext(info.Filename) != ".csv" {
					ctx.StatusCode(iris.StatusInternalServerError)
					ctx.HTML("Error while uploading:" + info.Filename + " <b>Wrong file type must be .csv</b>")
					return
				}

				if err != nil {
					ctx.StatusCode(iris.StatusInternalServerError)
					ctx.HTML("Error while uploading:" + info.Filename + " <b>" + err.Error() + "</b>")
					return
				}

				defer file.Close()

				// Create a file with the same name
				out, err := os.OpenFile("./uploads/"+user.UserID+".csv",
					os.O_WRONLY|os.O_CREATE, 0666)

				if err != nil {
					ctx.StatusCode(iris.StatusInternalServerError)
					ctx.HTML("Error while uploading: <b>" + err.Error() + "</b>")
					return
				}

				defer out.Close()

				io.Copy(out, file)

				tweets := readers.TweetReader{}.ReadTweetsCsv(user.UserID)
				swearWords := readers.SwearWordReader{}.ReadSwearWordsCsv()
				toDelete := services.TwitterService{}.CleanTweets(tweets, swearWords, true)
				services.TwitterService{}.DeleteTweets(client, toDelete)
			} else {
				services.TwitterAuthService{}.BeginAuthHandler(ctx, sessionsManager)
			}
		})

	app.Run(iris.Addr("localhost:3000"))
}