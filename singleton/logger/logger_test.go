package logger

import "testing"

func TestNewLogger(t *testing.T) {
	if NewLogger("test") == nil {
		t.Error("Logger should not be nil")
	}
	if NewLogger("test") != NewLogger("test") {
		t.Error("The two loggers should be the same!")
	}
	if NewLogger("test") != NewLogger("else") {
		t.Error("The two loggers should be the same with the same name!")
	}
}

func TestLog(t *testing.T) {
	var logger = NewLogger("test")
	if logger.Log(INFO, "everything is alright") != "test :: INFO everything is alright\n" {
		t.Error("the two sides are not equivalent")
	}
}
