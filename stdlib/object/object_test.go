package object_test

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/glaukiol1/gago/run"
)

func TestObject(t *testing.T) {
	out1 := captureOutput(func() {
		run.RunData("import object\nconst a = call object.create()\ncall print(a)", "<test>", false)
	})
	if strings.TrimSpace(out1) != "{}" {
		t.Errorf("Expected %s found %s", "{}", out1)
	}
	out2 := captureOutput(func() {
		run.RunData("import object\nconst a = call object.create()\ncall object.set(a, 'test', 1)\ncall print(a)", "<test>", false)
	})
	if strings.TrimSpace(out2) != "{test: 1}" {
		t.Errorf("Expected %s found %s", "{test: 1}", out2)
	}
	out3 := captureOutput(func() {
		run.RunData("import object\nconst a = call object.create()\ncall object.set(a, 'test', 1)\nconst b = call object.get(a, 'test')\ncall print(b)", "<test>", false)
	})
	if strings.TrimSpace(out3) != "1" {
		t.Errorf("Expected %s found %s", "1", out3)
	}
	out4 := captureOutput(func() {
		run.RunData("import object\nconst a = call object.create()\ncall object.set(a, 'test', 1)\ncall object.set(a, 'test1', 1)\ncall print(call object.keys(a))", "<test>", false)
	})
	if strings.TrimSpace(out4) != "[test, test1]" {
		t.Errorf("Expected %s found %s", "[test, test1]", out4)
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
