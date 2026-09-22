//go:build !unix
// +build !unix

package mem

// Alloc allocates the required number of bytes on the heap.
//
// On platforms without an anonymous mmap the allocation is served by the Go
// runtime, so the memory is garbage collected and [Free] is a no-op.
func Alloc(size int) ([]byte, error) {
	return make([]byte, size), nil
}

// Free frees the slice allocated by [Alloc].
func Free(b []byte) error {
	return nil
}
