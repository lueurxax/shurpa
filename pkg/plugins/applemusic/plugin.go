package applemusic

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

func NewPlugin(id string, host string) Plugin {
	return &plugin{id: id, host: host, SongFinder: NewFinder(), LinkResolver: NewResolver()}
}
