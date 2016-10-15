package app

import (
	"errors"
	"math"

	"fmt"
	"strconv"
)

type (
	BodyFatInput struct {
		metric string
		gender string
		height float64

		neck  float64
		waist float64
		hip   float64
	}

	BodyFat struct {
	}
)

// Compute using the US Navy Method. For more details see: https://en.wikipedia.org/wiki/Body_fat_percentage
// Formula: http://fitness.stackexchange.com/questions/1660/how-accurate-is-the-navy-body-fat-calculator?newreg=550aaea7d87447b8b8ef6e83e9d74927
func (bf BodyFat) Compute(param BodyFatInput) (float64, error) {
	// Calculating BF percentage for the transgender gets complicated.  See https://www.quora.com/Should-transgender-peoples-body-fat-percentages-be-compared-to-the-charts-for-their-gender-assigned-at-birth-or-their-post-transition-gender
	if param.gender != "male" && param.gender != "female" {
		return 0, errors.New("Please specify the gender. Options include: [male, female].")
	}

	// Convert all inches to cm
	if param.metric == "inches" {
		param.height = param.height * INCH_TO_CM
		param.neck = param.neck * INCH_TO_CM
		param.waist = param.waist * INCH_TO_CM
		param.hip = param.hip * INCH_TO_CM
	}

	var bodyFat float64
	bodyFat = 0

	logHeight := math.Log10(param.height)
	if param.gender == "male" {
		// 495 / (1.29579 - .35004 * log10(Waist - Neck) + 0.22100 * log10(Height)) - 450
		maleFactor := math.Log10(param.waist + param.neck)
		bodyFat = 495/(1.29579-0.35004*maleFactor+0.22100*logHeight) - 450
	} else {
		// female
		// 495 / (1.29579 - .35004 * log10(Waist + Hip - Neck) + 0.22100 * log10(Height)) - 450
		femaleFactor := math.Log10(param.waist + param.hip - param.neck)
		bodyFat = 495/(1.29579-0.35004*femaleFactor+0.22100*logHeight) - 450
	}

	trunc := fmt.Sprintf("%.2f", bodyFat)
	bodyFat, _ = strconv.ParseFloat(trunc, 64)

	return bodyFat, nil
}
