package scanner

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScan(t *testing.T) {
	type args struct {
		dest any
		val  any
	}
	tests := []struct {
		name    string
		args    args
		want    func() any
		wantErr error
	}{
		{
			name: "on success will return nil",
			args: args{
				dest: new(string),
				val:  "value",
			},
			want: func() any {
				res := "value"
				return &res
			},
			wantErr: nil,
		},
		{
			name: "on success when scanning int value to float64 destination return nil",
			args: args{
				dest: new(float64),
				val:  int(1),
			},
			want: func() any {
				res := float64(1)
				return &res
			},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		err := Scan(tt.args.dest, tt.args.val)
		assert.Equal(t, tt.wantErr, err)
		assert.Equal(t, tt.want(), tt.args.dest)
	}
}
