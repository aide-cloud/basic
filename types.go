package basic

type UInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type HInt interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	UInt | HInt | Float
}

type Complex interface {
	~complex64 | ~complex128
}

type String string

type Bool bool

type Byte byte

type Rune rune

type Any any

type Comparable interface {
	Number | String | Byte | Rune
}

type Basic interface {
	Comparable | Complex | Any
}
