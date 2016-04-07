package travisplay

import (
	"os"
	"testing"
)

func TestFoo(t *testing.T) {
	for _, v := range os.Environ() {
		t.Log(v)
	}
	t.Fail()
}
