package internal

import (
    "net/http"
    "bytes"
    "encoding/json"
    "fmt"

	admissionv1 "k8s.io/api/admission/v1"
    "github.com/sirupsen/logrus"
)

func RequestToAdmission(r *http.Request) (*admissionv1.AdmissionReview, error) {
    logrus.Debug("Invoking Request body to Admission")
	bodyBuffer := new(bytes.Buffer)
	bodyBuffer.ReadFrom(r.Body)
	body := bodyBuffer.Bytes()

	if len(body) == 0 {
		return nil, fmt.Errorf("Admission request body is empty")
	}

	var admissionBody admissionv1.AdmissionReview

	if err := json.Unmarshal(body, &admissionBody); err != nil {
		return nil, fmt.Errorf("Could not parse admission request: %v", err)
	}

	if admissionBody.Request == nil {
		return nil, fmt.Errorf("Request field is nil")
	}

	return &admissionBody, nil
}
