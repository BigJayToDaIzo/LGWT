package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Greeter interface {
	Greet(name string) (string, error)
}

func GreetSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("@DN")
	assert.NoError(t, err)
	assert.Equal(t, got, "Hello, @DN")
}

type GreetAdapter func(name string) string

func (g GreetAdapter) Greet(name string) (string, error) {
	return g(name), nil
}

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t *testing.T, meany MeanGreeter) {
	got, err := meany.Curse("DeeNuh")
	assert.NoError(t, err)
	assert.Equal(t, got, "Go to FLORIDA, DeeNuh!")
}
