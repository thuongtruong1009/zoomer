package interceptor

type InterceptorSuccessProps struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data"`
}

type InterceptorErrorProps struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"error"`
}
