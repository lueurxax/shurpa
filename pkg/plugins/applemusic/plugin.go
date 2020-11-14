package applemusic

// Plugin contain all apple music functional for searching music
type Plugin interface {
	SongFinder
	LinkResolver
	ID() string
	Host() string
}

type plugin struct {
	SongFinder
	LinkResolver
	id, host string
}

func (p *plugin) ID() string {
	return p.id
}

func (p *plugin) Host() string {
	return p.host
}

// NewPlugin construct new plugin
func NewPlugin(id string, host string) Plugin {
	return &plugin{id: id, host: host, SongFinder: NewFinder(), LinkResolver: NewResolver()}
}
