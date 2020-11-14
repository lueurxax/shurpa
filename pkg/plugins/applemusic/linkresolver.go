package applemusic

import "github.com/lueurxax/shurpa/models"

// LinkResolver resolve song info by link
type LinkResolver interface {
	ResolveLink(link string) (info models.SongInfo, err error)
}

type resolver struct {
}

func (r *resolver) ResolveLink(link string) (info models.SongInfo, err error) {
	panic("implement me")
}

// NewResolver construct new LinkResolver interface
func NewResolver() LinkResolver {
	return &resolver{}
}
