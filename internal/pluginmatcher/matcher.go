package pluginmatcher

import (
	"fmt"
	"net/url"
)

// Matcher match music service by song link
type Matcher interface {
	MatchPlugin(link string) (pluginIdentifier string, err error)
}

type matcher struct {
	plugins map[string]string
}

// ErrPluginNotFound error plugin not found
type ErrPluginNotFound struct {
	host string
}

func (e *ErrPluginNotFound) Error() string {
	return fmt.Sprintf("host %s not found in the plugins", e.host)
}

func (m *matcher) MatchPlugin(link string) (pluginIdentifier string, err error) {
	host, err := m.getHost(link)
	id, ok := m.plugins[host]
	if !ok {
		return "", &ErrPluginNotFound{host: host}
	}

	return id, nil
}

func (m *matcher) getHost(link string) (string, error) {
	data, err := url.Parse(link)
	if err != nil {
		return "", err
	}
	return data.Host, nil
}

// NewMatcher construct new matcher interface
func NewMatcher(plugins map[string]string) Matcher {
	return &matcher{plugins: plugins}
}
