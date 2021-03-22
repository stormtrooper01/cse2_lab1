package example_bin

import "testing"

func TestHello(t *testing.T) {
	res := Hello()

	if res != "<h1>May the force be with you</h1>" {
		t.Errorf("Text is incorrect")
	}
}
