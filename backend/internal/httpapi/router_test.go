package httpapi

import (
	"context"
	"encoding/json"
	"errors"
	"mime"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)
	os.Exit(m.Run())
}

type fakePinger struct {
	err error
}

func (f fakePinger) Ping(context.Context) error {
	return f.err
}

func TestHealth(t *testing.T) {
	r := NewRouter(fakePinger{})

	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Fatalf("status = %d, want %d", w.Code, http.StatusOK)
	}

	mediaType, params, err := mime.ParseMediaType(w.Header().Get("Content-Type"))
	if err != nil {
		t.Fatalf("parse Content-Type %q: %v", w.Header().Get("Content-Type"), err)
	}
	if mediaType != "application/json" {
		t.Fatalf("Content-Type = %q, want application/json", mediaType)
	}

	if got := params["charset"]; !strings.EqualFold(got, "utf-8") {
		t.Errorf("charset = %q, want %q", got, "utf-8")
	}

	var body struct {
		Status string `json:"status"`
	}
	if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
		t.Fatalf("decode body %q: %v", w.Body.String(), err)
	}
	if got, want := body.Status, "ok"; got != want {
		t.Fatalf("status field = %q, want %q", got, want)
	}
}

func TestRoutes(t *testing.T) {
	r := NewRouter(fakePinger{})

	tests := []struct {
		name   string
		method string
		path   string
		want   int
	}{
		{"health", http.MethodGet, "/health", http.StatusOK},
		{"unknown path", http.MethodGet, "/nie-istnieje", http.StatusNotFound},
		{"root", http.MethodGet, "/", http.StatusNotFound},
		{"health with wrong method", http.MethodPost, "/health", http.StatusMethodNotAllowed},
		{"ready", http.MethodGet, "/ready", http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req := httptest.NewRequest(tt.method, tt.path, nil)
			r.ServeHTTP(w, req)

			if w.Code != tt.want {
				t.Errorf("%s %s: status = %d, want %d", tt.method, tt.path, w.Code, tt.want)
			}
		})
	}
}

func TestReady(t *testing.T) {
	tests := []struct {
		name       string
		pingErr    error
		wantCode   int
		wantStatus string
	}{
		{"db up", nil, http.StatusOK, "ready"},
		{"db down", errors.New("connection refused"), http.StatusServiceUnavailable, "unavailable"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := NewRouter(fakePinger{err: tt.pingErr})

			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/ready", nil)
			r.ServeHTTP(w, req)

			if w.Code != tt.wantCode {
				t.Fatalf("status = %d, want %d", w.Code, tt.wantCode)
			}

			var body struct {
				Status string `json:"status"`
			}
			if err := json.Unmarshal(w.Body.Bytes(), &body); err != nil {
				t.Fatalf("decode body %q: %v", w.Body.String(), err)
			}
			if body.Status != tt.wantStatus {
				t.Errorf("status field = %q, want %q", body.Status, tt.wantStatus)
			}
		})
	}
}
