package lang

// All the Types will have at least
// Val() interface{}
// Reassign(newval interface{}) error
// Name() string
// IsConstant() bool
// SetConstant(bool)

type Type interface {
	Val() interface{}
	Reassign(newval interface{}) error
	Name() string
	IsConstant() bool
	SetConstant(bool)
}
