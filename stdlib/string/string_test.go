package string_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/glaukiol1/gago/run"
)

func TestString(t *testing.T) {
	out1 := captureOutput(func() {
		run.RunData("import string\nconst a = 'hello'\nconst b = 'world'\nconst c = call string.concat(a, ' ', b)\ncall print(c)", "<test>", false)
	})
	if strings.TrimSpace(out1) != "hello world" {
		t.Errorf("Expected %s found %s", "hello world", out1)
	}
	out2 := captureOutput(func() {
		run.RunData("import string\nconst a = 'hello world'\nconst b = call string.contains(a, 'world')\ncall print(b)", "<test>", false)
	})
	if strings.TrimSpace(out2) != "true" {
		t.Errorf("Expected %s found %s", "true", out2)
	}
	out3 := captureOutput(func() {
		run.RunData("import string\nconst a = 'hello world'\nconst b = call string.containsAny(a, 'o')\ncall print(b)", "<test>", false)
	})
	if strings.TrimSpace(out3) != "true" {
		t.Errorf("Expected %s found %s", "true", out3)
	}
	out4 := captureOutput(func() {
		run.RunData("import string\nconst a = '   hello world  '\nconst b = call string.trimSpace(a)\ncall print('|', b, '|')", "<test>", false)
	})
	if strings.TrimSpace(out4) != "| hello world |" {
		t.Errorf("Expected %s found %s", "| hello world |", out4)
	}
	out5 := captureOutput(func() {
		run.RunData("import string\nconst a = 'hello world'\nconst b = call string.index(a, 'h')\ncall print(b)", "<test>", false)
	})
	if strings.TrimSpace(out5) != "0" {
		t.Errorf("Expected %s found %s", "0", out5)
	}
	out6 := captureOutput(func() {
		run.RunData("import string\nconst a = 'hello world'\nconst b = call string.len(a)\ncall print(b)", "<test>", false)
	})
	if strings.TrimSpace(out6) != "11" {
		t.Errorf("Expected %s found %s", "11", out6)
	}
	out7 := captureOutput(func() {
		run.RunData("import string\nconst a = 'hello world'\nconst b = call string.charAt(a, 0)\ncall print(b)", "<test>", false)
	})
	if strings.TrimSpace(out7) != "h" {
		t.Errorf("Expected %s found %s", "h", out7)
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
