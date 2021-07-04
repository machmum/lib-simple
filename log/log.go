package liblog

import (
	"context"
	"encoding/json"
	"net/http"
	"os"
	"runtime"
	"sync"

	"github.com/labstack/echo"
	libcommon "github.com/machmum/lib-simple/common"
	libctx "github.com/machmum/lib-simple/context"
	"github.com/sirupsen/logrus"
)

const (
	// log msg
	LLvlAccess   = "access_log"
	LLvlService  = "service_log"
	LLvlPostgres = "mysql_log"
	LLvlInfo     = "info_log"
	LLvlGRPC     = "grpc_log"
	LLvlHTTP     = "http_log"
)

type Options struct {
	// Entry only supports:
	// ErrorLevel
	// WarnLevel
	// FatalLevel
	// InfoLevel
	// DebugLevel
	Level logrus.Level
	// Skip logs log
	Skip bool
}

// New: new json-formatter log with opts
func New(opt Options) {
	var init sync.Once
	init.Do(func() {
		if opt.Skip {
			skip = opt.Skip
		}

		logrus.SetOutput(os.Stdout)
		if opt.Level < 1 {
			logrus.SetLevel(logrus.DebugLevel)
		} else {
			logrus.SetLevel(opt.Level)
		}
		logrus.SetFormatter(&logrus.JSONFormatter{TimestampFormat: libcommon.DatetimeYMDFormat})
	})
	return
}

// NewStd: new json-formatter log
func NewStd() {
	opt := Options{Level: logrus.DebugLevel}
	New(opt)
}

func Access(id string, request *http.Request, response *echo.Response) {
	if isSuccess(response.Status) {
		AccessLog(id, request, response).Info(LLvlAccess)
	} else {
		AccessLog(id, request, response).Warn(LLvlAccess)
	}
}

// AccessLog logs access log
func AccessLog(id string, request *http.Request, response *echo.Response) *logger {
	return &logger{f: accessLog(id, request, response)}
}

func accessLog(id string, request *http.Request, response *echo.Response) logrus.Fields {
	fields := logrus.Fields{
		"id":     id,
		"host":   request.Host,
		"method": request.Method,
		"uri":    request.URL.String(),
		"status": response.Status,
	}
	return fields
}

func Service(ctx context.Context, param, request, response interface{}, err error) {
	if err != nil {
		ServiceLog(ctx, param, request, response, err).Error(LLvlService)
	} else {
		ServiceLog(ctx, param, request, response, err).Info(LLvlService)
	}
}

// ServiceLog logs param, request and response (error if any)
func ServiceLog(ctx context.Context, param, request, response interface{}, err error) Logger {
	return &logger{f: serviceLog(ctx, param, request, response, err)}
}

func serviceLog(ctx context.Context, param, request, response interface{}, err error) logrus.Fields {
	var requestID = libctx.FromRequestIDContext(ctx)
	var fields = make(logrus.Fields, 0)

	pc, file, line, ok := runtime.Caller(3)
	fn := runtime.FuncForPC(pc).Name()

	fields["caller"] = getCaller(ok, fn, line)
	fields["method"] = getMethod(fn, file)
	if requestID != "" {
		fields["id"] = requestID
	}
	if param != nil {
		rb, _ := json.Marshal(map[string]interface{}{"data": param})
		fields["param"] = GetFields(string(rb))
	}
	if request != nil {
		rb, _ := json.Marshal(map[string]interface{}{"data": request})
		fields["request"] = GetFields(string(rb))
	}
	if response != nil {
		rb, _ := json.Marshal(map[string]interface{}{"data": response})
		fields["response"] = GetFields(string(rb))
	}
	if err != nil {
		fields["error"] = err
	}
	return fields
}

func Sql(ctx context.Context, ql QueryHelper, query string, bind ...interface{}) {
	if ql.Error != nil {
		SqlLog(ctx, ql, query, bind...).Error(LLvlPostgres)
	} else {
		SqlLog(ctx, ql, query, bind...).Info(LLvlPostgres)
	}
}

func SqlLog(ctx context.Context, ql QueryHelper, query string, bind ...interface{}) Logger {
	return &logger{f: sqlLog(ctx, ql, query, bind)}
}

func sqlLog(ctx context.Context, ql QueryHelper, query string, bind ...interface{}) logrus.Fields {
	var requestID = libctx.FromRequestIDContext(ctx)
	var fields = make(logrus.Fields, 0)

	pc, file, line, ok := runtime.Caller(3)
	fn := runtime.FuncForPC(pc).Name()

	fields["caller"] = getCaller(ok, fn, line)
	fields["method"] = getMethod(fn, file)
	if requestID != "" {
		fields["id"] = requestID
	}
	fields["query"] = query
	if bind != nil {
		fields["parameter"] = bind[0]
	}
	if ql.Affected > 0 {
		fields["rows_affected"] = ql.Affected
	}
	if ql.Error != nil {
		fields["error"] = ql.Error
	}
	return fields
}

type QueryHelper struct {
	Error    error `json:"-"`
	Affected int64 `json:"-"`
}
