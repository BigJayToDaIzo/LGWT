package go_specs_greet_test

import (
	"testing"

	gsg "example.com/go-specs-greet"
	specs "example.com/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specs.GreetSpecification(t, specs.GreetAdapter(gsg.Greet))
}
