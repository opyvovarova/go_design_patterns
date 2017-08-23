package factory

import (
	"testing"
)

func CreateTest(t *testing.T) {
	var expected_name string = "ProductA"
	var c Creator = new(CreatorA)
	p := c.Create()
	if p.Name() != expected_name {
		t.Error("Expected name is not equal to actual one")
	}
}
