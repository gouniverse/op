package op

import "reflect"

func IfNilOrEmpty(mainValue interface{}, defaultValue interface{}) interface{} {
	if mainValue == nil {
		return defaultValue
	}

	if reflect.ValueOf(mainValue).String() == "" {
		return defaultValue
	}

	return mainValue
}

// import "reflect"

// func IfNilOrEmpty[T any](mainValue, defaultValue T) T {
// 	if reflect.ValueOf(mainValue).IsNil() {
// 		return defaultValue
// 	}

// 	if reflect.ValueOf(mainValue).String() == "" {
// 		return defaultValue
// 	}

// 	return mainValue
// }
