package spotify

import (
	"context"
	"github.com/zmb3/spotify"
	"golang.org/x/oauth2/clientcredentials"
	"log"
	"os"
)

const HOST = "spotify.com"

type Plugin interface {
	SongFinder
	LinkResolver
	ID() string
}

type plugin struct {
	SongFinder
	LinkResolver
	id string
}

func (p *plugin) ID() string {
	return p.id
}

func (p *plugin) Host() string {
	return HOST
}

func NewPlugin(id string) Plugin {
	config := &clientcredentials.Config{
		ClientID:     os.Getenv("SPOTIFY_ID"),
		ClientSecret: os.Getenv("SPOTIFY_SECRET"),
		TokenURL:     spotify.TokenURL,
	}
	token, err := config.Token(context.Background())
	if err != nil {
		log.Fatalf("couldn't get token: %v", err)
	}

	client := spotify.Authenticator{}.NewClient(token)

	return &plugin{id: id, SongFinder: NewFinder(&client), LinkResolver: NewResolver(&client)}
}
