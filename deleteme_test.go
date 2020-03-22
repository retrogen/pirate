// Copyright 2020 retrogen <github.com/retrogen>
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the root directory of this source tree.

package pirate

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloPirates(t *testing.T) {
	assert.Equal(t, "Aye Aye Captain!", helloPirates())
}
