package op

import (
	"reflect"
	"strings"
)

func IfNilOrWhitespace[T any](mainValue, defaultValue T) T {
	if reflect.ValueOf(mainValue).IsNil() {
		return defaultValue
	}

	if strings.TrimSpace(reflect.ValueOf(mainValue).String()) == "" {
		return defaultValue
	}

	return mainValue
}
