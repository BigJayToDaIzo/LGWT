package interactions_test

import (
	"testing"

	"example.com/go-specs-greet/domain/interactions"
	specs "example.com/go-specs-greet/specifications"
	"github.com/alecthomas/assert/v2"
)

func TestCurse(t *testing.T) {
	specs.CurseSpecification(t, specs.CurseAdapter(interactions.Curse))
	t.Run("default name to world if it's an empty string", func(t *testing.T) {
		assert.Equal(t, "Go to FLORIDA, World!", interactions.Curse(""))
	})
}
