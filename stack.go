package basic

type Stack interface {
	container
	Top() (element Any)
}

var _ Stack = (*arrayStack)(nil)
var _ Stack = (*linkStack)(nil)
var _ Stack = (*loopStack)(nil)
