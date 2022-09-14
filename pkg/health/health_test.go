package health

import (
    "io/ioutil"
    "testing"
    "net/http"
    "net/http/httptest"

    "github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
    req := httptest.NewRequest(http.MethodGet, "/heath", nil)
    w := httptest.NewRecorder()

    HealthCheckHandler(w, req)
    res := w.Result()

    defer res.Body.Close()

    data, err := ioutil.ReadAll(res.Body)
    assert.Nil(t, err, "Expected no error out")
    assert.Equal(t, "OK", string(data), "Expected OK")
}
