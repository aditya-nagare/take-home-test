package log

import (
	"context"
	"os"
	"runtime"

	uuid "github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

var logger *log.Logger

type key string

const (
	//REQUESTID string type key for context
	REQUESTID key = "requestID"
)

//Initialization for logs
func init() {

	logger = log.New()
	logger.SetLevel(log.TraceLevel)
	logger.Formatter = &log.TextFormatter{}
	log.SetOutput(os.Stdout)
}

//Logger with fields
func Logger(ctx context.Context) *log.Entry {
	var depth = 1
	var requestid string

	//Tracking Request Using Context
	if ctxRqID, ok := ctx.Value(REQUESTID).(string); ok {

		requestid = ctxRqID
	}
	function, file, line, _ := runtime.Caller(depth)
	functionObject := runtime.FuncForPC(function)
	entry := logger.WithFields(log.Fields{
		"requestid": requestid,
		"file":      file,
		"function":  functionObject.Name(),
		"line":      line,
	})

	logger.SetOutput(os.Stdout)
	return entry

}

// WithRqID returns a context with request ID or creates new a requestId and assigns to context
func WithRqID(ctx context.Context) context.Context {
	return context.WithValue(ctx, REQUESTID, generateRequestID())
}

func generateRequestID() string {
	requestID := uuid.New()
	return requestID.String()
}
