package dialogs

import (
	"reflect"
	"testing"
	"time"

	"github.com/gosnmp/gosnmp"
)

var test_dms = &gosnmp.GoSNMP{
	Target:    "192.168.9.76",
	Port:      161,
	Community: "public",
	Timeout:   2 * time.Second,
	Retries:   1,
	Version:   gosnmp.Version1,
	MaxOids:   500,
}

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
		wantResult  activatingMessageResult
	}{
		// TODO: Add test cases.
		{
			name: "test_1",
			args: args{
				dms:               test_dms,
				duration:          65535,
				priority:          255,
				messageMemoryType: 3,
				messageNumber:     1,
			},
			wantErr:    false,
			wantResult: activatingMessageResult{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := ActivatingMessage(tt.args.dms, tt.args.duration, tt.args.priority, tt.args.messageMemoryType, tt.args.messageNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("ActivatingMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("RetrievingMessage() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}

func TestDefiningMessage(t *testing.T) {
	type args struct {
		dms               *gosnmp.GoSNMP
		messageMemoryType int
		messageNumber     int
		mutiString        string
		ownerAddress      string
		priority          int
		beacon            int
		pixelService      int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test_1",
			args: args{
				dms:               test_dms,
				messageMemoryType: 3,
				messageNumber:     1,
				mutiString:        "TESTING[nl]",
				ownerAddress:      "127.0.0.1",
				priority:          255,
				beacon:            0,
				pixelService:      0,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DefiningMessage(tt.args.dms, tt.args.messageMemoryType, tt.args.messageNumber, tt.args.mutiString, tt.args.ownerAddress, tt.args.priority, tt.args.beacon, tt.args.pixelService); (err != nil) != tt.wantErr {
				t.Errorf("DefiningMessage() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRetrievingMessage(t *testing.T) {
	type args struct {
		dms               *gosnmp.GoSNMP
		messageMemoryType int
		messageNumber     int
	}
	tests := []struct {
		name       string
		args       args
		wantResult retrievingResult
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "test_1",
			args: args{
				dms:               test_dms,
				messageMemoryType: 3,
				messageNumber:     1,
			},
			wantResult: retrievingResult{},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := RetrievingMessage(tt.args.dms, tt.args.messageMemoryType, tt.args.messageNumber)
			if (err != nil) != tt.wantErr {
				t.Errorf("RetrievingMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResult, tt.wantResult) {
				t.Errorf("RetrievingMessage() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
