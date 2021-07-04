package libcommon

import (
	"net/http"
	"time"
)

type Status string

const (
	StdTimeout = 3 * time.Minute

	Received Status = "received"
	Success  Status = "success"

	HeaderContentType  = "Content-Type"
	HeaderAllowOrigin  = "Access-Control-Allow-Origin"
	HeaderAllowMethod  = "Access-Control-Allow-Methods"
	HeaderAllowHeaders = "Access-Control-Allow-Headers"
	HeaderAuth         = "Authorization"

	ContentTypeJson   = "application/json"
	DatetimeYMDFormat = "2006-01-02 15:04:05"
)

var AllowedHeaders = []string{
	HeaderContentType,
	HeaderAuth,
}

var AllowedMethods = []string{
	http.MethodPost,
	http.MethodGet,
}
