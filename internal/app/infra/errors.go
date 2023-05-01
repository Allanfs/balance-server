package infra

type InfraError struct {
	error
	Essential bool
	Name      string
	Message   string
}

func (i InfraError) Error() string {
	return i.error.Error()
}

func NewEssentialError(rootErr error, name, message string) *InfraError {

	return &InfraError{
		Essential: true,
		Name:      name,
		Message:   message,
		error:     rootErr,
	}
}
