package error

type InternalServerError struct {
	Err error
}

func (e *InternalServerError) GetError() error {
	return e.Err
}
