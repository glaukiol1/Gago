package array_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/glaukiol1/gago/run"
)

func TestArray(t *testing.T) {
	out1 := captureOutput(func() {
		run.RunData("const array1 = call array.create(1,2,3,4)\ncall print(array1)", "<test>", false)
	})
	if strings.TrimSpace(out1) == "[1, 2, 3, 4] " {
		t.Errorf("Expected %s found %s", "[1, 2, 3, 4]", out1)
	}
	out2 := captureOutput(func() {
		run.RunData("const a = call array.create(true,false,1,2)\nconst indx0 = call array.access(a, 0)\nconst indx1 = call array.access(a, 1)\nconst indx2 = call array.access(a, 2)\nconst indx3 = call array.access(a, 3)\ncall print(indx0,indx1,indx2,indx3)", "<test>", false)
	})
	if strings.TrimSpace(out2) != "true false 1 2" {
		t.Errorf("Expected %s found %s", "true false 1 2", out2)
	}
	out3 := captureOutput(func() {
		run.RunData("const a = call array.create(1,2,3,4)\ncall array.pop(a)\ncall print(a)", "<test>", false)
	})
	if strings.TrimSpace(out3) != "[1, 2, 3]" {
		t.Errorf("Expected %s found %s", "[1, 2, 3]", out3)
	}
	out4 := captureOutput(func() {
		run.RunData("const a = call array.create(1,2,3,4)\ncall array.shift(a)\ncall print(a)", "<test>", false)
	})
	if strings.TrimSpace(out4) != "[2, 3, 4]" {
		t.Errorf("Expected %s found %s", "[2, 3, 4]", out4)
	}
	out5 := captureOutput(func() {
		run.RunData("const a = call array.create(1,2,3,4)\nconst b = call array.subslice(a, 1, -1)\ncall print(b)", "<test>", false)
	})
	if strings.TrimSpace(out5) != "[2, 3, 4]" {
		t.Errorf("Expected %s found %s", "[2, 3, 4]", out5)
	}
}

func captureOutput(s func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	s()

	outC := make(chan string)

	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()

	// back to normal state
	w.Close()
	os.Stdout = old
	out := <-outC

	return out
}
