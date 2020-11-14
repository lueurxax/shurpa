package applemusic

import (
	"context"

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
	r.APIClient.GetSongInfo(ctx, "")
	return nil, nil
}

// NewResolver construct new LinkResolver interface
func NewResolver(client APIClient) LinkResolver {
	return &resolver{
		APIClient: client,
	}
}
