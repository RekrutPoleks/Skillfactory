package calc

import "errors"

type calculator struct {
	num1     float64
	num2     float64
	operator string
}

func NewCalculator() calculator {
	return calculator{}
}

func (cl *calculator) Calculate(num1, num2 float64, operator string) (float64, error) {
	cl.num1, cl.num2 = num1, num2
	cl.operator = operator
	if callOperator, err := cl.hub(); err != nil {
		return 0, err
	} else {
		return callOperator()
	}

}

func (cl *calculator) hub() (func() (float64, error), error) {
	switch cl.operator {
	case "+":
		return cl.sum, nil
	case "-":
		return cl.diff, nil
	case "*":
		return cl.compos, nil
	case "/":
		return cl.div, nil
	default:
		return nil, errors.New("There is no such operator")
	}
}

func (cl *calculator) sum() (float64, error) {
	return cl.num1 + cl.num2, nil
}

func (cl *calculator) diff() (float64, error) {
	return cl.num1 - cl.num2, nil
}

func (cl *calculator) div() (float64, error) {
	if cl.num2 == 0 {
		return 0, errors.New("You can't divide it by zero")
	}
	return cl.num1 / cl.num2, nil
}

func (cl *calculator) compos() (float64, error) {
	return cl.num1 * cl.num2, nil
}
