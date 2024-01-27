package resources

type ErrorResource struct {
	Message string `json:"message"`
	Errors  string `json:"errors"`
}

func NewErrorResource(message string, errors string) *ErrorResource {
	return &ErrorResource{
		Message: message,
		Errors:  errors,
	}
}
