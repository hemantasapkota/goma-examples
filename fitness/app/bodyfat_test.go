package app

import (
	"testing"
)

func TestBodyFat(t *testing.T) {
	//male, cm
	bfM, _ := BodyFat{}.Compute(BodyFatInput{metric: "cm", gender: "male", height: 179, neck: 39, waist: 89, hip: 0})
	t.Log(bfM, "%")

	//female, cm
	bfX, _ := BodyFat{}.Compute(BodyFatInput{metric: "cm", gender: "female", height: 179, neck: 39, waist: 89, hip: 97})
	t.Log(bfX, "%")

	//male, inchesj
	bfM, _ = BodyFat{}.Compute(BodyFatInput{metric: "inches", gender: "male", height: 70, neck: 15, waist: 35, hip: 0})
	t.Log(bfM, "%")

	//female, inches
	bfX, _ = BodyFat{}.Compute(BodyFatInput{metric: "inches", gender: "female", height: 70, neck: 15, waist: 35, hip: 38})
	t.Log(bfX, "%")
}
