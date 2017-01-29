package app

import (
	"github.com/Knetic/govaluate"
)

type Formula struct {
	expr string
}

func NewFormula(expr string) *Formula {
	return &Formula{expr: expr}
}

func (self *Formula) Evaluate() (float64, error) {
	// split by tokens
	expression, err := govaluate.NewEvaluableExpression(self.expr)

	if err != nil {
		return 0, err
	}

	result, err := expression.Evaluate(nil)

	if err != nil {
		return 0, err
	}

	return result.(float64), nil
}
