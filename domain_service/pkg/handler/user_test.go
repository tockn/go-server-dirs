package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/entity"

	"github.com/stretchr/testify/assert"

	"github.com/tockn/go-dirs/domain_service/pkg/mock"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/repository"
)

func TestHandler_GetUser(t *testing.T) {
	tests := []struct {
		name       string
		reqBody    []byte
		service    repository.User
		expectBody []byte
		expectCode int
	}{
		{
			name:    "success",
			reqBody: []byte(`{"name":"hoge"}`),
			service: &mock.UserService{
				ExpectedUser: user1(),
			},
			expectBody: []byte(`{"id":1,"name":"hoge"}
`),
			expectCode: 200,
		},
		{
			name:       "bad json format",
			reqBody:    []byte(`not json`),
			service:    &mock.UserService{},
			expectCode: 400,
		},
		{
			name:    "internal error",
			reqBody: []byte(`{"name":"hoge"}`),
			service: &mock.UserService{
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
				User: User{tt.service},
			}
			h.Router().ServeHTTP(recorder, req)

			actual := recorder.Body.Bytes()

			assert.Equal(t, tt.expectCode, recorder.Code)
			assert.Equal(t, tt.expectBody, actual)
		})
	}
}

var fixedTime = time.Date(2020, 7, 5, 0, 0, 0, 0, time.UTC)

func user1() *entity.User {
	return &entity.User{
		ID:   1,
		Name: "hoge",
	}
}
