package aritmetica

import (
	"errors"
	"math"
)

type Service interface {
	SumSubtNumbers(req numbersRequest) (float64, error)
	DivideTwoNumbers(req divideRequest) (float64, error)
	MultiplyNumbers(req multiplyRequest) (float64, error)
	RootNumbers(req rootRequest) (float64, error)
	PowerNumbers(req powerRequest) (float64, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &service{
		repository: repository,
	}
}

func (s service) SumSubtNumbers(req numbersRequest) (float64, error) {
	var total float64

	for _, number := range req.ListNumbers {
		total = total + number
	}
	return total, nil
}

func (s service) DivideTwoNumbers(req divideRequest) (float64, error) {

	if req.Denominator == 0 {
		return 0, errors.New("indefinido")
	}

	total := req.Numerator / req.Denominator

	return total, nil
}

func (s service) MultiplyNumbers(req multiplyRequest) (float64, error) {
	var total float64 = 1

	for _, number := range req.ListNumbers {
		total = total * number
	}
	return total, nil
}

func (s service) RootNumbers(req rootRequest) (float64, error) {
	var total float64
	var RootDecimal float64

	if req.RootNumber%2 == 0 && req.NumberToRoot < 0 {
		return 0, errors.New("raíz negativa")
	}
	RootDecimal = 1 / float64(req.RootNumber)
	total = math.Pow(req.NumberToRoot, RootDecimal)

	return total, nil
}

func (s service) PowerNumbers(req powerRequest) (float64, error) {
	var total float64 = 1

	total = math.Pow(req.NumberToPower, req.PowerNumber)

	return total, nil
}


