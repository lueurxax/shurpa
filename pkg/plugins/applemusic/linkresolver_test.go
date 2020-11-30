package applemusic

import "testing"

func Test_resolver_getID(t *testing.T) {
	type args struct {
		link string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "test simple link",
			args: args{link: "https://music.apple.com/ru/album/soft/1209133183"},
			want: "1209133183",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &resolver{}
			got, err := r.getID(tt.args.link)
			if (err != nil) != tt.wantErr {
				t.Errorf("getID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
