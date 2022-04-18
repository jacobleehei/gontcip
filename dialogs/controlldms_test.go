package dialogs

import (
	"reflect"
	"testing"
	"time"

	"github.com/gosnmp/gosnmp"
)

// for test
// dms: &gosnmp.GoSNMP{
// 					Target:    "192.168.9.76",
// 					Port:      161,
// 					Community: "public",
// 					Timeout:   2 * time.Second,
// 					Retries:   1,
// 					Version:   gosnmp.Version1,
// 					MaxOids:   500,
// 				},

func TestActivatingMessage(t *testing.T) {
	type args struct {
		dms               *gosnmp.GoSNMP
		duration          int
		priority          int
		messageMemoryType int
		messageNumber     int
	}
	tests := []struct {
		name        string
		args        args
		wantResults []string
		wantErr     bool
	}{
		// TODO: Add test cases.
		{
			name: "test_1",
			args: args{
				dms: &gosnmp.GoSNMP{
					Target:    "192.168.9.76",
					Port:      161,
					Community: "public",
					Timeout:   2 * time.Second,
					Retries:   1,
					Version:   gosnmp.Version1,
					MaxOids:   500,
				},
				duration:          65535,
				priority:          255,
				messageMemoryType: 3,
				messageNumber:     5,
			},
			wantResults: []string{},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResults, err := ActivatingMessage(tt.args.dms, tt.args.duration, tt.args.priority, tt.args.messageMemoryType, tt.args.messageNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("ActivatingMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResults, tt.wantResults) {
				t.Errorf("ActivatingMessage() = %v, want %v", gotResults, tt.wantResults)
			}
		})
	}
}
