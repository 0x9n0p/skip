package skip

type Skip[T any] struct {
	Value T
	Error error
}

func (s *Skip[T]) Catch(f func(err error) error) *Skip[T] {
	if s.Error != nil {
		s.Error = f(s.Error)
	}
	return s
}

func (s *Skip[T]) Replace(err error) *Skip[T] {
	return s.Catch(func(_ error) error {
		return err
	})
}

func (s *Skip[T]) Return() T {
	if s.Error != nil {
		panic(&ReturnError{s.Error})
	}
	return s.Value
}
