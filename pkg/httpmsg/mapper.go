package httpmsg

import (
	"net/http"

	"github.com/iam-benyamin/hellofresh/pkg/errmsg"
	"github.com/iam-benyamin/hellofresh/pkg/richerror"
)

func Error(err error) (message string, code int) {
	switch err.(type) {
	case richerror.RichError:
		re := err.(richerror.RichError)
		msg := re.Message()
		kind := mapKindToStatusCode(re.Kind())
		if kind >= 500 {
			msg = errmsg.ErrorMsgSomeThingWentWrong
		}
		return msg, kind
	default:
		return err.Error(), http.StatusBadRequest
	}
}

func mapKindToStatusCode(kind richerror.Kind) int {
	switch kind {
	case richerror.KindNotFound:
		return http.StatusNotFound
	case richerror.KindUnexpected:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
