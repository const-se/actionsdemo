package main

import (
	"testing"

	"github.com/stretchr/testify/require"
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

func Test_name(t *testing.T) {
	t.Parallel()

	type args struct {
		args []string
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "custom",
			args: args{
				args: []string{"app", "John Smith"},
			},
			want: "John Smith",
		},
		{
			name: "default",
			args: args{
				args: []string{"app"},
			},
			want: "John Doe",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got := name(tt.args.args)
			require.Equal(t, tt.want, got)
		})
	}
}
