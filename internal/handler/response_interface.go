package handler

import "github.com/gin-gonic/gin"

type ContentType string

const (
	CONTENT_TYPE_JSON ContentType = "application/json"
	CONTENT_TYPE_XML  ContentType = "application/xml"
	CONTENT_TYPE_M3U8 ContentType = "application/vnd.apple.mpegurl"
	CONTENT_TYPE_TEXT ContentType = "text/plain"
)

// ResponseController interface defines the methods that a response controller should implement.
type ResponseController interface {
	// IsAbort checks if the request should be aborted based on the status code or errors.
	IsAbort() bool

	// GetStatusCode returns the HTTP status code.
	GetStatusCode() int

	// GetErrors returns a slice of errors, if any.
	GetErrors() []error

	// GetResponse returns the response object.
	GetResponse() interface{}

	// SetStatusCode sets the HTTP status code.
	SetStatusCode(statusCode int)

	// SetErrors sets the slice of errors.
	SetErrors(errors []error)

	// AddError appends a single error to the errors slice.
	AddError(err error)

	// SetResponse sets the response object.
	SetResponse(response interface{})

	// SetResult is a shortcut to set the status code and response object.
	SetResult(status int, response interface{})

	// Write writes the response to the Gin context with the appropriate status code.
	Write(ctx *gin.Context)

	// SetContentType sets the content type of the response.
	SetContentType(contentType ContentType)
}

// baseResponseController is a concrete implementation of the ResponseController interface.
type baseResponseController struct {
	contentType ContentType
	statusCode  int
	err         []error
	response    interface{}
}

func (b *baseResponseController) IsAbort() bool {
	return b.statusCode >= 400 || len(b.err) > 0
}

func (b *baseResponseController) GetStatusCode() int {
	return b.statusCode
}

func (b *baseResponseController) GetErrors() []error {
	return b.err
}

func (b *baseResponseController) GetResponse() interface{} {
	return b.response
}

func (b *baseResponseController) SetStatusCode(statusCode int) {
	b.statusCode = statusCode
}

func (b *baseResponseController) SetErrors(errors []error) {
	b.err = []error{}

	for _, err := range errors {
		b.err = append(b.err, err)
	}
}

func (b *baseResponseController) AddError(err error) {
	b.err = append(b.err, err)
}

func (b *baseResponseController) SetResponse(response interface{}) {
	b.response = response
}

func (b *baseResponseController) SetResult(status int, response interface{}) {
	b.statusCode = status
	b.response = response
}

func (b *baseResponseController) Write(ctx *gin.Context) {
	ctx.Header("Content-Type", string(b.contentType))
	switch b.contentType {
	case CONTENT_TYPE_M3U8:
		ctx.Writer.Write([]byte(b.response.(string)))
	case CONTENT_TYPE_TEXT:
		ctx.String(b.statusCode, b.response.(string))
	default:
		ctx.JSON(b.statusCode, b.response)
	}
}

func (b *baseResponseController) SetContentType(contentType ContentType) {
	b.contentType = contentType
}

// NewJsonResponseController creates a new instance of baseResponseController and returns it as a ResponseController.
func NewJsonResponseController() ResponseController {
	return &baseResponseController{
		contentType: CONTENT_TYPE_JSON,
	}
}
