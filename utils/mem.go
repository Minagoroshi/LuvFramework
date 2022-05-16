package luvutils

import "C"
import (
	"runtime"
)

//The AllocMem function is used to allocate memory
//Prone to garbage collection
func AllocMem(size int) ([]byte, error) {
	//Add a buffer to the size
	size += 1024

	mem := make([]byte, size)

	go func() {
		for {
			runtime.KeepAlive(mem)
		}
	}()
	return mem, nil
}

//The FreeMem function is used to free memory
func FreeMem(ptr []byte) {
	//Do nothing
}

//The MemCpy function is used to reallocate memory
func MemCpy(dst, src []byte) {
	copy(dst, src)
}

//AllocMemMulti is used to allocate memory for multiple variables
//Prone to garbage collection
func AllocMemMulti(size int, count int) ([][]byte, error) {

	var mem [][]byte

	memToAloc := size / count

	for i := 0; i < count; i++ {
		mem = append(mem, make([]byte, memToAloc))
	}

	return mem, nil

}

// https://github.com/golang/go/wiki/cgo#turning-c-arrays-into-go-slices
// https://dgraph.io/blog/post/manual-memory-management-golang-jemalloc/

func MallocMemory(size int) ([]byte, error) {
	var slice []byte

	// Make this actually allocate memory

	//ptr := C.je_calloc(C.size_t(n), 1)
	//if ptr == nil {
	//	// NB: throw is like panic, except it guarantees the process will be
	//	// terminated. The call below is exactly what the Go runtime invokes when
	//	// it cannot allocate memory.
	//	throw("out of memory")
	//}
	//uptr := unsafe.Pointer(ptr)
	//
	//atomic.AddInt64(&numBytes, int64(n))
	//// Interpret the C pointer as a pointer to a Go array, then slice.
	//return (*[MaxArrayLen]byte)(uptr)[:n:n]

	return slice, nil
}
