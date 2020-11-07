package applemusic

import "github.com/lueurxax/shurpa/models"

type LinkResolver interface {
	ResolveLink(link string) (info models.SongInfo, err error)
}

type resolver struct {
}

func (r *resolver) ResolveLink(link string) (info models.SongInfo, err error) {
	panic("implement me")
}

func NewResolver() LinkResolver {
	return &resolver{}
}
