package response

// ErrorsEnvelope represents errors envelop, unsafe for concurrent usage
type ErrorsEnvelope struct {
	Errors map[string][]string `json:"errors"`
}

// NewErrorsEnvelope creates errors envelope
func NewErrorsEnvelope() *ErrorsEnvelope {
	return &ErrorsEnvelope{
		Errors: make(map[string][]string),
	}
}

// Set sets new value for the specific name
func (e *ErrorsEnvelope) Set(name, value string) *ErrorsEnvelope {
	e.Errors[name] = []string{value}
	return e
}

// Add adds new value for the specific name
func (e *ErrorsEnvelope) Add(name, value string) *ErrorsEnvelope {
	errors, ok := e.Errors[name]

	if ok {
		errors = append(errors, value)
	} else {
		errors = []string{value}
	}

	e.Errors[name] = errors
	return e
}
