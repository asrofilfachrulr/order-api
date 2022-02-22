package error

type NotFoundError struct {
	Err error
}

func (e *NotFoundError) GetError() error {
	return e.Err
}
