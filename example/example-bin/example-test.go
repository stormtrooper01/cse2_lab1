package example_bin

import "testing"

func TestForce(t *testing.T) {
	res := Force()

	if res != "<h1>May the force be with you</h1>" {
		t.Errorf("Text is incorrect")
	}
}
