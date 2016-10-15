package app

import (
	"testing"
)

func TestBodyFat(t *testing.T) {
	bf1, err := BodyFat{}.Compute(BodyFatInput{metric: "cm", gender: "male", height: 179, neck: 39, waist: 89, hip: 0})
	if err != nil {
		t.Log(err)
		return
	}

	if bf1 != 19 {
		t.Log("Expected 19 percent but got ", bf1, "%")
	}

	t.Log(bf1)

	bf1 = 0

	bf1, err = BodyFat{}.Compute(BodyFatInput{metric: "inches", gender: "male", height: 70, neck: 15, waist: 35, hip: 0})
	if bf1 != 19 {
		t.Log("Expected 19 percent but got %", bf1)
	}
}
