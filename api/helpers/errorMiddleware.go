package api_helpers

import (
	"net/http"

	"github.com/viitoormb/go-cleanarch-api/domain/shared"
)

type ResponseError struct {
	Err        string `json:error`
	Message    string `json:message`
	StatusCode int    `json:status_code`
	Details    string `json:details`
}

func HandleError(err error, r *http.Request) (response ResponseError) {
	var statusCode int
	var message string
	response = ResponseError{}

	if e, domainError := err.(*shared.DomainError); domainError {
		response = ResponseError{
			Err:        e.Err.Error(),
			Details:    e.Details,
			StatusCode: e.StatusCode,
			Message:    e.Message,
		}
	} else {
		if statusCode == 0 {
			statusCode = 500
		}

		response = ResponseError{
			Err:        err.Error(),
			Details:    r.RequestURI,
			StatusCode: statusCode,
			Message:    message,
		}
	}

	return response
}
