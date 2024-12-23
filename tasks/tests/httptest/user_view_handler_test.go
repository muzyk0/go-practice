package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var usersDB = map[string]User{
	"u1": {
		ID:        "u1",
		FirstName: "Misha",
		LastName:  "Popov",
	},
	"u2": {
		ID:        "u2",
		FirstName: "Sasha",
		LastName:  "Popov",
	},
}

func TestUserViewHandler(t *testing.T) {
	type want struct {
		code        int
		contentType string
		user        User
	}
	tests := []struct {
		name    string
		request string
		want    want
		usersDB map[string]User
	}{
		{
			name:    "positive test #1",
			request: "/users?user_id=u2",
			want: want{
				code:        200,
				contentType: "application/json",
				user: User{
					ID:        "u2",
					FirstName: "Sasha",
					LastName:  "Popov",
				},
			},
			usersDB: usersDB,
		},
		{
			name:    "with not found",
			request: "/users?user_id=u3",
			want: want{
				code:        404,
				contentType: "application/json",
				user: User{
					ID:        "u2",
					FirstName: "Sasha",
					LastName:  "Popov",
				},
			},
			usersDB: usersDB,
		},
		{
			name:    "user_id is empty",
			request: "/users?",
			want: want{
				code:        400,
				contentType: "application/json",
				user: User{
					ID:        "u2",
					FirstName: "Sasha",
					LastName:  "Popov",
				},
			},
			usersDB: usersDB,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodGet, test.request, nil)
			// создаём новый Recorder
			w := httptest.NewRecorder()
			UserViewHandler(test.usersDB)(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, res.StatusCode, test.want.code)
			// проверяем Content-Type

			// получаем и проверяем тело запроса
			resBody, err := io.ReadAll(res.Body)

			require.NoError(t, err)

			err = res.Body.Close()
			require.NoError(t, err)

			if res.StatusCode == http.StatusOK {
				assert.Equal(t, res.Header.Get("Content-Type"), test.want.contentType)

				var user User
				err = json.Unmarshal(resBody, &user)
				require.NoError(t, err)

				assert.Equal(t, test.want.user, user)
			}

		})
	}
}
