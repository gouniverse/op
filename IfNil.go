package op

func IfNil(mainValue interface{}, defaultValue interface{}) interface{} {
	if mainValue == nil {
		return defaultValue
	}

	return mainValue
}
