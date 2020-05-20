package ucmsgrecv

import (
	"testing"
)

func Test_parser(t *testing.T) {
	type args struct {
		text string
	}
	tests := []struct {
		name    string
		args    args
		wantTo  string
		wantMsg string
		wantErr bool
	}{
		{
			name: "simple good string",
			args: args{
				text: "@username blah blah blah",
			},
			wantErr: false,
			wantMsg: "blah blah blah",
			wantTo:  "username",
		},
		{
			name: "good string with left spaces",
			args: args{
				text: "  @username blah blah blah",
			},
			wantErr: false,
			wantMsg: "blah blah blah",
			wantTo:  "username",
		},
		{
			name: "good string with right spaces",
			args: args{
				text: "@username blah blah blah  ",
			},
			wantErr: false,
			wantMsg: "blah blah blah",
			wantTo:  "username",
		},
		{
			name: "good string with left new line",
			args: args{
				text: `  
				@username blah blah blah`,
			},
			wantErr: false,
			wantMsg: "blah blah blah",
			wantTo:  "username",
		},
		{
			name: "bad string without @ symbol",
			args: args{
				text: "username blah blah blah",
			},
			wantErr: true,
			wantMsg: "",
			wantTo:  "",
		},
		{
			name: "bad string without message",
			args: args{
				text: "@username",
			},
			wantErr: true,
			wantMsg: "",
			wantTo:  "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTo, gotMsg, err := parser(tt.args.text)
			if (err != nil) != tt.wantErr {
				t.Errorf("parser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotTo != tt.wantTo {
				t.Errorf("parser() gotTo = %v, want %v", gotTo, tt.wantTo)
			}
			if gotMsg != tt.wantMsg {
				t.Errorf("parser() gotMsg = %v, want %v", gotMsg, tt.wantMsg)
			}
		})
	}
}
