package skip

type Error interface {
	error
	Unwrap() error
	Skippable() bool
}

type ReturnError struct {
	error
}

func (r *ReturnError) Unwrap() error {
	return r.error
}

func (r *ReturnError) Skippable() bool {
	return false
}
