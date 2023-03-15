package views

import "net/http"

type Response struct {
	Status         int          `json:"status"`
	Message        string       `json:"message"`
	Error          *CustomError `json:"error,omitempty"`
	AdditionalInfo interface{}  `json:"additional_info,omitempty"`
	Payload        interface{}  `json:"payload,omitempty"`
}

type CustomError struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
}

type ErrorType string

const (
	Err_InvalidPayload ErrorType = "INVALID_PAYLOAD"
	Err_BadRequest     ErrorType = "BAD_REQUEST"
	Err_Repository     ErrorType = "REPOSITORY_ERROR"
	Err_NotFound       ErrorType = "DATA_NOT_FOUND"
	Err_InternalServer ErrorType = "INTERNAL_SERVER_ERROR"
)

type MsgType string

const (
	Success_Created MsgType = "CREATED_SUCCESS"
	Success_FindAll MsgType = "FIND_ALL_SUCCESS"
	Err_Created     MsgType = "CREATED_FAIL"
)

func SuccessCreated(payload interface{}) *Response {
	return buildSuccess(payload, http.StatusCreated, Success_Created)
}

func SuccessFindAll(payload interface{}) *Response {
	return buildSuccess(payload, http.StatusOK, Success_FindAll)
}

func BadRequest(err error) *Response {
	return buildError(err, http.StatusBadRequest, Err_BadRequest)
}
func RepositoryError(err error) *Response {
	return buildError(err, http.StatusInternalServerError, Err_Repository)
}
func InternalServerError(err error) *Response {
	return buildError(err, http.StatusInternalServerError, Err_InternalServer)
}

func NotFound(err error) *Response {
	return buildError(err, http.StatusNotFound, Err_NotFound)
}

func buildSuccess(payload interface{}, status int, msg MsgType) *Response {
	return &Response{
		Status:  status,
		Payload: payload,
		Message: string(msg),
	}
}

func buildError(err error, status int, errType ErrorType) *Response {
	return &Response{
		Status: status,
		Error: &CustomError{
			Type:    errType,
			Message: err.Error(),
		},
		AdditionalInfo: err,
	}
}
