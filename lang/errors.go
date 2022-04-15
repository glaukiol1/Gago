package lang

import (
	"fmt"
	"os"
)

type BaseError struct {
	Type    string
	Message string
	Stack   string
	IsFatal bool
}

func (baseError *BaseError) GetType() string {
	return baseError.Type
}

func (baseError *BaseError) GetMessage() string {
	return baseError.Type
}

func (baseError *BaseError) Error() string {
	return baseError.Type + ": " + baseError.Message
}

func (baseError *BaseError) Run() {
	if baseError.IsFatal {
		fmt.Print(baseError.Type + ": " + baseError.Message)
		fmt.Println(baseError.Stack)
		os.Exit(1)
	}
	fmt.Print(baseError.Type + ": " + baseError.Message)
	fmt.Println(baseError.Stack)
}

func Errorf(errtype, message, Stack string, isFatal bool) *BaseError {
	return &BaseError{errtype, message, Stack, isFatal}
}
