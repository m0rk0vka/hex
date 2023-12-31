package api

import "github.com/m0rk0vka/hex/internal/ports"

type Adapter struct {
	db    ports.DbPort
	arith ports.ArithmeticPort
}

func NewAdapter(db ports.DbPort, arith ports.ArithmeticPort) *Adapter {
	return &Adapter{db: db, arith: arith}
}

func (apia Adapter) GetAddition(a, b int32) (int32, error) {
	answer, err := apia.arith.Addition(a, b)
	if err != nil {
		return 0, err
	}

	if err := apia.db.AddToHistory(answer, "addition"); err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia Adapter) GetSubstruction(a, b int32) (int32, error) {
	answer, err := apia.arith.Substraction(a, b)
	if err != nil {
		return 0, err
	}

	if err := apia.db.AddToHistory(answer, "substraction"); err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia Adapter) GetMultiplication(a, b int32) (int32, error) {
	answer, err := apia.arith.Multitplication(a, b)
	if err != nil {
		return 0, err
	}

	if err := apia.db.AddToHistory(answer, "multiplication"); err != nil {
		return 0, err
	}

	return answer, nil
}

func (apia Adapter) GetDivision(a, b int32) (int32, error) {
	answer, err := apia.arith.Division(a, b)
	if err != nil {
		return 0, err
	}

	if err := apia.db.AddToHistory(answer, "division"); err != nil {
		return 0, err
	}

	return answer, nil
}
