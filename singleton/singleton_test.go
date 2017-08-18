package singleton

import "testing"

func TestNewSingleton(t *testing.T) {
	if NewSingleton() == nil {
		t.Error("Singleton should not be nil")
	}

	if NewSingleton() != NewSingleton() {
		t.Error("Singleton should return the same object every time")
	}
}
