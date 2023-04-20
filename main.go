package main

import (
	"context"
	"github.com/1-samuel/hoot-cal/match"
	"github.com/1-samuel/hoot-cal/owl"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2/clientcredentials"
	"net/http"
	"os"
)

func main() {
	client := configureApiClient()

	repo := owl.NewRepositoryApi(*client)
	resource := match.NewMatchResource(repo)

	r := gin.Default()

	r.GET("/api/v1/matches", resource.GetAll)
	r.GET("/owl.ics", resource.GetCalendar)

	r.Run(":8080")
}

func configureApiClient() *http.Client {
	clientID := os.Getenv("BLIZZARD_CLIENT_ID")
	clientSecret := os.Getenv("BLIZZARD_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		panic("env not set")
	}

	config := clientcredentials.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		TokenURL:     "https://oauth.battle.net/token",
	}

	client := config.Client(context.TODO())
	return client
}
