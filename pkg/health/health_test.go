package health

import (
    "io/ioutil"
    "testing"
    "net/http"
    "net/http/httptest"
)

func TestHealthCheck(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/heath", nil)
    w := httptest.NewRecorder()

    HealthCheckHandler(w, req)
    res := w.Result()

    defer res.Body.Close()

    data, err := ioutil.ReadAll(res.Body)
    if err != nil {
        t.Errorf("Expected error to be nil got %v", err)
    }
    if string(data) != "OK" {
        t.Errorf("Expected OK got %v", string(data))
    }
}
