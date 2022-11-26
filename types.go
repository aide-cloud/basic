package basic

type UInt interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

type Int interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Float interface {
	~float32 | ~float64
}

type Number interface {
	UInt | Int | Float
}

type Complex interface {
	~complex64 | ~complex128
}

type String interface {
	~string
}

type Bool interface {
	~bool
}

type Byte interface {
	~byte
}

type Rune interface {
	~rune
}

type Any interface {
	any
}

type Comparable interface {
	Number | String | Bool | Byte | Rune
}

type Basic interface {
	Comparable | Complex | Any
}
