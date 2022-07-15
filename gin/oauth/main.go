package main

import (
	"fmt"
	"os"

	dotenv "github.com/joho/godotenv"
	"golang.org/x/oauth2"
)

const (
	authEP  = "https://accounts.google.com/o/oauth2/v2/auth"
	tokenEP = "https://www.googleapis.com/oauth2/v4/token"
)

func main() {
	if err := dotenv.Load(); err != nil {
		panic(err)
	}

	config := getConnect()
	fmt.Println(config.AuthCodeURL(""))
}

func getConnect() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     os.Getenv("clientID"),
		ClientSecret: os.Getenv("clientSecret"),
		Endpoint: oauth2.Endpoint{
			AuthURL:  authEP,
			TokenURL: tokenEP,
		},
		Scopes:      []string{"openid", "email", "profile"},
		RedirectURL: "http://localhost:8080/google/callback",
	}
}
