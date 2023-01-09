package strings

import (
	"strings"
	"testing"
)

const MSG_INDEX = "%s (parte: %s) - Índices: esperado (%d) <> encontrado (%d)."

func TestIndex(t *testing.T) {
	tests := []struct {
		text  string
		part  string
		index int
	}{
		{"Olá mundo", "mundo", 5},
		{"", "", 0},
		{"OPA", "opa", -1},
		{"felipe", "lipe", 2},
	}

	for _, test := range tests {
		t.Logf("Test: %v", test)

		actual := strings.Index(test.text, test.part)

		if actual != test.index {
			t.Errorf(MSG_INDEX, test.text, test.part, test.index, actual)
		}
	}
}
