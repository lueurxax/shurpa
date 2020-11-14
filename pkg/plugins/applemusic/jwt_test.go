package applemusic

import (
	"io/ioutil"
	"testing"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/kelseyhightower/envconfig"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/lueurxax/shurpa/common/crypto"
)

type testConfig struct {
	KeyPath string `envconfig:"KEY" required:"true"`
	Issuer  string `envconfig:"ISSUER" required:"true"`
	KID     string `envconfig:"KID" required:"true"`
}

func Test_jwt_CreateToken(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping testing in short mode")
	}
	t.Run("create token", func(t *testing.T) {
		var cfg testConfig
		if err := envconfig.Process("", &cfg); err != nil {
			require.NoError(t, err)
		}
		key, err := ioutil.ReadFile(cfg.KeyPath)
		require.NoError(t, err)
		j, err := NewJwt(key, cfg.KID, cfg.Issuer, time.Hour*24)
		require.NoError(t, err)
		tokenString, err := j.CreateToken()
		require.NoError(t, err)
		ecKey, err := crypto.ParsePKCS8PrivateKeyFromPEM(key)
		require.NoError(t, err)
		token, err := jwtgo.Parse(tokenString, func(token *jwtgo.Token) (interface{}, error) {
			return ecKey.Public(), nil
		})
		require.NoError(t, err)
		assert.Equal(t, cfg.KID, token.Header[kidField])
	})
}
