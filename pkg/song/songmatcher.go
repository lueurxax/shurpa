package song

import "github.com/lueurxax/shurpa/models"

// Matcher match song in source service with song in destination service by link
type Matcher interface {
	MatchSong(link string, destination string) (response string, err error)
}

type plugin interface {
	linkResolver
	songFinder
}

type linkResolver interface {
	ResolveLink(link string) (info models.SongInfo, err error)
}

type songFinder interface {
	SearchSong(info *models.SongInfo) (link string, err error)
}

type pluginMatcher interface {
	MatchPlugin(link string) (pluginIdentifier string, err error)
}

type matcher struct {
	plugins map[string]plugin
	pluginMatcher
}

func (m *matcher) MatchSong(link string, destination string) (string, error) {
	sourcePlugin, err := m.pluginMatcher.MatchPlugin(link)
	if err != nil {
		return "", err
	}
	info, err := m.plugins[sourcePlugin].ResolveLink(link)
	if err != nil {
		return "", err
	}
	return m.plugins[destination].SearchSong(&info)
}

// NewMatcher construct new Matcher interface
func NewMatcher(pluginMatcher pluginMatcher) Matcher {
	return &matcher{pluginMatcher: pluginMatcher}
}
