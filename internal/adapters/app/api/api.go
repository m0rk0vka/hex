package api

import "github.com/m0rk0vka/hex/internal/ports"

type Adapter struct {
	arith ports.ArithmeticPort
}

func NewAdapter(arith ports.ArithmeticPort) *Adapter {
	return &Adapter{arith: arith}
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia Adapter) Substruction(a, b int32) (int32, error) {
	answer, err := apia.arith.Substraction(a, b)
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia Adapter) Multitplication(a, b int32) (int32, error) {
	answer, err := apia.arith.Multitplication(a, b)
	if err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	return answer, nil
}
