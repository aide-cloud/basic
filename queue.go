package basic

type container interface {
	Push(element Any) bool
	Pop() (element Any)
	IsEmpty() bool
	IsFull() bool
	Size() int
	String() string
	Clear()
	Values() []Any
}

type Queue interface {
	container
}

var _ Queue = (*linkedQueue)(nil)
var _ Queue = (*arrayQueue)(nil)
var _ Queue = (*loopQueue)(nil)
