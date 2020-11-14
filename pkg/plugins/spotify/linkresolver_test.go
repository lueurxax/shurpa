package spotify

import (
	"github.com/lueurxax/shurpa/models"
	"github.com/zmb3/spotify"
	"reflect"
	"testing"
)

func Test_resolver_ResolveLink(t *testing.T) {
	client := mustNewClient()
	type fields struct {
		client *spotify.Client
	}
	type args struct {
		link string
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantInfo *models.SongInfo
		wantErr  bool
	}{
		{
			name:   "full link",
			fields: fields{client: client},
			args: args{
				link: "http://open.spotify.com/track/6rqhFgbbKwnb9MLmUQDhG6",
			},
			wantInfo: &models.SongInfo{
				Name:   "Speak To Me - 2011 Remastered Version",
				Album:  "The Dark Side Of The Moon (2011 Remastered Version)",
				Artist: "Pink Floyd",
				Other:  nil,
			},
			wantErr: false,
		},
		{
			name:   "invalid id",
			fields: fields{client: client},
			args: args{
				link: "http://open.spotify.com/track/выапролдж",
			},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:   "invalid url",
			fields: fields{client: client},
			args: args{
				link: "dsfghjkl;",
			},
			wantInfo: nil,
			wantErr:  true,
		},
		{
			name:   "track id",
			fields: fields{client: client},
			args: args{
				link: "spotify:track:6rqhFgbbKwnb9MLmUQDhG6",
			},
			wantInfo: &models.SongInfo{
				Name:   "Speak To Me - 2011 Remastered Version",
				Album:  "The Dark Side Of The Moon (2011 Remastered Version)",
				Artist: "Pink Floyd",
				Other:  nil,
			},
			wantErr: false,
		},
		{
			name:   "invalid track id",
			fields: fields{client: client},
			args: args{
				link: "spotify:track:фыавпрол",
			},
			wantInfo: nil,
			wantErr:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &resolver{
				client: tt.fields.client,
			}
			gotInfo, err := r.ResolveLink(tt.args.link)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResolveLink() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInfo, tt.wantInfo) {
				t.Errorf("ResolveLink() gotInfo = %v, want %v", gotInfo, tt.wantInfo)
			}
		})
	}
}
