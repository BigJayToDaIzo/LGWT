package interactions_test

import (
	"testing"

	"example.com/go-specs-greet/domain/interactions"
	specs "example.com/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specs.GreetSpecification(t, specs.GreetAdapter(interactions.Greet))
}
