package fullname

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFullname(t *testing.T) {
	type fields struct {
		FirstName string
		LastName  string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "simple test",
			fields: fields{
				FirstName: "Vladislav",
				LastName:  "Muzyka",
			},
			want: "Vladislav Muzyka",
		},
		{
			name: "long name",
			fields: fields{
				FirstName: "Vladislav Vladimirovich",
				LastName:  "Muzyka",
			},
			want: "Vladislav Vladimirovich Muzyka",
		},
		{
			name: "Without LastName",
			fields: fields{
				FirstName: "Vladislav ",
				LastName:  "",
			},
			want: "Vladislav",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := User{
				FirstName: tt.fields.FirstName,
				LastName:  tt.fields.LastName,
			}

			got := u.FullName()

			assert.Equal(t, got, tt.want)

			if got != tt.want {
				t.Errorf("FullName() = %v, want %v", got, tt.want)
			}
		})
	}
}
