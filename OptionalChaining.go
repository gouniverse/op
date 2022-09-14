package op

import "reflect"

func OptionalChaining[T1, T2 any](val *T1, f func() *T2) *T2 {
	if reflect.ValueOf(val).IsNil() {
		return nil
	}

	return f()
}
