package op

import (
	"testing"
)

func TestIfNil(t *testing.T) {
	var displayMessage interface{}
	if IfNil(displayMessage, "failed") != "failed" {
		t.Error("IfNil failed on nil")
	}

	displayMessage = "ok"
	if IfNil(displayMessage, "failed") == "failed" {
		t.Error("IfNil failed on full")
	}
}
