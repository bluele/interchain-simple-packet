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
	v, ok := h.Get("k0")
	require.True(ok)
	require.Equal([]byte("v0"), v)

	h.Set("k1", []byte("v1"))
	require.Equal(h.Keys(), []string{"k0", "k1"})
	v, ok = h.Get("k1")
	require.Equal([]byte("v1"), v)
}
