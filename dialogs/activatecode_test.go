package dialogs

import (
	"encoding/hex"
	"log"
	"testing"
)

func TestEncodeActivateMessageCode(t *testing.T) {
	type args struct {
		multiString      string
		beacon           int
		pixelService     int
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
				multiString:      "[jp3]TEST [fl]Flashing[/fl]",
				beacon:           0,
				pixelService:     0,
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
			got, err := EncodeActivateMessageCode(tt.args.multiString, tt.args.beacon, tt.args.pixelService, tt.args.messageType, tt.args.duration, tt.args.priority, tt.args.messageNumber, tt.args.requestIPAddress)
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

func Test_calcChecksum(t *testing.T) {
	type args struct {
		multiString  string
		beacon       int
		pixelService int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{
			name: "Normal",
			args: args{
				multiString:  "[g1][g1][fo2][jl]TEST TEST TEST[nl][jl]TEST TEST TEST[nl][jl]TEST TEST TEST[nl][fo6][hc6e2c][hc8a66][hc9032][hc884c][hc4e2d][hc0020][hc6e2c][hc8a66][hc9032][hc884c][hc4e2d]",
				beacon:       1,
				pixelService: 0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcChecksum(tt.args.multiString, tt.args.beacon, tt.args.pixelService); got != tt.want {
				log.Println(got)
			}
		})
	}
}
