package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Controller interface defines the contract for transaction handling.
type Controller interface {
	// SetParams sets the path parameters from the request.
	SetParams(params gin.Params)

	// SetHost sets the host information.
	SetHost(host string)

	// SetPath sets the path of the request.
	SetPath(path string)

	// SetHeader sets the headers of the request.
	SetHeader(header http.Header)

	// SetContext sets the Gin context.
	SetContext(context *gin.Context)

	// SetDefaultLog sets the logging mechanism.
	SetDefaultLog(log *logrus.Entry)

	// GetHost retrieves the host information.
	GetHost() string

	// GetParam retrieves a specific path parameter by key.
	GetParam(key string) string

	// GetQueryParam retrieves a specific query parameter by key.
	GetQueryParam(key string) string

	// GetClientIP retrieves the IP address of the client.
	GetClientIP() string

	// Execute performs the actual transaction and returns a ResponseController.
	Execute(payload interface{}) ResponseController
}

// TransactionControllerImpl is a concrete implementation of the TransactionController interface.
type TransactionControllerImpl struct {
	Params     gin.Params
	context    *gin.Context
	host       string
	Path       string
	header     http.Header
	DefaultLog *logrus.Entry
}

func (t *TransactionControllerImpl) SetParams(params gin.Params) {
	t.Params = params
}

func (t *TransactionControllerImpl) SetHost(host string) {
	t.host = host
}

func (t *TransactionControllerImpl) SetPath(path string) {
	t.Path = path
}

func (t *TransactionControllerImpl) SetHeader(header http.Header) {
	t.header = header
}

func (t *TransactionControllerImpl) SetContext(context *gin.Context) {
	t.context = context
}

func (t *TransactionControllerImpl) SetDefaultLog(log *logrus.Entry) {
	t.DefaultLog = log
}

func (t *TransactionControllerImpl) GetHost() string {
	return t.host
}

func (t *TransactionControllerImpl) GetParam(key string) string {
	return t.Params.ByName(key)
}

func (t *TransactionControllerImpl) GetQueryParam(key string) string {
	return t.context.Query(key)
}

func (t *TransactionControllerImpl) GetClientIP() string {
	return t.context.ClientIP()
}
