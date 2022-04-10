package vm

import (
	"fmt"
	"os"
)

// print function using the VMs stdout

// print() prints the `data` string to the stdout
func (vm *VM) print(data string) {
	dt := []byte(data)
	_, err := vm.stdout.Write(dt)
	if err != nil {
		fmt.Println("Error printing to the specified stdout...")
		os.Exit(1)
	}
}

// println uses `print`, but adds a newline character.
func (vm *VM) println(data string) {
	vm.print(data + "\n")
}
