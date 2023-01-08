package math

import "testing"

const DEFAULT_ERROR = "Valor esperado %v, mas o resultado foi %v"

func TestMedia(t *testing.T) {
	expectedValue := 7.28
	value := Media(7.2, 9.9, 6.1, 5.9)

	if value != expectedValue {
		t.Errorf(DEFAULT_ERROR, expectedValue, value)
	}
}
