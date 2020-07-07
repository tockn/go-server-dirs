package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/tockn/go-dirs/domain_service/pkg/mock"

	"github.com/tockn/go-dirs/domain_service/pkg/domain/service"
)

func TestPayment_Pay(t *testing.T) {
	tests := []struct {
		name       string
		reqBody    []byte
		service    service.Payment
		expectBody []byte
		expectCode int
	}{
		{
			name:    "success",
			reqBody: []byte(`{"payerID":1,"payeeID":2,"amount":100}`),
			service: &mock.PaymentService{},
			expectBody: []byte(`{"status":"succeeded"}
`),
			expectCode: 201,
		},
		{
			name:       "bad json format",
			reqBody:    []byte(`not json`),
			service:    &mock.PaymentService{},
			expectCode: 400,
		},
		{
			name:    "pay error",
			reqBody: []byte(`{"payerID":1,"payeeID":2,"amount":100}`),
			service: &mock.PaymentService{
				ExpectedError: errors.New("failed message"),
			},
			expectBody: []byte(`{"status":"failed","failed_message":"failed message"}
`),
			expectCode: 500,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			recorder := httptest.NewRecorder()
			req, err := http.NewRequest(http.MethodPost, "/payments", bytes.NewBuffer(tt.reqBody))
			if err != nil {
				t.Fatal(err)
			}

			h := &Handler{
				Payment: Payment{Service: tt.service},
			}
			h.Router().ServeHTTP(recorder, req)

			actual := recorder.Body.Bytes()

			assert.Equal(t, tt.expectCode, recorder.Code)
			assert.Equal(t, tt.expectBody, actual)
		})
	}
}
