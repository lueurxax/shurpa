package applemusic

import "github.com/lueurxax/shurpa/models"

// SongFinder searching music by song info
type SongFinder interface {
	SearchSong(info *models.SongInfo) (link string, err error)
}

type finder struct {
}

func (f *finder) SearchSong(info *models.SongInfo) (link string, err error) {
	panic("implement me")
}

// NewFinder construct new SongFinder interface
func NewFinder() SongFinder {
	return &finder{}
}
