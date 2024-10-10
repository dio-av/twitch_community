package auth

import (
	"log"
	"os"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"github.com/joho/godotenv"
	"github.com/markbates/goth"
	"github.com/markbates/goth/gothic"
	"github.com/markbates/goth/providers/twitch"
)

const (
	cookieExpiration = "48h"

	isProd = true
)

func NewAuth() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error loading .env file %w", err)
		return
	}

	twitchClientId := os.Getenv("CLIENT_ID")
	twitchClientSecret := os.Getenv("CLIENT_SECRET")

	secureKey := securecookie.GenerateRandomKey(64)
	if secureKey == nil {
		log.Fatalf("error generationg secure cookie key")
		return
	}

	store := sessions.NewCookieStore(secureKey)
	cookieDuration, err := time.ParseDuration(cookieExpiration)
	if err != nil {
		log.Fatalf("error parsing cookie expiration date")
		return
	}

	store.MaxAge(int(cookieDuration.Seconds()))

	store.Options.Path = "/"
	store.Options.HttpOnly = true
	store.Options.Secure = isProd

	gothic.Store = store

	goth.UseProviders(
		twitch.New(twitchClientId, twitchClientSecret, "http://localhost:3000/auth/twitch/callback"),
	)
}
