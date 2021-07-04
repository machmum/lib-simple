package libresp

import (
	"net/http"

	"github.com/labstack/echo"
	libcommon "github.com/machmum/lib-simple/common"
	liberr "github.com/machmum/lib-simple/error"
)

var debug bool

func Init(d bool) {
	debug = d
}

func mappingError(resCode int, err error) (int, string) {
	if debug && err != nil {
		return resCode, err.Error()
	} else {
		// return original error if err.AsMessage set true
		if err != nil {
			if e, ok := err.(liberr.Error); ok && e.AsMessage() {
				return resCode, e.Error()
			}
		}

		switch resCode {
		case http.StatusOK:
			return resCode, ""
		case http.StatusUnauthorized:
			return resCode, "User is not authorized"
		case http.StatusForbidden:
			return resCode, "Access Denied"
		case http.StatusNotFound:
			return resCode, "Data Not Found"
		case http.StatusMethodNotAllowed:
			return resCode, "Action Not Allowed"
		case http.StatusUnprocessableEntity:
			return resCode, "Data can't be processed"
		case http.StatusInternalServerError:
			return resCode, "Internal server error"
		case http.StatusInsufficientStorage:
			return resCode, "Error Database"
		default:
			return resCode, "Unexpected error"
		}
	}
}

// response struct
type Response struct {
	Message string      `json:"message"`
	Code    int         `json:"-"`
	Data    interface{} `json:"data"`
}

func JSON(c echo.Context, code int, result interface{}, err error) error {
	c.Response().Header().Set(libcommon.HeaderContentType, libcommon.ContentTypeJson)

	code, msg := mappingError(code, err)
	if code == http.StatusCreated {
		return c.NoContent(code)
	}

	resp := new(Response)
	resp.Code, resp.Message = code, msg
	resp.Data = result
	return c.JSON(code, resp)
}
