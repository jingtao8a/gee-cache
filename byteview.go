package main

type ByteView struct {
	b []byte
}

func (v ByteView) Len() int {
	return len(v.b)
}

func (v ByteView) BytesSlice() []byte {
	c := make([]byte, len(v.b))
	copy(c, v.b)
	return c
}

func (v ByteView) String() string {
	return string(v.b)
}
