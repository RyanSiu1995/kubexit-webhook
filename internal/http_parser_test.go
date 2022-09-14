package internal

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestToAdmission(t *testing.T) {
	reader, err := os.Open("./fixtures/admission_request.json")
	assert.Nil(t, err, "Can open up the test JSON")

	req := httptest.NewRequest(http.MethodPost, "/mutation/kubexit", reader)
	out, err := RequestToAdmission(req)
	assert.Nil(t, err, "No error out from RequestToAdmission")
	assert.NotNil(t, out, "Get the admission request out")
}

func TestRequestToAdmissionWithNoBody(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/mutation/kubexit", nil)
	out, err := RequestToAdmission(req)
	assert.Equal(t, fmt.Errorf("Admission request body is empty"), err, "Error out from RequestToAdmission")
	assert.Nil(t, out, "No out from RequestToAdmission")
}

func TestRequestToAdmissionWithBrokenBody(t *testing.T) {
	reader, err := os.Open("./fixtures/broken.json")
	assert.Nil(t, err, "Can open up the test JSON")
	req := httptest.NewRequest(http.MethodPost, "/mutation/kubexit", reader)

	out, err := RequestToAdmission(req)
	assert.True(t, strings.HasPrefix(err.Error(), "Could not parse admission request: "), "Error out from RequestToAdmission")
	assert.Nil(t, out, "No out from RequestToAdmission")
}

func TestRequestToAdmissionWithNoRequestField(t *testing.T) {
	reader, err := os.Open("./fixtures/invalid.json")
	assert.Nil(t, err, "Can open up the test JSON")
	req := httptest.NewRequest(http.MethodPost, "/mutation/kubexit", reader)

	out, err := RequestToAdmission(req)
	assert.Equal(t, fmt.Errorf("Request field is nil"), err, "Error out from RequestToAdmission")
	assert.Nil(t, out, "No out from RequestToAdmission")
}
