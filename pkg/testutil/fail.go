package testutil

import "testing"

func TestFailed(t *testing.T, expected, actual interface{}) {
	t.Errorf("test failed\n"+
		"expected: %v\n"+
		"actual:   %v\n", expected, actual)
}
