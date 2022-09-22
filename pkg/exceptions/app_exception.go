package exceptions

type AppException struct {
	Message string `json:"message"`
	Code    int
	Stack   error
}

func NewAppException(message string, err error, code int) *AppException {
	return &AppException{
		Message: message,
		Stack:   err,
		Code:    code,
	}
}

func (a *AppException) Error() string {
	return a.Message
}
