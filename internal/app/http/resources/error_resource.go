package resources

type ErrorResource struct {
	Message string `json:"message"`
	Errors  string `json:"errors"`
}
