package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHeader(t *testing.T) {
	require := require.New(t)

	h := NewHeader()
	require.Equal(h.Keys(), []string(nil))

	h.Set("k0", []byte("v0"))
	require.Equal(h.Keys(), []string{"k0"})
	require.Equal(&HeaderField{"k0", []byte("v0")}, h.Get("k0"))

	h.Set("k1", []byte("v1"))
	require.Equal(h.Keys(), []string{"k0", "k1"})
	require.Equal(&HeaderField{"k1", []byte("v1")}, h.Get("k1"))
}
