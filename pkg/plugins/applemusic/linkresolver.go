package applemusic

import (
	"context"
	"errors"
	"net/url"
	"strings"

	"github.com/lueurxax/shurpa/models"
)

// LinkResolver resolve song info by link
type LinkResolver interface {
	ResolveLink(link string) (info *models.SongInfo, err error)
}

type resolver struct {
	APIClient
}

func (r *resolver) ResolveLink(link string) (*models.SongInfo, error) {
	ctx := context.Background()
	id, err := r.getID(link)
	songs, err := r.APIClient.GetSongInfo(ctx, id)
	if err != nil {
		return nil, err
	}
	if len(songs) == 0 {
		return nil, errors.New("song not found")
	}
	return &models.SongInfo{
		Name:   songs[0].Attributes.Name,
		Album:  songs[0].Attributes.AlbumName,
		Artist: songs[0].Attributes.ArtistName,
	}, nil
}

func (r *resolver) getID(link string) (string, error) {
	u, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	parts := strings.Split(u.Path, "/")
	if len(parts) != 5 {
		return "", errors.New("unsupported link")
	}
	return parts[4], nil
}

// NewResolver construct new LinkResolver interface
func NewResolver(client APIClient) LinkResolver {
	return &resolver{
		APIClient: client,
	}
}
