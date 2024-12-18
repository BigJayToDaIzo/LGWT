package interactions_test

import (
	"testing"

	"example.com/go-specs-greet/domain/interactions"
	specs "example.com/go-specs-greet/specifications"
	"github.com/alecthomas/assert/v2"
)

func TestGreet(t *testing.T) {
	specs.GreetSpecification(t, specs.GreetAdapter(interactions.Greet))
	t.Run("default name to world if it's an empty string", func(t *testing.T) {
		assert.Equal(t, "Hello, World", interactions.Greet(""))
	})
}
