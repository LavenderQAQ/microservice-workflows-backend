package router

import (
	"net/http"

	"github.com/Lavender-QAQ/microservice-workflows-backend/pkg/handler"
	"github.com/go-logr/logr"
)

var RouterLogger logr.Logger

func NewRouter(ip_port string) {
	logger := RouterLogger

	http.HandleFunc("/deploy", handler.DeployHandler)
	http.HandleFunc("/save", handler.SaveHandler)

	logger.WithValues("listening ip address and port", ip_port).Info("Start listen")

	err := http.ListenAndServe(ip_port, nil)
	if err != nil {
		logger.Error(err, "Router listen error")
		return
	}
}
