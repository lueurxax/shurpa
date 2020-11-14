package spotify

import (
	"errors"
	"fmt"
	"github.com/lueurxax/shurpa/models"
	"github.com/zmb3/spotify"
)

var ErrInvalidInfo = errors.New("invalid info")
var ErrNotFound = errors.New("not found")

type SongFinder interface {
	SearchSong(info *models.SongInfo) (link string, err error)
}

type finder struct {
	client *spotify.Client
}

func songInfoToQuery(info *models.SongInfo) string {
	return fmt.Sprintf("%s %s %s", info.Artist, info.Album, info.Name)
}

func (f *finder) SearchSong(info *models.SongInfo) (link string, err error) {
	if info == nil {
		return "", ErrInvalidInfo
	}
	limit := 1
	country := "RU"
	res, err := f.client.SearchOpt(songInfoToQuery(info), spotify.SearchTypeTrack, &spotify.Options{
		Country:   &country,
		Limit:     &limit,
		Offset:    nil,
		Timerange: nil,
	})
	if err != nil {
		return "", err
	}

	if len(res.Tracks.Tracks) == 0 {
		return "", ErrNotFound
	}

	return res.Tracks.Tracks[0].Endpoint, nil
}

func NewFinder(client *spotify.Client) SongFinder {
	return &finder{
		client: client,
	}
}
