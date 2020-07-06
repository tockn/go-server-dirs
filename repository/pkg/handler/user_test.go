package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/tockn/go-dirs/repository/pkg/mock"
	"github.com/tockn/go-dirs/repository/pkg/model"

	"github.com/tockn/go-dirs/repository/pkg/repository"
)

func TestHandler_GetUser(t *testing.T) {
	tests := []struct {
		name       string
		reqBody    []byte
		repo       repository.User
		expectBody []byte
		expectCode int
	}{
		{
			name:    "success",
			reqBody: []byte(`{"name":"hoge"}`),
			repo: &mock.UserRepository{
				ExpectedUser: modelUser1(),
			},
			expectBody: []byte(`{"id":1,"name":"hoge"}
`),
			expectCode: 200,
		},
		{
			name:       "bad json format",
			reqBody:    []byte(`not json`),
			repo:       &mock.UserRepository{},
			expectCode: 400,
		},
		{
			name:    "internal error",
			reqBody: []byte(`{"name":"hoge"}`),
			repo: &mock.UserRepository{
				ExpectedError: errors.New("error"),
			},
			expectCode: 500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/users", bytes.NewBuffer(tt.reqBody))
			if err != nil {
				t.Fatal(err)
			}

			h := &Handler{
				UserRepository: tt.repo,
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
		ID:   1,
		Name: "hoge",
	}
}
