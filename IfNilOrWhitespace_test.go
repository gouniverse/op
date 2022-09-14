package op

import (
	"testing"
)

func TestIfNilOrWhitespace(t *testing.T) {
	var displayMessage interface{}
	if IfNilOrWhitespace(displayMessage, "failed") != "failed" {
		t.Error("IfNilOrWhitespace failed on nil")
	}

	displayMessage = ""
	if IfNilOrWhitespace(displayMessage, "failed") != "failed" {
		t.Error("IfNilOrWhitespace failed on empty")
	}

	displayMessage = "    "
	if IfNilOrWhitespace(displayMessage, "failed") != "failed" {
		t.Error("IfNilOrWhitespace failed on spaces")
	}

	displayMessage = "  \n \t   "
	if IfNilOrWhitespace(displayMessage, "failed") != "failed" {
		t.Error("IfNilOrWhitespace failed on spaces and newlines")
	}

	displayMessage = "ok"
	if IfNilOrWhitespace(displayMessage, "failed") == "failed" {
		t.Error("IfNilOrWhitespace failed on full")
	}

}
