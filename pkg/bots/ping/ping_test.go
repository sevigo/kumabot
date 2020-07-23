package ping

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_pingBot_Match(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want bool
	}{
		{
			name: "case 1",
			in:   "Ping",
			want: true,
		},
		{
			name: "case 2",
			in:   "ping",
			want: true,
		},
		{
			name: "case 3",
			in:   "test",
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := New()
			got := p.Match(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}

func Test_pingBot_Run(t *testing.T) {
	res := New().Run("ping")
	assert.Equal(t, "pong", res)
}
