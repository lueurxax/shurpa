package spotify

import (
	"github.com/lueurxax/shurpa/models"
	"github.com/zmb3/spotify"
	"testing"
)

func Test_finder_SearchSong(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}

	client := mustNewClient()

	type fields struct {
		client *spotify.Client
	}
	type args struct {
		info *models.SongInfo
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantLink string
		wantErr  bool
	}{
		{
			name:   "ok",
			fields: fields{client: client},
			args: args{
				info: &models.SongInfo{
					Name:   "Speak To Me - 2011 Remastered Version",
					Album:  "The Dark Side Of The Moon (2011 Remastered Version)",
					Artist: "Pink Floyd",
					Other:  nil,
				},
			},
			wantLink: "https://api.spotify.com/v1/tracks/4rQYDXfKFikLX4ad674jhg",
			wantErr:  false,
		},
		{
			name:   "not_found",
			fields: fields{client: client},
			args: args{
				info: &models.SongInfo{
					Name:   "dsfghjkl",
					Album:  "asdfgfhjyk",
					Artist: "fsghjykuiokp",
					Other:  nil,
				},
			},
			wantLink: "",
			wantErr:  true,
		},
		{
			name:   "empty",
			fields: fields{client: client},
			args: args{
				info: &models.SongInfo{
					Name:   "",
					Album:  "",
					Artist: "",
					Other:  nil,
				},
			},
			wantLink: "",
			wantErr:  true,
		},
		{
			name:   "null",
			fields: fields{client: client},
			args: args{
				info: nil,
			},
			wantLink: "",
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &finder{
				client: tt.fields.client,
			}
			gotLink, err := f.SearchSong(tt.args.info)
			if (err != nil) != tt.wantErr {
				t.Errorf("SearchSong() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotLink != tt.wantLink {
				t.Errorf("SearchSong() gotLink = %v, want %v", gotLink, tt.wantLink)
			}
		})
	}
}
