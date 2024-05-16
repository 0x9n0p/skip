package skip_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/0x9n0p/skip"
)

const (
	KeyID   = "id"
	KeyName = "name"
)

var ErrBadRequest = errors.New("bad request")

type User struct {
	ID   int
	Name string
}

func TestExample(t *testing.T) {
	defer skip.Recoverf(func(err skip.Error) {
		t.Log(err)
	})

	req := map[string]any{
		KeyID:   1,
		KeyName: "Hitler",
	}

	user := User{
		ID:   parse[int](req, KeyID).Replace(ErrBadRequest).Return(),
		Name: parse[string](req, KeyName).Replace(ErrBadRequest).Return(),
	}

	t.Log(user)
}

func parse[T any](req map[string]any, key string) (skp *skip.Skip[T]) {
	skp = new(skip.Skip[T])

	v, found := req[key]
	if !found {
		skp.Error = errors.New("key not found")
		return
	}

	f, ok := v.(T)
	if !ok {
		skp.Error = fmt.Errorf("got data of type %T but wanted %T", v, skp.Value)
		return
	}

	skp.Value = f
	return
}
