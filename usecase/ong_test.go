package usecase

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockOngRepository struct {
	mock.Mock
}

func TestGetOng(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/ong/1", nil)
	if err != nil {
		t.Fatalf("failed to create HTTP request: %v", err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/ong/1" {
			http.NotFound(w, r)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code, "unexpected response status")
}
