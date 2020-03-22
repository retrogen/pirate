package pirate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloPirates(t *testing.T) {
	assert.Equal(t, "Aye Aye Captain!", helloPirates())
}
