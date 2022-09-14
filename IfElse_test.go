package op

import (
	"testing"
)

func TestIfElse(t *testing.T) {
	if IfElse(true, 1, 2) != 1 {
		t.Error("IfElse failed")
	}

	if IfElse(false, 1, 2) != 2 {
		t.Error("IfElse failed")
	}
}
