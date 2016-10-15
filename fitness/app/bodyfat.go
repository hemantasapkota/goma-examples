package app

import (
	"errors"
	"math"
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
//  men: %Fat = 86.010*LOG(abdomen - neck) - 70.041*LOG(height) + 30.30
//  women: %Fat = 163.205*LOG(abdomen + hip - neck) - 97.684*LOG(height) - 78.387
func (bf BodyFat) Compute(param BodyFatInput) (float64, error) {
	// Calculating BF percentage for the transgender gets complicated.  See https://www.quora.com/Should-transgender-peoples-body-fat-percentages-be-compared-to-the-charts-for-their-gender-assigned-at-birth-or-their-post-transition-gender
	if param.gender != "male" && param.gender != "female" {
		return 0, errors.New("Please specify the gender. Options include: [male, female].")
	}

	// Convert all inches to cm
	if param.metric == "inches" {
		param.height = math.Ceil(param.height * INCH_TO_CM)
		param.neck = math.Ceil(param.neck * INCH_TO_CM)
		param.waist = math.Ceil(param.waist * INCH_TO_CM)
		param.hip = math.Ceil(param.hip * INCH_TO_CM)
	}

	var bodyFat float64
	bodyFat = 0

	logHeight := math.Log10(param.height)
	if param.gender == "male" {
		maleFactor := math.Log10(param.waist - param.neck)
		bodyFat = 86.010*maleFactor - 70.041*logHeight + 30.30
	} else {
		// female
		femaleFactor := math.Log10(param.waist + param.hip - param.neck)
		bodyFat = 163.205*femaleFactor - 97.684*logHeight - 78.387
	}

	return math.Ceil(bodyFat), nil
}
