package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCP437Plus(t *testing.T) {
	tests := []struct {
		name string
		x    string
		y    string
		want string
	}{
		{
			name: "plus empty",
			x:    cp437Empty,
			y:    cp437Vertical,
			want: cp437Vertical,
		},
		{
			name: "corner top right + corner top left",
			x:    cp437BottomVerticalLeft,
			y:    cp437BottomVerticalRight,
			want: cp437BottomVerticalLeftRight,
		},
		{
			name: "corner bottom right + corner bottom left",
			x:    cp437TopVerticalLeft,
			y:    cp437TopVerticalRight,
			want: cp437TopVerticalLeftRight,
		},
		{
			name: "corner top right + corner top left + corner bottom right",
			x:    cp437BottomVerticalLeftRight,
			y:    cp437TopVerticalLeft,
			want: cp437VerticalLeftRight,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := cp437Plus(tc.x, tc.y)
			assert.Equal(t, tc.want, got)
		})
	}
}
