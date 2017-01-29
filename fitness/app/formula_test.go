package app

import (
	"testing"
)

func TestFormula(t *testing.T) {
	result, err := NewFormula("200*4").Evaluate()

	if result != 800 {
		t.Log("Formula should've evaluated to 800.")
	}

	result, err = NewFormula("200").Evaluate()

	t.Log(result)
	t.Log(err)

}
