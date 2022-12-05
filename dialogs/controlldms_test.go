package dialogs

import (
	"log"
	"reflect"
	"testing"
	"time"

	"github.com/gosnmp/gosnmp"
)

var test_dms = &gosnmp.GoSNMP{
	Conn:               nil,
	Target:             "10.0.11.41",
	Port:               1000,
	Transport:          "tcp",
	Community:          "Public",
	Version:            gosnmp.Version1,
	Context:            nil,
	Timeout:            3 * time.Second,
	Retries:            5,
	ExponentialTimeout: false,
	Logger:             gosnmp.Logger{},
	PreSend: func(*gosnmp.GoSNMP) {
	},
	OnSent: func(*gosnmp.GoSNMP) {
	},
	OnRecv: func(*gosnmp.GoSNMP) {
	},
	OnRetry: func(*gosnmp.GoSNMP) {
	},
	OnFinish: func(*gosnmp.GoSNMP) {
	},
	MaxOids:                 500,
	MaxRepetitions:          0,
	NonRepeaters:            0,
	UseUnconnectedUDPSocket: false,
	LocalAddr:               "",
	AppOpts:                 map[string]interface{}{},
	MsgFlags:                0,
	SecurityModel:           0,
	SecurityParameters:      nil,
	ContextEngineID:         "",
	ContextName:             "",
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
				messageNumber:     6,
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
				t.Errorf("ActivatingMessage() = %v, want %v", gotResult, tt.wantResult)
			}
			log.Println(gotResult)
		})
	}
}

func TestDefiningMessage(t *testing.T) {
	type args struct {
		dms               *gosnmp.GoSNMP
		messageMemoryType int
		messageNumber     int
		multiString       string
		ownerAddress      string
		priority          int
		beacon            int
		pixelService      int
	}
	tests := []struct {
		name             string
		args             args
		wantDefineResult definingMessageResult
		wantErr          bool
	}{
		// TODO: Add test cases.
		{
			name: "test_1",
			args: args{
				dms:               test_dms,
				messageMemoryType: 3,
				messageNumber:     3,
				multiString:       "[jl3][nl][jl3][nl][jl3][nl][fo6]",
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
			gotDefineResult, err := DefiningMessage(tt.args.dms, tt.args.messageMemoryType, tt.args.messageNumber, tt.args.multiString, tt.args.ownerAddress, tt.args.priority, tt.args.beacon, tt.args.pixelService)
			if (err != nil) != tt.wantErr {
				t.Errorf("DefiningMessage() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			log.Printf("DefiningMessage() = %v", gotDefineResult)
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
				messageNumber:     6,
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
			log.Println(gotResult)
		})
	}
}

func TestGetSnmp(t *testing.T) {
	t.Run("Get Snmp", func(t *testing.T) {
		test_dms.Connect()
		result, err := test_dms.Get([]string{"1.3.6.1.4.1.1206.4.2.3.6.17"})
		if err != nil {
			t.Errorf(err.Error())
			return
		}
		if result.Error.String() != "NoError" {
			t.Errorf(result.Error.String())
			return
		}
		for _, value := range result.Variables {
			log.Println(value.Name, value.Type.String(), value.Value)
		}
	})

}
