package models

type Config struct {
	ConsumerKey string `json:"consumerKey"`
	ConsumerSecret string `json:"consumerSecret"`
	CallbackUrl string `json:"callbackUrl"`
	CookieHashKey string `json:"cookieHashKey"`
	CookieBlockKey string `json:"cookieBlockKey"`
}