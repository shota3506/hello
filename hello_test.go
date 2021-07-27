package hello

import "testing"

func TestHello(t *testing.T) {
	expected := "hello"
	actual := Hello()
	if expected != actual {
		t.Errorf("Not equal: \n"+
			"expected: %s\n"+
			"actual  : %s", expected, actual)
	}
}
