package Calculadora

import "testing"

func TestAdd1(t *testing.T) {
	expected := 2
	actual := add(1, 1)
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}
func TestAdd_5_3(t *testing.T) {
	expected := 8
	actual := add(5, 3)
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}
func TestSubs_5_3(t *testing.T) {
	expected := 2
	actual := sub(5, 3)
	if actual != expected {
		t.Errorf("Test failed, expected: '%d', got: '%d'", expected, actual)
	}
}
