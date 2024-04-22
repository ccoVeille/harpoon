package syscalls

import (
	"bytes"
	"testing"
)

func TestPrint(t *testing.T) {
	type args struct {
		syscalls []uint32
	}
	tests := []struct {
		name       string
		args       args
		wantWriter string
		wantErr    bool
	}{
		// TODO: Add test cases.
		{
			name: "should_print_without_errors",
			args: args{
				syscalls: []uint32{0, 1, 2, 3},
			},
			wantWriter: "read\nwrite\nopen\nclose\n",
			wantErr:    false,
		},
		{
			name: "should_return_error",
			args: args{
				syscalls: []uint32{0, 1, 2, 99999},
			},
			wantWriter: "read\nwrite\nopen\n",
			wantErr:    true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			writer := &bytes.Buffer{}
			if err := Print(writer, tt.args.syscalls); (err != nil) != tt.wantErr {
				t.Errorf("Print() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotWriter := writer.String(); gotWriter != tt.wantWriter {
				t.Errorf("Print() = %v, want %v", gotWriter, tt.wantWriter)
			}
		})
	}
}
