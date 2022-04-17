package dialogs

import (
	"reflect"
	"testing"
	"time"

	"github.com/gosnmp/gosnmp"
)

func TestActivatingMessage(t *testing.T) {
	type args struct {
		dms                      *gosnmp.GoSNMP
		dmsActivateMessageStruct []byte
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
				dmsActivateMessageStruct: []byte{0x01, 0x0B, 0x37, 0x04, 0x00, 0x05, 0x95, 0xF9, 0x67, 0x08, 0x09, 0x0A},
			},
			wantResults: []string{},
			wantErr:     false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResults, err := ActivatingMessage(tt.args.dms, tt.args.dmsActivateMessageStruct)
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
