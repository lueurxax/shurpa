package models

// SongInfo contains meta information about song for
// search this song in other services
type SongInfo struct {
	Name   string
	Album  string
	Artist string
	Other  map[string]string
}
