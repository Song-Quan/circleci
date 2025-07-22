package mathutil

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)
	expected := 5
	if result != expected {
		t.Errorf("Add(2, 3) = %d; want %d", result, expected)
	}
}

func TestSub(t *testing.T) {
	result := Sub(2, 3)
	expected := -1
	if result != expected {
		t.Errorf("Sub(2, 3) = %d; want %d", result, expected)
	}
}

func TestDivide4(t *testing.T) {
	result := Divide4(1, 1)
	expected := 1
	if result != expected {
		t.Errorf("Divide4(1, 1) = %d; want %d", result, expected)
	}
}
