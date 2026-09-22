package mem

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAllocFree(t *testing.T) {
	req := require.New(t)

	b, err := Alloc(4096)
	req.NoError(err)
	req.Len(b, 4096)

	for i := range b {
		b[i] = byte(i)
	}
	for i := range b {
		req.Equal(byte(i), b[i])
	}

	req.NoError(Free(b))
}
