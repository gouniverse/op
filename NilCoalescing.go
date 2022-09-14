package op

import "reflect"

func NilCoalescing[T any](mainValue, defaultValue T) T {
	if reflect.ValueOf(mainValue).IsNil() {
		return defaultValue
	}

	return mainValue
}
