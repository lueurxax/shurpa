package applemusic

import (
	"context"

	applemusic "github.com/minchao/go-apple-music"
)

const storefront = "ru"

type APIClient interface {
	Init() error
	GetSongInfo(ctx context.Context, id string) ([]applemusic.Song, error)
}

type tokenCreator interface {
	CreateToken() (token string, err error)
}

type client struct {
	tokenCreator
	storefront string
	client     *applemusic.Client
}

func (c *client) Init() error {
	token, err := c.tokenCreator.CreateToken()
	if err != nil {
		return err
	}
	tp := applemusic.Transport{Token: token}
	c.client = applemusic.NewClient(tp.Client())
	return nil
}

func (c *client) GetSongInfo(ctx context.Context, id string) ([]applemusic.Song, error) {
	songs, _, err := c.client.Catalog.GetSong(ctx, c.storefront, id, nil)
	if err != nil {
		return nil, err
	}
	return songs.Data, nil
}

func NewClient(tokenCreator tokenCreator) APIClient {
	return &client{tokenCreator: tokenCreator, storefront: "ru"}
}
