package dialogs

import (
	"encoding/hex"
	"testing"
)

func TestEncodeActivateMessageCode(t *testing.T) {
	type args struct {
		mutiString       string
		beacon           int
		pixelservice     int
		messageType      int
		duration         int
		priority         int
		messageNumber    int
		requestIPAddress string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_1",
			args: args{
				mutiString:       "[jp3]TEST [fl]Flashing[/fl]",
				beacon:           0,
				pixelservice:     0,
				messageType:      4,
				duration:         267,
				priority:         55,
				messageNumber:    5,
				requestIPAddress: "103.8.9.10",
			},
			want: "010B3704000595F96708090A",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := EncodeActivateMessageCode(tt.args.mutiString, tt.args.beacon, tt.args.pixelservice, tt.args.messageType, tt.args.duration, tt.args.priority, tt.args.messageNumber, tt.args.requestIPAddress)
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodeActivateMessageCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			wanted, _ := hex.DecodeString(tt.want)
			if string(got) != string(wanted) {
				t.Errorf("EncodeActivateMessageCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
