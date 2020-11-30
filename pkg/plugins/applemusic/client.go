package applemusic

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	applemusic "github.com/minchao/go-apple-music"

	"github.com/lueurxax/shurpa/models"
)

const storefront = "ru"
const limit = 25

type APIClient interface {
	Init() error
	GetSongInfo(ctx context.Context, id string) ([]applemusic.Song, error)
	SearchSong(ctx context.Context, songInfo *models.SongInfo) (link string, err error)
}

type tokenCreator interface {
	CreateToken() (token string, err error)
}

type client struct {
	tokenCreator
	storefront string
	client     *applemusic.Client
}

type Song struct {
	Attributes struct {
		Name string `json:"name"`
	} `json:"attributes"`
	Href string `json:"href"`
}

func (c *client) SearchSong(ctx context.Context, songInfo *models.SongInfo) (string, error) {
	for offset := 0; ; offset += limit {
		artists, err := c.searchArtist(ctx, songInfo.Artist, offset)
		if err != nil {
			return "", err
		}
		if len(artists) == 0 {
			return "", errors.New("not found")
		}
		for _, artist := range artists {
			// search album in
			album, err := c.findAlbum(ctx, artist, songInfo.Album)
			if err != nil {
				return "", err
			}
			if album == nil {
				continue
			}
			song, err := c.findSong(album, songInfo.Name)
			if err != nil {
				return "", err
			}
			if song == nil {
				continue
			}
			return song.Href, nil
		}
	}
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

func (c *client) searchSong(ctx context.Context, songName string, offset int) ([]applemusic.Song, error) {
	search, _, err := c.client.Catalog.Search(ctx, c.storefront, &applemusic.SearchOptions{
		Term:   strings.ReplaceAll(songName, " ", "+"),
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}
	if search != nil && search.Results.Songs != nil {
		return search.Results.Songs.Data, nil
	}
	return nil, nil
}

func (c *client) searchArtist(ctx context.Context, artist string, offset int) ([]applemusic.Artist, error) {
	search, _, err := c.client.Catalog.Search(ctx, c.storefront, &applemusic.SearchOptions{
		Term:   strings.ReplaceAll(artist, " ", "+"),
		Limit:  limit,
		Offset: offset,
		Types:  "artists",
	})
	if err != nil {
		return nil, err
	}
	if search != nil && search.Results.Artists != nil {
		return search.Results.Artists.Data, nil
	}
	return nil, nil
}

func (c *client) findAlbum(ctx context.Context, artist applemusic.Artist, albumName string) (*applemusic.Album, error) {
	for _, album := range artist.Relationships.Albums.Data {
		albumInfo, _, err := c.client.Catalog.GetAlbum(ctx, c.storefront, album.Id, nil)
		if err != nil {
			return nil, err
		}
		for _, albumDetail := range albumInfo.Data {
			if albumDetail.Attributes.Name == albumName {
				return &albumDetail, nil
			}
		}
	}
	return nil, nil
}

func (c *client) findSong(album *applemusic.Album, songName string) (*Song, error) {
	for _, songData := range album.Relationships.Tracks.Data {
		song := new(Song)
		if err := json.Unmarshal(songData.RawMessage, song); err != nil {
			return nil, err
		}
		if song.Attributes.Name == songName {
			return song, nil
		}
	}
	return nil, nil
}

func NewClient(tokenCreator tokenCreator) APIClient {
	return &client{tokenCreator: tokenCreator, storefront: storefront}
}
