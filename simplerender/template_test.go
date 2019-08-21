package simplerender

import (
	"testing"
)

func TestRenderVariable(t *testing.T) {
	got := renderVariable("hello {{name}}", "name", "world")
	expected := "hello world"
	if got != expected {
		t.Errorf("Expected '%s', got '%s'", expected, got)
	}
}
