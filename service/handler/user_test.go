package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/tockn/go-dirs/service/mock"
	"github.com/tockn/go-dirs/service/model"

	"github.com/tockn/go-dirs/service/repository"
)

func TestHandler_GetUser(t *testing.T) {
	tests := []struct {
		name       string
		repo       repository.User
		expectBody []byte
		expectCode int
	}{
		{
			name: "success",
			repo: &mock.UserRepository{
				ExpectedUser: modelUser1(),
			},
			expectBody: []byte(`{"id":"1","name":"hoge"}
`),
			expectCode: 200,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodGet, "/users/1", nil)
			if err != nil {
				t.Fatal(err)
			}

			h := &Handler{
				userRepository: tt.repo,
			}
			h.Router().ServeHTTP(recorder, req)

			actual := recorder.Body.Bytes()

			assert.Equal(t, tt.expectCode, recorder.Code)
			assert.Equal(t, tt.expectBody, actual)
		})
	}
}

var fixedTime = time.Date(2020, 7, 5, 0, 0, 0, 0, time.UTC)

func modelUser1() *model.User {
	return &model.User{
		ID:        "1",
		Name:      "hoge",
		CreatedAt: fixedTime,
		UpdatedAt: fixedTime,
	}
}
