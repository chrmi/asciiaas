package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/** Main should generally not be tested, as any functionality should live in local or imported modules,
 * this no-op file is to ensure go test does not produce a false error or indicator of a non-working system.
 */
func TestMain(t *testing.T) {
	assert.Equal(t, "noop", "noop")
}
