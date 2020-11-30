package applemusic

import (
	"context"
	"io/ioutil"
	"log"
	"testing"
	"time"

	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/require"

	"github.com/lueurxax/shurpa/models"
)

const (
	getInfoSuccess = "get song info success"
	getLinkSuccess = "get song link success"
)

func Test_client_GetSongInfo(t *testing.T) {
	var cfg testConfig
	if err := envconfig.Process("", &cfg); err != nil {
		require.NoError(t, err)
	}
	key, err := ioutil.ReadFile(cfg.KeyPath)
	require.NoError(t, err)
	j, err := NewJwt(key, cfg.KID, cfg.Issuer, time.Hour)
	require.NoError(t, err)
	t.Run(getInfoSuccess, func(t *testing.T) {
		c := NewClient(j)
		require.NoError(t, c.Init())

		got, err := c.GetSongInfo(context.Background(), "1524378510")
		require.NoError(t, err)
		log.Print(got[0].Attributes.Name)
	})
	t.Run(getLinkSuccess, func(t *testing.T) {
		c := NewClient(j)
		require.NoError(t, c.Init())

		got, err := c.SearchSong(context.Background(), &models.SongInfo{
			Name:   "Breath",
			Album:  "Shallow Bay: The Best of Breaking Benjamin (Deluxe Edition)",
			Artist: "Breaking Benjamin",
		})
		require.NoError(t, err)
		log.Print(got)
	})
}
