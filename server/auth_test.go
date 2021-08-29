package server

import (
	"testing"

	"gotest.tools/assert"
)

func testAuthTokenIsUnique(t *testing.T) {
	t.Parallel()
	token1, err := createToken("user", "pass")
	assert.NilError(t, err)
	token2, err := createToken("user", "pass")
	assert.NilError(t, err)
	assert.Assert(t, token1 != token2)
}
