package op

import (
	"testing"
)

func TestIfNilOrEmpty(t *testing.T) {
	var displayMessage interface{}
	if IfNilOrEmpty(displayMessage, "failed") != "failed" {
		t.Error("IfNilOrEmpty failed on nil")
	}

	displayMessage = ""
	if IfNilOrEmpty(displayMessage, "failed") != "failed" {
		t.Error("IfNilOrEmpty failed on empty")
	}

	displayMessage = "ok"
	if IfNilOrEmpty(displayMessage, "failed") == "failed" {
		t.Error("IfNilOrEmpty failed on full")
	}

}
