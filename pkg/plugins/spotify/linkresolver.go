package spotify

import (
	"github.com/lueurxax/shurpa/models"
	"github.com/zmb3/spotify"
	"strings"
)

type LinkResolver interface {
	ResolveLink(link string) (info *models.SongInfo, err error)
}

type resolver struct {
	client *spotify.Client
}

func getIdFromLink(link string) (string, error) {
	if strings.HasPrefix(link, "http://open.spotify.com/track/") {
		return strings.TrimPrefix(link, "http://open.spotify.com/track/"), nil
	}
	if strings.HasPrefix(link, "spotify:track:") {
		return strings.TrimPrefix(link, "spotify:track:"), nil
	}
	return link, nil
}

func (r *resolver) ResolveLink(link string) (info *models.SongInfo, err error) {
	id, err := getIdFromLink(link)
	if err != nil {
		return nil, err
	}
	res, err := r.client.GetTrack(spotify.ID(id))
	if err != nil {
		return nil, err
	}
	return &models.SongInfo{
		Name:   res.Name,
		Album:  res.Album.Name,
		Artist: res.Artists[0].Name,
		Other:  nil,
	}, nil

}

func NewResolver(client *spotify.Client) LinkResolver {
	return &resolver{
		client: client,
	}
}
