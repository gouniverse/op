package op

func NilCoalescing(mainValue interface{}, defaultValue interface{}) interface{} {
	return IfNil(mainValue, defaultValue)
}

// func NilCoalescing[T any](mainValue, defaultValue T) T {
// 	if reflect.ValueOf(mainValue).IsNil() {
// 		return defaultValue
// 	}

// 	return mainValue
// }
