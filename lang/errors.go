package lang

import (
	"fmt"
	"os"
)

type BaseError struct {
	Type    string
	Message string
	stack   string
	IsFatal bool
}

func (baseError *BaseError) GetType() string {
	return baseError.Type
}

func (baseError *BaseError) GetMessage() string {
	return baseError.Type
}

func (baseError *BaseError) Run() {
	if baseError.IsFatal {
		fmt.Print(baseError.Type + ": " + baseError.Message)
		fmt.Println(baseError.stack)
		os.Exit(0)
	}
	fmt.Print(baseError.Type + ": " + baseError.Message)
	fmt.Println(baseError.stack)
}

func Errorf(errtype, message, stack string, isFatal bool) *BaseError {
	return &BaseError{errtype, message, stack, isFatal}
}
