package gontcip

import (
	"fmt"
	"testing"
)

func Test_formatShortErrorStatusParameter(t *testing.T) {
	type args struct {
		getResult interface{}
	}
	tests := []struct {
		name       string
		args       args
		wantResult []string
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "test_1",
			args: args{
				getResult: 8258,
			},
			wantResult: []string{"Message Error", "Invalid"},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResult, err := formatShortErrorStatusParameter(tt.args.getResult)
			if (err != nil) != tt.wantErr {
				t.Errorf("formatShortErrorStatusParameter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if fmt.Sprint(gotResult) != fmt.Sprint(tt.wantResult) {
				t.Errorf("formatShortErrorStatusParameter() = %v, want %v", gotResult, tt.wantResult)
			}
		})
	}
}
