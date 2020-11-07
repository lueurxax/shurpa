package applemusic

import "github.com/lueurxax/shurpa/models"

type SongFinder interface {
	SearchSong(info *models.SongInfo) (link string, err error)
}

type finder struct {
}

func (f *finder) SearchSong(info *models.SongInfo) (link string, err error) {
	panic("implement me")
}

func NewFinder() SongFinder {
	return &finder{}
}
