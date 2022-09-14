package health

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	logrus.WithField("requestUri", r.RequestURI).Debug("health")
	fmt.Fprint(w, "OK")
}
