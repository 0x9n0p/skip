package skip

type Error interface {
	error
	Skippable() bool
}

type ReturnError struct {
	error
}

func (r *ReturnError) Skippable() bool {
	return false
}
