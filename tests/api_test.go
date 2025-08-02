
package api

import (
    "net/http"
    "net/http/httptest"
    "testing"
)

func TestHealthCheck(t *testing.T) {
    req := httptest.NewRequest("GET", "/health", nil)
    w := httptest.NewRecorder()
    HealthCheck(w, req)

    if w.Code != http.StatusOK {
        t.Errorf("Expected 200 OK but got %d", w.Code)
    }
}
