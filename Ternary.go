package op

func Ternary[T any](condition bool, trueValue T, falseValue T) T {
	return IfElse(condition, trueValue, falseValue)
}
