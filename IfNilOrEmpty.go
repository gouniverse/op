package op

import "reflect"

func IfNilOrEmpty[T any](mainValue, defaultValue T) T {
	if reflect.ValueOf(mainValue).IsNil() {
		return defaultValue
	}

	if reflect.ValueOf(mainValue).String() == "" {
		return defaultValue
	}

	return mainValue
}
