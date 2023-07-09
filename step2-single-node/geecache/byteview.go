package geecache

/*
byteview是对实际缓存的⼀层封装，因为实际的缓存值是⼀个byte切⽚存储的，
⽽切⽚的底层是⼀个指向底层数组的指针，⼀个记录⻓度的变量和⼀个记录容量的变量。
如果获取缓存值时直接返回缓存值的切⽚，那个切⽚只是原切⽚三个变量的拷⻉，
真正的缓存值就可能被外部恶意修改。
所以⽤byteView进⾏⼀层封装，返回缓存值时的byteView则是⼀个原切⽚的深拷⻉。
*/


// A ByteView holds an immutable view of bytes.
type ByteView struct {
	b []byte
}

// Len returns the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice returns a copy of the data as a byte slice.
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String returns the data as a string, making a copy if necessary.
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
