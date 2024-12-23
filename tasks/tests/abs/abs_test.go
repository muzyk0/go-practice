package abs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbs(t *testing.T) {
	tests := []struct {
		name  string
		value float64
		want  float64
	}{
		{
			name:  "with negative value",
			value: -3,
			want:  3,
		}, {
			name:  "with positive value",
			value: 3,
			want:  3,
		}, {
			name:  "with negative float",
			value: -2.000001,
			want:  2.000001,
		}, {
			name:  "with negatie zero float",
			value: -0.000000003,
			want:  0.000000003,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.want, Abs(test.value))

			if result := Abs(test.value); result != test.want {
				t.Errorf("Abs() = %f, want = %f", result, test.want)

			}

		})
	}
}
