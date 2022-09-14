package main

import (
	  "net/http"
    "os"

    "github.com/sirupsen/logrus"
    "github.com/RyanSiu1995/kubexit-webhook/pkg/health"
)

func main() {
    http.HandleFunc("/health", health.HealthCheckHandler)

    cert, certOk := os.LookupEnv("SERVER_TLS_CERT")
    key, keyOk := os.LookupEnv("SERVER_TLS_KEY")
    if certOk && keyOk {
		    logrus.Print("Listening on port 443...")
		    logrus.Fatal(http.ListenAndServeTLS(":443", cert, key, nil))
    } else {
		    logrus.Print("Listening on port 8080...")
		    logrus.Fatal(http.ListenAndServe(":8080", nil))
    }
}
