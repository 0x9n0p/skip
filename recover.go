package skip

import (
	"errors"
)

func Recoverf(f func(Error)) {
	v := recover()
	if v == nil {
		return
	}

	err, ok := v.(error)
	if !ok {
		panic(v)
		return
	}

	var nerr Error
	if errors.As(err, &nerr) {
		if !nerr.Skippable() {
			f(nerr)
		}
		return
	}

	panic(err)
}
