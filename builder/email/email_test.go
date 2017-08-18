package email

import "testing"

func TestBuilder(t *testing.T) {
	e := EmailBuilder().To("john@doe.com").From("jane@doe.com").Subject("hello").Body("Hello Darling").Build()
	if e.To() != "john@doe.com" || e.From() != "jane@doe.com" || e.Subject() != "hello" || e.Body() != "Hello Darling" {
		t.Error("Builder does not work")
	}
}
