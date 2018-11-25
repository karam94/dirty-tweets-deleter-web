package services

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/sessions"
	"github.com/markbates/goth"
)

type TwitterAuthService struct {}

func (service TwitterAuthService) BeginAuthHandler(ctx iris.Context, sessionsManager *sessions.Sessions) {
	url, err := GetTwitterAuthURL(ctx, sessionsManager)
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.Writef("%v", err)
		return
	}

	ctx.Redirect(url, iris.StatusTemporaryRedirect)
}

func GetTwitterAuthURL(ctx iris.Context, sessionsManager *sessions.Sessions) (string, error) {
	provider, err := goth.GetProvider("twitter")
	if err != nil {
		return "", err
	}

	sess, err := provider.BeginAuth(SetState(ctx))
	if err != nil {
		return "", err
	}

	url, err := sess.GetAuthURL()
	if err != nil {
		return "", err
	}
	session := sessionsManager.Start(ctx)
	session.Set("twitter", sess.Marshal())
	return url, nil
}

// Checks if the "state" request parameter exists.
// If not, it is set.
// The "state" parameter is used to protected against CSRF in OAuth.
var SetState = func(ctx iris.Context) string {
	state := ctx.URLParam("state")
	if len(state) > 0 {
		return state
	}

	return "state"
}

// GetState gets the state returned by the provider during the callback.
// This is used to prevent CSRF attacks, see
// http://tools.ietf.org/html/rfc6749#section-10.12
var GetState = func(ctx iris.Context) string {
	return ctx.URLParam("state")
}