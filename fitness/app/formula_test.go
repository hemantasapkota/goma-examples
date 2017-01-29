package app

import (
	"testing"
)

func TestFormula(t *testing.T) {
	result, err := NewFormula("200+40").Evaluate()

	if result != 800 {
		t.Log("Formula should've evaluated to 800.")
	}

	t.Log(result)

	result, err = NewFormula("200").Evaluate()

	t.Log(result)
	t.Log(err)

}
