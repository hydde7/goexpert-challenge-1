package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Controller is an interface that defines the methods required to handle requests.
type Controller interface {
	// SetParams sets the parameters of the request.
	SetParams(params gin.Params)

	// SetHost sets the host of the request.
	SetHost(host string)

	// SetPath sets the path of the request.
	SetPath(path string)

	// SetHeader sets the header of the request.
	SetHeader(header http.Header)

	// SetContext sets the context of the request.
	SetContext(context *gin.Context)

	// SetDefaultLog sets the default log entry for the request.
	SetDefaultLog(log *logrus.Entry)

	// GetHost returns the host of the request.
	GetHost() string

	// GetParam returns the value of the specified parameter.
	GetParam(key string) string

	// GetQueryParam returns the value of the specified query parameter.
	GetQueryParam(key string) string

	// GetClientIP returns the IP address of the client.
	GetClientIP() string

	// Execute executes the request with the given payload and returns a response.
	Execute(payload interface{}) ResponseController
}

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
