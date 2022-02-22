package error

type AuthFailedError struct {
	Err error
}

func (e *AuthFailedError) GetError() error {
	return e.Err
}
