package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_greeting(t *testing.T) {
	t.Parallel()

	type args struct {
		name string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "success",
			args: args{
				name: "John Smith",
			},
			want: "Hello, John Smith!",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := greeting(tt.args.name)
			require.Equal(t, tt.want, got)
		})
	}
}
