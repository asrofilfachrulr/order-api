package error

type BadRequestError struct {
	Err error
}

func (e *BadRequestError) GetError() error {
	return e.Err
}
