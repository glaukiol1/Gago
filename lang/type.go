package lang

// All the Types will have at least
// these two functions:
// Val() interface{}
// Reassign(newval interface{})
// Name() string

type Type interface {
	Val() interface{}
	Reassign(newval interface{}) error
	Name() string
	IsConstant() bool
	SetConstant(bool)
}
