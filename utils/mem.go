package luvutils

//The AllocMem function is used to allocate memory
func AllocMem(size int) ([]byte, error) {
	//Add a buffer to the size
	size += 1024
	return make([]byte, size), nil
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
func AllocMemMulti(size int, count int) ([][]byte, error) {

	var mem [][]byte

	memToAloc := size / count

	for i := 0; i < count; i++ {
		mem = append(mem, make([]byte, memToAloc))
	}

	return mem, nil

}
