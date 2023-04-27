package main

import (
	"context"
	"github.com/1-samuel/hoot-cal/match"
	"github.com/1-samuel/hoot-cal/mongodb"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"net/http"
	"os"
)

func main() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	database := client.Database("owl")

	repo := mongodb.NewRepositoryMongo(database)
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
